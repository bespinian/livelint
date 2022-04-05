package livelint

import (
	"context"
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedapiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func TestCheckIsMountingPendingPVC(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it              string
		pods            []apiv1.Pod
		pvcs            []apiv1.PersistentVolumeClaim
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if no PVCs are pending",
			pods: []apiv1.Pod{
				{
					Spec: apiv1.PodSpec{
						Volumes: []apiv1.Volume{
							{
								VolumeSource: apiv1.VolumeSource{
									PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
										ClaimName: "PVC",
									},
								},
							},
						},
					},
				},
			},
			pvcs: []apiv1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "PVC",
					},
					Status: apiv1.PersistentVolumeClaimStatus{
						Phase: apiv1.ClaimBound,
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "You are not mounting any PENDING PersistentVolumeClaims",
		},
		{
			it:              "succeeds if there are no pods",
			pods:            []apiv1.Pod{},
			expectedToFail:  false,
			expectedMessage: "You are not mounting any PENDING PersistentVolumeClaims",
		},
		{
			it: "succeeds if pod has no volumes",
			pods: []apiv1.Pod{
				{
					Spec: apiv1.PodSpec{
						Volumes: []apiv1.Volume{},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "You are not mounting any PENDING PersistentVolumeClaims",
		},
		{
			it: "fails if there are pending PVCs",
			pods: []apiv1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "POD",
					},
					Spec: apiv1.PodSpec{
						Volumes: []apiv1.Volume{
							{
								VolumeSource: apiv1.VolumeSource{
									PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
										ClaimName: "PVC",
									},
								},
							},
						},
					},
				},
			},
			pvcs: []apiv1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "PVC",
					},
					Status: apiv1.PersistentVolumeClaimStatus{
						Phase: apiv1.ClaimPending,
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "You are mounting a PENDING PersistentVolumeClaim PVC in Pod POD",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			k8s := &kubernetesInterfaceMock{
				CoreV1Func: func() typedapiv1.CoreV1Interface {
					return &apiv1InterfaceMock{
						PersistentVolumeClaimsFunc: func(string) typedapiv1.PersistentVolumeClaimInterface {
							return &apiv1PVCInterfaceMock{
								ListFunc: func(context.Context, metav1.ListOptions) (*apiv1.PersistentVolumeClaimList, error) {
									return &apiv1.PersistentVolumeClaimList{Items: tc.pvcs}, nil
								},
							}
						},
					}
				},
			}
			ll := Livelint{
				k8s: k8s,
			}
			result := ll.checkIsMountingPendingPVC(tc.pods, "NAMESPACE")

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
