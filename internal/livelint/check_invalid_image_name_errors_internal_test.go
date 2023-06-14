package livelint

import (
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
)

func TestCheckInvalidImageName(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pod             apiv1.Pod
		container       apiv1.Container
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds, if the container is not found in the pod",
			pod: apiv1.Pod{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name: "container1",
						},
						{
							Name: "container2",
						},
					},
				},
			},
			container:       apiv1.Container{Name: "container3"},
			expectedToFail:  false,
			expectedMessage: "All image names are valid",
		},
		{
			it: "succeeds, if no container in the pod is in state waiting",
			pod: apiv1.Pod{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name: "container1",
						},
						{
							Name: "container2",
						},
					},
				},
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name: "container1",
							State: apiv1.ContainerState{
								Running: &apiv1.ContainerStateRunning{},
							},
						},
						{
							Name: "container2",
							State: apiv1.ContainerState{
								Running: &apiv1.ContainerStateRunning{},
							},
						},
					},
				},
			},
			container:       apiv1.Container{Name: "container3"},
			expectedToFail:  false,
			expectedMessage: "All image names are valid",
		},
		{
			it: "succeeds, if there are containers in the pod in state waiting, but the reason does not match",
			pod: apiv1.Pod{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name: "container1",
						},
						{
							Name: "container2",
						},
					},
				},
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name: "container1",
							State: apiv1.ContainerState{
								Waiting: &apiv1.ContainerStateWaiting{
									Reason: "SomeOtherReason",
								},
							},
						},
						{
							Name: "container2",
							State: apiv1.ContainerState{
								Running: &apiv1.ContainerStateRunning{},
							},
						},
					},
				},
			},
			container:       apiv1.Container{Name: "container3"},
			expectedToFail:  false,
			expectedMessage: "All image names are valid",
		},
		{
			it: "fails, if there are containers in the pod in state waiting and the reason matches",
			pod: apiv1.Pod{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name: "container1",
						},
						{
							Name: "container2",
						},
					},
				},
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name: "container1",
							State: apiv1.ContainerState{
								Waiting: &apiv1.ContainerStateWaiting{
									Reason:  "InvalidImageName",
									Message: "message",
								},
							},
						},
						{
							Name: "container2",
							State: apiv1.ContainerState{
								Running: &apiv1.ContainerStateRunning{},
							},
						},
					},
				},
			},
			container: apiv1.Container{
				Name:  "container1",
				Image: "image1",
			},
			expectedToFail:  true,
			expectedMessage: "The image name \"image1\" is invalid",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result := checkInvalidImageName(tc.pod, tc.container)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
