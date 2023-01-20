package livelint_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/bespinian/livelint/internal/livelint"
	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedapiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func TestCheckReadinessProbe(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pods            []apiv1.Pod
		events          []apiv1.Event
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if no events indicating failed readiness probes are present for any pod within the last 5 minutes",
			pods: []apiv1.Pod{
				{ObjectMeta: metav1.ObjectMeta{Name: "pod1", Namespace: "namespace"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "pod2", Namespace: "namespace"}},
			},
			events: []apiv1.Event{
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod1", Namespace: "namespace"},
					Reason:         "Running",
					Count:          2,
				},
			},
			expectedToFail:  false,
			expectedMessage: "There are no failing Readiness probes",
		},
		{
			it: "fails if there's events indicating failed readiness probes are present for a pod within the last 5 minutes",
			pods: []apiv1.Pod{
				{ObjectMeta: metav1.ObjectMeta{Name: "pod1", Namespace: "namespace"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "pod2", Namespace: "namespace"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "pod3", Namespace: "namespace"}},
			},
			events: []apiv1.Event{
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod1", Namespace: "namespace"},
					Reason:         "Running",
					Count:          2,
				},
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod1", Namespace: "namespace"},
					Reason:         "Running",
					Count:          2,
				},
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod2", Namespace: "namespace"},
					Reason:         "Unhealthy",
					Count:          4,
					Message:        "Readiness probe failed: ERROR",
					LastTimestamp:  metav1.NewTime(time.Now().Add(time.Minute * -1)),
				},
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod3", Namespace: "namespace"},
					Reason:         "Unhealthy",
					Count:          4,
					Message:        "Readiness probe failed: ERROR",
					LastTimestamp:  metav1.NewTime(time.Now().Add(time.Minute * -4)),
				},
			},
			expectedToFail:  true,
			expectedMessage: "There are 2 failing ReadinessProbes",
		},
		{
			it: "succeeds if a readiness probe has failed but no more than three times",
			pods: []apiv1.Pod{
				{ObjectMeta: metav1.ObjectMeta{Name: "pod1", Namespace: "namespace"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "pod2", Namespace: "namespace"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "pod3", Namespace: "namespace"}},
			},
			events: []apiv1.Event{
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod1", Namespace: "namespace"},
					Reason:         "Running",
					Count:          2,
				},
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod1", Namespace: "namespace"},
					Reason:         "Running",
					Count:          2,
				},
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod2", Namespace: "namespace"},
					Reason:         "Unhealthy",
					Count:          3,
					Message:        "Readiness probe failed: ERROR",
					LastTimestamp:  metav1.NewTime(time.Now().Add(time.Minute * -1)),
				},
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod3", Namespace: "namespace"},
					Reason:         "Unhealthy",
					Count:          2,
					Message:        "Readiness probe failed: ERROR",
					LastTimestamp:  metav1.NewTime(time.Now().Add(time.Minute * -1)),
				},
			},
			expectedToFail:  false,
			expectedMessage: "There are no failing Readiness probes",
		},
		{
			it: "succeeds if a readiness probe has failed but not within the last 5 minutes",
			pods: []apiv1.Pod{
				{ObjectMeta: metav1.ObjectMeta{Name: "pod1", Namespace: "namespace"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "pod2", Namespace: "namespace"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "pod3", Namespace: "namespace"}},
			},
			events: []apiv1.Event{
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod1", Namespace: "namespace"},
					Reason:         "Running",
					Count:          2,
				},
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod1", Namespace: "namespace"},
					Reason:         "Running",
					Count:          2,
				},
				{
					InvolvedObject: apiv1.ObjectReference{Name: "pod3", Namespace: "namespace"},
					Reason:         "Unhealthy",
					Count:          2,
					Message:        "Readiness probe failed: ERROR",
					LastTimestamp:  metav1.NewTime(time.Now().Add(time.Minute * -6)),
				},
			},
			expectedToFail:  false,
			expectedMessage: "There are no failing Readiness probes",
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
								ListFunc: func(ctx context.Context, opts metav1.ListOptions) (*apiv1.EventList, error) {
									matchingEvents := []apiv1.Event{}
									for _, e := range tc.events {
										if strings.Contains(opts.FieldSelector, e.InvolvedObject.Name) {
											matchingEvents = append(matchingEvents, e)
										}
									}
									return &apiv1.EventList{Items: matchingEvents}, nil
								},
							}
						},
					}
				},
			}
			ll := livelint.Livelint{
				K8s: k8s,
			}
			result := ll.CheckReadinessProbe(tc.pods)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
