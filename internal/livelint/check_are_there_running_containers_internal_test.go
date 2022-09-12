package livelint

import (
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
)

func TestCheckAreThereRunningContainers(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pod             apiv1.Pod
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if at least one container is running",
			pod: apiv1.Pod{
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "STATUS",
							State: apiv1.ContainerState{Running: &apiv1.ContainerStateRunning{}},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "There are RUNNING containers",
		},
		{
			it: "fails if no container is running",
			pod: apiv1.Pod{
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name: "STATUS",
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "There are no RUNNING containers",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result := checkAreThereRunningContainers(tc.pod)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
