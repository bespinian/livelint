package livelint_test

import (
	"context"
	"testing"

	"github.com/bespinian/livelint/internal/livelint"
	"github.com/matryer/is"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func TestCheckAreResourceQuotasHit(t *testing.T) {
	t.Parallel()

	one := int32(1)

	cases := []struct {
		it              string
		replicaSets     []appsv1.ReplicaSet
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds if quotas are not hit",
			replicaSets: []appsv1.ReplicaSet{
				{
					ObjectMeta: metav1.ObjectMeta{
						OwnerReferences: []metav1.OwnerReference{
							{
								Kind: "Deployment",
								Name: "DEPLOYMENT",
							},
						},
					},
					Spec: appsv1.ReplicaSetSpec{
						Replicas: &one,
					},
					Status: appsv1.ReplicaSetStatus{
						AvailableReplicas: 1,
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "You are well within the ResourceQuota limits",
		},
		{
			it:              "succeeds if there are no replica sets",
			replicaSets:     []appsv1.ReplicaSet{},
			expectedToFail:  false,
			expectedMessage: "You are well within the ResourceQuota limits",
		},
		{
			it: "succeeds if there are no replica sets for deployment",
			replicaSets: []appsv1.ReplicaSet{
				{
					ObjectMeta: metav1.ObjectMeta{
						OwnerReferences: []metav1.OwnerReference{
							{
								Kind: "Deployment",
								Name: "other",
							},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "You are well within the ResourceQuota limits",
		},
		{
			it: "fails if quotas are exceeded",
			replicaSets: []appsv1.ReplicaSet{
				{
					ObjectMeta: metav1.ObjectMeta{
						OwnerReferences: []metav1.OwnerReference{
							{
								Kind: "Deployment",
								Name: "DEPLOYMENT",
							},
						},
					},
					Spec: appsv1.ReplicaSetSpec{
						Replicas: &one,
					},
					Status: appsv1.ReplicaSetStatus{
						AvailableReplicas: 0,
						Conditions: []appsv1.ReplicaSetCondition{
							{
								Type:    appsv1.ReplicaSetReplicaFailure,
								Message: "exceeded quota:",
							},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "You are hitting the ResourceQuota limits",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			k8s := &KubernetesInterfaceMock{
				AppsV1Func: func() typedappsv1.AppsV1Interface {
					return &Appsv1InterfaceMock{
						ReplicaSetsFunc: func(string) typedappsv1.ReplicaSetInterface {
							return &ReplicaSetInterfaceMock{
								ListFunc: func(context.Context, metav1.ListOptions) (*appsv1.ReplicaSetList, error) {
									return &appsv1.ReplicaSetList{Items: tc.replicaSets}, nil
								},
							}
						},
					}
				},
			}
			ll := livelint.Livelint{
				K8s: k8s,
			}
			result := ll.CheckAreResourceQuotasHit("NAMESPACE", "DEPLOYMENT")

			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			is.Equal(result.Message, tc.expectedMessage)  // Message
		})
	}
}
