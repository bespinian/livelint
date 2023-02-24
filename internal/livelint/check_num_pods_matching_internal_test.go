package livelint

import (
	"context"
	"testing"

	"github.com/matryer/is"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func TestCheckNumberOfPods(t *testing.T) {
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
			it: "succeeds, if the number of pods match the number of replicas in the deployment spec",
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
			expectedMessage: "Desired number of pods is running",
		},
		{
			it: "fails, if the number of pods is lower than the number of replicas in the deployment spec",
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
			expectedMessage: "Number of pods is lower then expected",
		},
		{
			it: "succeeds, if the number of pods is higher than the number of replicas in the deployment spec",
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
			expectedMessage: "Your cluster is in intermediary state. Number of pods is bigger then expected. Rerun livelint once cluster is in stable state.",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			k8s := &kubernetesInterfaceMock{
				AppsV1Func: func() typedappsv1.AppsV1Interface {
					return &appsv1InterfaceMock{
						DeploymentsFunc: func(namespace string) typedappsv1.DeploymentInterface {
							return &deploymentInterfaceMock{
								GetFunc: func(ctx context.Context, name string, opts v1.GetOptions) (*appsv1.Deployment, error) {
									return &tc.deployment, nil
								},
							}
						},
						ReplicaSetsFunc: func(namespace string) typedappsv1.ReplicaSetInterface {
							return &replicaSetInterfaceMock{
								ListFunc: func(ctx context.Context, opts v1.ListOptions) (*appsv1.ReplicaSetList, error) {
									return &appsv1.ReplicaSetList{Items: tc.replicaSets}, nil
								},
							}
						},
					}
				},
			}

			ll := Livelint{
				k8s: k8s,
			}
			result := ll.CheckIsNumberOfPodsMatching("NAMESPACE", "DEPLOYMENT")

			is.Equal(result.Message, tc.expectedMessage)  // Message
			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
		})
	}
}
