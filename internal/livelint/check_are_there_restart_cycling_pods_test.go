package livelint_test

import (
	"context"
	"testing"
	"time"

	"github.com/bespinian/livelint/internal/livelint"
	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedapiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func TestCheckAreThereRestartCyclingPods(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pods            []apiv1.Pod
		events          []apiv1.Event
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if no pods are restart cycling",
			pods: []apiv1.Pod{
				{},
			},
			events: []apiv1.Event{
				{
					Reason: "BackOff",
					Count:  2,
				},
			},
			expectedToFail:  false,
			expectedMessage: "No Pods are restart cycling",
		},
		{
			it:              "succeeds if there are no pods",
			expectedToFail:  false,
			expectedMessage: "No Pods are restart cycling",
		},
		{
			it: "succeeds if pod has no events",
			pods: []apiv1.Pod{
				{},
			},
			expectedToFail:  false,
			expectedMessage: "No Pods are restart cycling",
		},
		{
			it: "fails if pods are restart cycling",
			pods: []apiv1.Pod{
				{},
			},
			events: []apiv1.Event{
				{
					Reason: "BackOff",
					Count:  6,
				},
				{
					Reason:        "Unhealthy",
					LastTimestamp: metav1.NewTime(time.Now()),
				},
			},
			expectedToFail:  true,
			expectedMessage: "There is 1 Pod cycling between RUNNING and CRASHING",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			k8s := &KubernetesInterfaceMock{
				CoreV1Func: func() typedapiv1.CoreV1Interface {
					return &Apiv1InterfaceMock{
						EventsFunc: func(string) typedapiv1.EventInterface {
							return &Apiv1EventInterfaceMock{
								ListFunc: func(context.Context, metav1.ListOptions) (*apiv1.EventList, error) {
									return &apiv1.EventList{Items: tc.events}, nil
								},
							}
						},
					}
				},
			}
			ll := livelint.Livelint{
				K8s: k8s,
			}
			result := ll.CheckAreThereRestartCyclingPods(tc.pods)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
