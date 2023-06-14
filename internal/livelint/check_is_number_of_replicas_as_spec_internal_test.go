package livelint_test

import (
	"context"
	"testing"

	"github.com/bespinian/livelint/internal/livelint"
	"github.com/matryer/is"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func TestIsNumberOfReplicasAsCorrect(t *testing.T) {
	t.Parallel()

	two := int32(2)

	cases := []struct {
		it              string
		deployment      appsv1.Deployment
		replicaSets     []appsv1.ReplicaSet
		expectedToFail  bool
		expectedMessage string
	}{
		{
			it: "succeeds, if the number of replicas match the number in the deployment spec",
			deployment: appsv1.Deployment{
				TypeMeta:   v1.TypeMeta{},
				ObjectMeta: v1.ObjectMeta{Name: "DEPLOYMENT"},
				Spec: appsv1.DeploymentSpec{
					Replicas: &two,
				},
				Status: appsv1.DeploymentStatus{
					Replicas: 2,
				},
			},
			replicaSets: []appsv1.ReplicaSet{
				{
					ObjectMeta: v1.ObjectMeta{
						OwnerReferences: []v1.OwnerReference{
							{
								Kind: "Deployment",
								Name: "DEPLOYMENT",
							},
						},
					},
					Spec: appsv1.ReplicaSetSpec{
						Replicas: &two,
					},
					Status: appsv1.ReplicaSetStatus{
						Replicas: 2,
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "Number of replicas is as desired",
		},
		{
			it: "fails, if the number of replicas is lower than the number of replicas in the deployment spec",
			deployment: appsv1.Deployment{
				TypeMeta:   v1.TypeMeta{},
				ObjectMeta: v1.ObjectMeta{Name: "DEPLOYMENT"},
				Spec: appsv1.DeploymentSpec{
					Replicas: &two,
				},
				Status: appsv1.DeploymentStatus{
					Replicas: 2,
				},
			},
			replicaSets: []appsv1.ReplicaSet{
				{
					ObjectMeta: v1.ObjectMeta{
						OwnerReferences: []v1.OwnerReference{
							{
								Kind: "Deployment",
								Name: "DEPLOYMENT",
							},
						},
					},
					Spec: appsv1.ReplicaSetSpec{
						Replicas: &two,
					},
					Status: appsv1.ReplicaSetStatus{
						Replicas: 1,
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "Number of replicas is lower than desired",
		},
		{
			it: "succeeds, if the number of replicas is higher than the one in the deployment spec",
			deployment: appsv1.Deployment{
				TypeMeta:   v1.TypeMeta{},
				ObjectMeta: v1.ObjectMeta{Name: "DEPLOYMENT"},
				Spec: appsv1.DeploymentSpec{
					Replicas: &two,
				},
				Status: appsv1.DeploymentStatus{
					Replicas: 2,
				},
			},
			replicaSets: []appsv1.ReplicaSet{
				{
					ObjectMeta: v1.ObjectMeta{
						OwnerReferences: []v1.OwnerReference{
							{
								Kind: "Deployment",
								Name: "DEPLOYMENT",
							},
						},
					},
					Spec: appsv1.ReplicaSetSpec{
						Replicas: &two,
					},
					Status: appsv1.ReplicaSetStatus{
						Replicas: 3,
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "Cluster in intermediary state. Number of replicas is larger than desired. Re-run livelint once cluster is in stable state.",
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
						DeploymentsFunc: func(namespace string) typedappsv1.DeploymentInterface {
							return &DeploymentInterfaceMock{
								GetFunc: func(ctx context.Context, name string, opts v1.GetOptions) (*appsv1.Deployment, error) {
									return &tc.deployment, nil
								},
							}
						},
						ReplicaSetsFunc: func(namespace string) typedappsv1.ReplicaSetInterface {
							return &ReplicaSetInterfaceMock{
								ListFunc: func(ctx context.Context, opts v1.ListOptions) (*appsv1.ReplicaSetList, error) {
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
			result := ll.CheckIsNumberOfReplicasCorrect("NAMESPACE", "DEPLOYMENT")

			is.Equal(result.Message, tc.expectedMessage)  // Message
			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
		})
	}
}
