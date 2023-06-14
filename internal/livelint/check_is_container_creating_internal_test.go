package livelint

import (
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCheckIsContainerCreating(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pod             apiv1.Pod
		services        []apiv1.Service
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "returns a failing result if the pod has a container with status waiting and reason ContainerCreating",
			pod: apiv1.Pod{
				ObjectMeta: v1.ObjectMeta{Name: "TESTPOD"},
				Status: apiv1.PodStatus{
					Phase: apiv1.PodPending,
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "TESTCONTAINER",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "ContainerCreating"}},
						},
					},
				},
			},

			expectedToFail:  true,
			expectedMessage: "There is 1 container still being created",
		},
		{
			it: "returns a non failing result if the pod has a container with status waiting and reason ContainerCreating but is not pending",
			pod: apiv1.Pod{
				ObjectMeta: v1.ObjectMeta{Name: "TESTPOD"},
				Status: apiv1.PodStatus{
					Phase: apiv1.PodRunning,
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "TESTCONTAINER",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "ContainerCreating"}},
						},
					},
				},
			},

			expectedToFail:  false,
			expectedMessage: "No Container is still being created",
		},
		{
			it: "returns a failing result if the pod has an init container with status waiting and reason ContainerCreating",
			pod: apiv1.Pod{
				ObjectMeta: v1.ObjectMeta{Name: "TESTPOD"},
				Status: apiv1.PodStatus{
					Phase: apiv1.PodPending,
					InitContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "TESTCONTAINER",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "ContainerCreating"}},
						},
					},
				},
			},

			expectedToFail:  true,
			expectedMessage: "There is 1 container still being created",
		},
		{
			it: "returns a failing result if the pod has one of many containers with status waiting and reason ContainerCreating",
			pod: apiv1.Pod{
				ObjectMeta: v1.ObjectMeta{Name: "TESTPOD"},
				Status: apiv1.PodStatus{
					Phase: apiv1.PodPending,
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "ALLGOOD",
							State: apiv1.ContainerState{Running: &apiv1.ContainerStateRunning{}},
						},
						{
							Name:  "TESTCONTAINER",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "ContainerCreating"}},
						},
						{
							Name:  "ALLGOOD",
							State: apiv1.ContainerState{Running: &apiv1.ContainerStateRunning{}},
						},
					},
				},
			},

			expectedToFail:  true,
			expectedMessage: "There is 1 container still being created",
		},
		{
			it: "returns a non-failing result if the pod has a container with status waiting and reason != ContainerCreating",
			pod: apiv1.Pod{
				ObjectMeta: v1.ObjectMeta{Name: "TESTPOD"},
				Status: apiv1.PodStatus{
					ContainerStatuses: []apiv1.ContainerStatus{
						{
							Name:  "TESTCONTAINER",
							State: apiv1.ContainerState{Waiting: &apiv1.ContainerStateWaiting{Reason: "SOMETHING"}},
						},
					},
				},
			},

			expectedToFail:  false,
			expectedMessage: "No Container is still being created",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result := checkIsContainerCreating(tc.pod)

			is.Equal(result.Message, tc.expectedMessage)  // Message
			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
		})
	}
}
