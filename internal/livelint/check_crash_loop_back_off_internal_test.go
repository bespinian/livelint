package livelint

import (
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
)

func TestCheckCrashLoopBackOff(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pod             apiv1.Pod
		containerName   string
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it:            "succeeds if the container is not waiting with reason CrashLoopBackOff",
			containerName: "container1",
			pod: apiv1.Pod{
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "container1",
							State: apiv1.ContainerState{Running: &apiv1.ContainerStateRunning{}},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "The Pod is not in a crash loop",
		},
		{
			it:            "succeeds if there is only another container waiting with reason CrashLoopBackOff",
			containerName: "container1",
			pod: apiv1.Pod{
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "container1",
							State: apiv1.ContainerState{Running: &apiv1.ContainerStateRunning{}},
						},
						{
							Name:  "container2",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "CrashLoopBackOff"}},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "The Pod is not in a crash loop",
		},
		{
			it:            "fails if the specified container is waiting with reason CrashLoopBackOff",
			containerName: "container2",
			pod: apiv1.Pod{
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "container1",
							State: apiv1.ContainerState{Running: &apiv1.ContainerStateRunning{}},
						},
						{
							Name:  "container2",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "CrashLoopBackOff"}},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "The Pod status is CrashLoopBackOff",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result := checkCrashLoopBackOff(tc.pod, tc.containerName)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
