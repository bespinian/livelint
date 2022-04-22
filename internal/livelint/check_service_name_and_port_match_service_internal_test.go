package livelint

import (
	"context"
	"errors"
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	typedapiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func TestCheckTargetPortMatchesContainerPort(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pods            []apiv1.Pod
		services        []apiv1.Service
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if the service targetport is exposed by the container",
			pods: []apiv1.Pod{
				{
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							{
								Ports: []apiv1.ContainerPort{
									{
										ContainerPort: 42,
										Protocol:      "TCP",
									},
								},
							},
						},
					},
				},
			},
			services: []apiv1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "WRONGSERVICENAME",
					},
					Spec: apiv1.ServiceSpec{
						Ports: []apiv1.ServicePort{
							{
								TargetPort: intstr.FromInt(999),
								Protocol:   "TCP",
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "SERVICENAME",
					},
					Spec: apiv1.ServiceSpec{
						Ports: []apiv1.ServicePort{
							{
								TargetPort: intstr.FromInt(42),
								Protocol:   "TCP",
							},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "The targetPorts on the Service match the containerPorts in the Pod",
		},
		{
			it: "succeeds if multiple service ports are correctly exposed by thee container",
			pods: []apiv1.Pod{
				{
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							{
								Ports: []apiv1.ContainerPort{
									{
										ContainerPort: 41,
										Protocol:      "TCP",
									},
									{
										ContainerPort: 42,
										Protocol:      "TCP",
									},
									{
										ContainerPort: 43,
										Protocol:      "TCP",
									},
								},
							},
						},
					},
				},
			},
			services: []apiv1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "SERVICENAME",
					},
					Spec: apiv1.ServiceSpec{
						Ports: []apiv1.ServicePort{
							{
								TargetPort: intstr.FromInt(42),
								Protocol:   "TCP",
							},
							{
								TargetPort: intstr.FromInt(43),
								Protocol:   "TCP",
							},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "The targetPorts on the Service match the containerPorts in the Pod",
		},
		{
			it: "fails if the service targetport is not exposed by the container",
			pods: []apiv1.Pod{
				{
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							{
								Ports: []apiv1.ContainerPort{
									{
										ContainerPort: 42,
										Protocol:      "TCP",
									},
								},
							},
						},
					},
				},
			},
			services: []apiv1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "SERVICENAME",
					},
					Spec: apiv1.ServiceSpec{
						Ports: []apiv1.ServicePort{
							{
								TargetPort: intstr.FromInt(999),
								Protocol:   "TCP",
							},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "The targetPorts TCP 999 on the Service don't match the containerPort in the Pod",
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
									return nil, errServiceNotFound
								},
							}
						},
					}
				},
			}
			ll := Livelint{
				k8s: k8s,
			}
			result := ll.checkTargetPortMatchesContainerPort(tc.pods, "SERVICENAME", "NAMESPACE")

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}

var errServiceNotFound = errors.New("Service was not found in provided list")
