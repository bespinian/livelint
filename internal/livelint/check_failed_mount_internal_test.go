package livelint

import (
	"context"
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedapiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func TestCheckFailedMount(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pod             apiv1.Pod
		podEvents       []apiv1.Event
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if there is not event with reason FailedMount",
			pod: apiv1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "PODNAME",
					Namespace: "NAMESPACE",
				},
			},
			podEvents:       []apiv1.Event{},
			expectedToFail:  false,
			expectedMessage: "There appear to be no issues mounting volumes.",
		},
		{
			it: "fails if there is at least one event for that pod with reason FailedMount",
			pod: apiv1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "PODNAME",
					Namespace: "NAMESPACE",
				},
			},
			podEvents: []apiv1.Event{
				{
					Reason: "FailedMount",
				},
			},
			expectedToFail:  true,
			expectedMessage: "The Pod is unable to mount a volume",
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
						EventsFunc: func(namespace string) typedapiv1.EventInterface {
							return &apiv1EventInterfaceMock{
								ListFunc: func(ctx context.Context, opts metav1.ListOptions) (*apiv1.EventList, error) {
									return &apiv1.EventList{Items: tc.podEvents}, nil
								},
							}
						},
					}
				},
			}
			ll := Livelint{
				k8s: k8s,
			}
			result := ll.checkFailedMount(tc.pod)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
