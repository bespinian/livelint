package livelint

import (
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
)

func TestCheckAreAllPodsReady(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pods            []apiv1.Pod
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if all pods are ready",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Conditions: []apiv1.PodCondition{
							{
								Type:   apiv1.ContainersReady,
								Status: apiv1.ConditionTrue,
							},
							{
								Type:   apiv1.ContainersReady,
								Status: apiv1.ConditionTrue,
							},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "All Pods are READY",
		},
		{
			it:              "succeeds if there are no pods",
			pods:            []apiv1.Pod{},
			expectedToFail:  false,
			expectedMessage: "All Pods are READY",
		},
		{
			it: "fails if there are non-ready pods",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Conditions: []apiv1.PodCondition{
							{
								Type:   apiv1.ContainersReady,
								Status: apiv1.ConditionFalse,
							},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "There is 1 Pod that is not READY",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result := checkAreAllPodsReady(tc.pods)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
