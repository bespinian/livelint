package livelint

import (
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
)

func TestCheckAreAllPodsRunning(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pods            []apiv1.Pod
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if all pods are running",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase: apiv1.PodRunning,
						ContainerStatuses: []apiv1.ContainerStatus{
							{
								State: apiv1.ContainerState{
									Running: &apiv1.ContainerStateRunning{},
								},
							},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "All Pods are RUNNING",
		},
		{
			it:              "succeeds if there are no pods",
			pods:            []apiv1.Pod{},
			expectedToFail:  false,
			expectedMessage: "All Pods are RUNNING",
		},
		{
			it: "fails if there are non-running pods",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase: apiv1.PodFailed,
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "There is 1 Pod that is not RUNNING",
		},
		{
			it: "fails if there are pods with non-running containers",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase: apiv1.PodRunning,
						ContainerStatuses: []apiv1.ContainerStatus{
							{
								State: apiv1.ContainerState{
									Running: nil,
								},
							},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "There is 1 Pod that is not RUNNING",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result := checkAreAllPodsRunning(tc.pods)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
