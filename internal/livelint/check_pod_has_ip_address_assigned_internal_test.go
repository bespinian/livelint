package livelint

import (
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCheckPodHasIPAddressAssigned(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pods            []apiv1.Pod
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds all pods have an IP assigned",
			pods: []apiv1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod1"},
					Status: apiv1.PodStatus{
						PodIP: "1.3.3.7",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod2"},
					Status: apiv1.PodStatus{
						PodIP: "1.3.3.7",
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "The Pods have an IP address assigned",
		},
		{
			it: "fails if at least one pod does not have an IP assigned",
			pods: []apiv1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod1"},
					Status: apiv1.PodStatus{
						PodIP: "1.3.3.7",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod2"},
					Status:     apiv1.PodStatus{},
				},
			},
			expectedToFail:  true,
			expectedMessage: "The Pod pod2 has no IP address assigned",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result := checkPodHasIPAddressAssigned(tc.pods)

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
