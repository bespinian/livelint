package livelint_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/bespinian/livelint/internal/livelint"
	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

var services = []apiv1.Service{
	{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "service1",
			Namespace: "namespace",
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{},
			Ports: []apiv1.ServicePort{
				{
					Port:       80,
					Protocol:   apiv1.ProtocolTCP,
					TargetPort: intstr.FromInt(8080),
					Name:       "http",
				},
			},
		},
	},
}

var ingresses = []netv1.Ingress{
	{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ingress",
			Namespace: "namespace",
		},
		Spec: netv1.IngressSpec{
			Rules: []netv1.IngressRule{
				{
					Host: "host.test.com",
					IngressRuleValue: netv1.IngressRuleValue{
						HTTP: &netv1.HTTPIngressRuleValue{
							Paths: []netv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathTypePrefix,
									Backend: netv1.IngressBackend{
										Service: &netv1.IngressServiceBackend{
											Name: "service1",
											Port: netv1.ServiceBackendPort{
												Number: 80,
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
	},
}

var pathTypePrefix = netv1.PathTypePrefix

var errConnection = errors.New("unable to resolve host")

func TestCheckCanVisitPublicApp(t *testing.T) {
	t.Parallel()
	cases := []struct {
		it              string
		ingresses       []netv1.Ingress
		roundTripFunc   RoundTripFunc
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it:        "succeeds, if every url deducable from ingress is reachable",
			ingresses: ingresses,
			roundTripFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
					Header:     make(http.Header),
				}, nil
			},
			expectedToFail:  false,
			expectedMessage: "You can visit the app from the public internet",
		},
		{
			it:        "fails, if no ingresses are found",
			ingresses: []netv1.Ingress{},
			roundTripFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
					Header:     make(http.Header),
				}, nil
			},
			expectedToFail:  true,
			expectedMessage: "You cannot visit the app from the public internet",
		},
		{
			it:        "fails, if an url is reachable but produces an http error",
			ingresses: ingresses,
			roundTripFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(bytes.NewBufferString(`Error`)),
					Header:     make(http.Header),
				}, nil
			},
			expectedToFail:  true,
			expectedMessage: "You cannot visit the app from the public internet",
		},
		{
			it:        "fails, if an url is not reachable",
			ingresses: ingresses,
			roundTripFunc: func(req *http.Request) (*http.Response, error) {
				return nil, errConnection
			},
			expectedToFail:  true,
			expectedMessage: "You cannot visit http://host.test.com:80/ from public internet: Get \"http://host.test.com:80/\": unable to resolve host",
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			k8s := &KubernetesInterfaceMock{
				NetworkingV1Func: func() v1.NetworkingV1Interface {
					return &NetworkingV1InterfaceMock{
						IngressesFunc: func(namespace string) v1.IngressInterface {
							return &IngressInterfaceMock{
								ListFunc: func(ctx context.Context, opts metav1.ListOptions) (*netv1.IngressList, error) {
									return &netv1.IngressList{
										Items: tc.ingresses,
									}, nil
								},
							}
						},
					}
				},
			}
			ll := livelint.Livelint{
				K8s:  k8s,
				HTTP: NewTestClient(tc.roundTripFunc),
			}
			result := ll.CheckCanVisitPublicApp("namespace", services)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
