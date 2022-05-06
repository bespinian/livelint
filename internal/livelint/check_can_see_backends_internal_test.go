package livelint

import (
	"context"
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedapiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func TestCheckCanSeeBackends(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		ingress         netv1.Ingress
		services        []apiv1.Service
		endpoints       []apiv1.Endpoints
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if there are visible backends for all ingress rules",
			ingress: netv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "INGRESSNAME"},
				Spec: netv1.IngressSpec{
					Rules: []netv1.IngressRule{
						{
							IngressRuleValue: netv1.IngressRuleValue{
								HTTP: &netv1.HTTPIngressRuleValue{
									Paths: []netv1.HTTPIngressPath{
										{
											Path: "/path1",
											Backend: netv1.IngressBackend{
												Service: &netv1.IngressServiceBackend{
													Name: "service1",
													Port: netv1.ServiceBackendPort{Number: 80},
												},
											},
										},
									},
								},
							},
						},
						{
							IngressRuleValue: netv1.IngressRuleValue{
								HTTP: &netv1.HTTPIngressRuleValue{
									Paths: []netv1.HTTPIngressPath{
										{
											Path: "/path2",
											Backend: netv1.IngressBackend{
												Service: &netv1.IngressServiceBackend{
													Name: "service2",
													Port: netv1.ServiceBackendPort{Number: 8080},
												},
											},
										},
										{
											Path: "/path3",
											Backend: netv1.IngressBackend{
												Service: &netv1.IngressServiceBackend{
													Name: "service2",
													Port: netv1.ServiceBackendPort{Number: 8081},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			services: []apiv1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "service1"},
					Spec:       apiv1.ServiceSpec{Ports: []apiv1.ServicePort{{Port: 80}}},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "service2"},
					Spec: apiv1.ServiceSpec{Ports: []apiv1.ServicePort{
						{Port: 8080},
						{Port: 8081},
					}},
				},
			},
			endpoints: []apiv1.Endpoints{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "service1"},
					Subsets:    []apiv1.EndpointSubset{{}},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "service2"},
					Subsets:    []apiv1.EndpointSubset{{}},
				},
			},
			expectedToFail:  false,
			expectedMessage: "There are Backends available for Ingress INGRESSNAME",
		},
		{
			it: "fails if there is no service with port and name matches to the ingress path",
			ingress: netv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "INGRESSNAME"},
				Spec: netv1.IngressSpec{
					Rules: []netv1.IngressRule{
						{
							IngressRuleValue: netv1.IngressRuleValue{
								HTTP: &netv1.HTTPIngressRuleValue{
									Paths: []netv1.HTTPIngressPath{
										{
											Path: "/path1",
											Backend: netv1.IngressBackend{
												Service: &netv1.IngressServiceBackend{
													Name: "service1",
													Port: netv1.ServiceBackendPort{Number: 80},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			services: []apiv1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "service1"},
					Spec:       apiv1.ServiceSpec{Ports: []apiv1.ServicePort{{Port: 1337}}},
				},
			},
			endpoints: []apiv1.Endpoints{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "service1"},
					Subsets:    []apiv1.EndpointSubset{{}},
				},
			},
			expectedToFail:  true,
			expectedMessage: "No backends available for Ingress path /path1",
		},
		{
			it: "fails if there are no endpoint subsets available for a service backend",
			ingress: netv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "INGRESSNAME"},
				Spec: netv1.IngressSpec{
					Rules: []netv1.IngressRule{
						{
							IngressRuleValue: netv1.IngressRuleValue{
								HTTP: &netv1.HTTPIngressRuleValue{
									Paths: []netv1.HTTPIngressPath{
										{
											Path: "/path1",
											Backend: netv1.IngressBackend{
												Service: &netv1.IngressServiceBackend{
													Name: "service1",
													Port: netv1.ServiceBackendPort{Number: 80},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			services: []apiv1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "service1"},
					Spec:       apiv1.ServiceSpec{Ports: []apiv1.ServicePort{{Port: 80}}},
				},
			},
			endpoints: []apiv1.Endpoints{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "service1"},
					Subsets:    []apiv1.EndpointSubset{},
				},
			},
			expectedToFail:  true,
			expectedMessage: "No backends available for Ingress path /path1, because service service1 has no endpoints",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			k8s := &kubernetesInterfaceMock{
				CoreV1Func: func() typedapiv1.CoreV1Interface {
					return &apiv1InterfaceMock{
						ServicesFunc: func(namespace string) typedapiv1.ServiceInterface {
							return &apiv1ServiceInterfaceMock{
								GetFunc: func(ctx context.Context, serviceName string, options metav1.GetOptions) (*apiv1.Service, error) {
									for _, service := range tc.services {
										if service.Name == serviceName {
											return &service, nil
										}
									}
									return nil, nil
								},
							}
						},
						EndpointsFunc: func(namespace string) typedapiv1.EndpointsInterface {
							return &apiv1EndpointsInterfaceMock{
								GetFunc: func(ctx context.Context, name string, opts metav1.GetOptions) (*apiv1.Endpoints, error) {
									for _, endpoint := range tc.endpoints {
										if endpoint.Name == name {
											return &endpoint, nil
										}
									}
									return nil, nil
								},
							}
						},
					}
				},
			}
			ll := Livelint{
				k8s: k8s,
			}
			result := ll.checkCanSeeBackends(tc.ingress, "NAMESPACE")

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
