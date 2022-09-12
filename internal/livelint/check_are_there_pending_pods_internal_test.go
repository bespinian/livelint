package livelint

import (
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
)

func TestCheckAreTherePendingPods(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pods            []apiv1.Pod
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if no pods are pending",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase: apiv1.PodRunning,
						ContainerStatuses: []apiv1.ContainerStatus{
							{Name: "STATUS"},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "There are no PENDING Pods",
		},
		{
			it:              "succeeds if there are no pods",
			pods:            []apiv1.Pod{},
			expectedToFail:  false,
			expectedMessage: "There are no PENDING Pods",
		},
		{
			it: "fails if there are pending pods",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase: apiv1.PodPending,
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "There is 1 PENDING Pod",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result := checkAreTherePendingPods(tc.pods)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
