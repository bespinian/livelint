package livelint

import (
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
)

func TestCheckImagePullErrors(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pod             apiv1.Pod
		containerName   string
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it:            "succeeds if the container does not have a waiting status",
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
			expectedMessage: "The Pod is not in status ImagePullBackOff",
		},
		{
			it:            "succeeds if the container does not have a waiting status with reason ErrImagePull or ImagePullBackOff",
			containerName: "container1",
			pod: apiv1.Pod{
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "container1",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "SomethingElse"}},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "The Pod is not in status ImagePullBackOff",
		},
		{
			it:            "succeeds if other containers have a waiting status with reason ErrImagePull or ImagePullBackOff",
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
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "ErrImagePull"}},
						},
						{
							Name:  "container3",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "ImagePullBackOff"}},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "The Pod is not in status ImagePullBackOff",
		},
		{
			it:            "fails if the container has a waiting status with reason ErrImagePull",
			containerName: "container1",
			pod: apiv1.Pod{
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "container1",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "ErrImagePull"}},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "The Pod is in status ErrImagePull",
		},
		{
			it:            "fails if the container has a waiting status with reason ImagePullBackOff",
			containerName: "container1",
			pod: apiv1.Pod{
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "container1",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "ImagePullBackOff"}},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "The Pod is in status ImagePullBackOff",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result := checkImagePullErrors(tc.pod, tc.containerName)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
