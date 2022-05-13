package livelint

import (
	"context"
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedapiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func TestCheckCanSeeEndpoints(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		endpoint        apiv1.Endpoints
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if there is an endpoint with subsets available for the service",
			endpoint: apiv1.Endpoints{
				ObjectMeta: metav1.ObjectMeta{Name: "service1"},
				Subsets:    []apiv1.EndpointSubset{{}},
			},
			expectedToFail:  false,
			expectedMessage: "Endpoints exists for service SERVICENAME",
		},
		{
			it: "fails if there is no endpoitn with subsets available for the service",
			endpoint: apiv1.Endpoints{
				ObjectMeta: metav1.ObjectMeta{Name: "service1"},
				Subsets:    []apiv1.EndpointSubset{},
			},
			expectedToFail:  true,
			expectedMessage: "No endpoints exists on the service SERVICENAME",
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
						EndpointsFunc: func(namespace string) typedapiv1.EndpointsInterface {
							return &apiv1EndpointsInterfaceMock{
								GetFunc: func(ctx context.Context, name string, opts metav1.GetOptions) (*apiv1.Endpoints, error) {
									return &tc.endpoint, nil
								},
							}
						},
					}
				},
			}
			ll := Livelint{
				k8s: k8s,
			}
			result := ll.checkCanSeeEndpoints("SERVICENAME", "NAMESPACE")

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
