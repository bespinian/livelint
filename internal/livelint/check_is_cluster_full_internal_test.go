package livelint

import (
	"context"
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedapiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func makeContainerWithResources(cpu, memory string) apiv1.Container {
	return apiv1.Container{
		Resources: apiv1.ResourceRequirements{
			Requests: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse(cpu),
				apiv1.ResourceMemory: resource.MustParse(memory),
			},
		},
	}
}

func makeNodeWithResources(cpu, memory string) apiv1.Node {
	return apiv1.Node{
		ObjectMeta: v1.ObjectMeta{Name: "NODENAME"},
		Status: apiv1.NodeStatus{
			Allocatable: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse(cpu),
				apiv1.ResourceMemory: resource.MustParse(memory),
			},
		},
	}
}

func TestCheckIsClusterFull(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it           string
		pods         []apiv1.Pod
		nodePodPairs []struct {
			node     apiv1.Node
			nodePods []apiv1.Pod
		}
		expectedToFail  bool
		expectedMessage string
		expectedDetails string
	}{
		{
			it: "succeeds if no pods are unschedulable with a message that there are insufficient compute resources",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase: apiv1.PodRunning,
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "The cluster is not full",
			expectedDetails: "Did not detect any pod with insufficient CPU or memory.",
		},
		{
			it: "succeeds if a pod is unschedulable with a message that there are insufficient compute resources, but a node with sufficient resources could be found",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase:             apiv1.PodPending,
						ContainerStatuses: []apiv1.ContainerStatus{},
						Conditions: []apiv1.PodCondition{
							{
								Type:    apiv1.PodScheduled,
								Reason:  apiv1.PodReasonUnschedulable,
								Status:  apiv1.ConditionFalse,
								Message: "0/1 nodes are available: 1 Insufficient cpu.",
							},
						},
					},
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							makeContainerWithResources("400m", "128Mi"),
						},
					},
				},
			},
			nodePodPairs: []struct {
				node     apiv1.Node
				nodePods []apiv1.Pod
			}{
				{
					node: makeNodeWithResources("2000m", "2Gi"),
					nodePods: []apiv1.Pod{
						{
							Spec: apiv1.PodSpec{
								Containers: []apiv1.Container{
									makeContainerWithResources("400m", "128Mi"),
								},
							},
						},
					},
				},
			},
			expectedToFail:  false,
			expectedMessage: "The cluster is not full",
			expectedDetails: "Found node NODENAME with sufficient CPU and memory. (There may be other constraints on these nodes that prohibit a pod from being scheduled here.)",
		},
		{
			it: "fails if a pod is unschedulable with a message that there are insufficient compute resources and it tries to schedule too much cpu using a single container",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase:             apiv1.PodPending,
						ContainerStatuses: []apiv1.ContainerStatus{},
						Conditions: []apiv1.PodCondition{
							{
								Type:    apiv1.PodScheduled,
								Reason:  apiv1.PodReasonUnschedulable,
								Status:  apiv1.ConditionFalse,
								Message: "0/1 nodes are available: 1 Insufficient cpu.",
							},
						},
					},
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							makeContainerWithResources("800m", "128Mi"),
						},
					},
				},
			},
			nodePodPairs: []struct {
				node     apiv1.Node
				nodePods []apiv1.Pod
			}{
				{
					node: makeNodeWithResources("1000m", "1Gi"),
					nodePods: []apiv1.Pod{
						{
							Spec: apiv1.PodSpec{
								Containers: []apiv1.Container{
									makeContainerWithResources("400m", "128Mi"),
								},
							},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "The cluster is full",
			expectedDetails: "Checked 1 schedulable nodes and found none with sufficient CPU and memory.",
		},
		{
			it: "fails if a pod is unschedulable with a message that there are insufficient compute resources and it tries to schedule too much cpu using multiple containers",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase:             apiv1.PodPending,
						ContainerStatuses: []apiv1.ContainerStatus{},
						Conditions: []apiv1.PodCondition{
							{
								Type:    apiv1.PodScheduled,
								Reason:  apiv1.PodReasonUnschedulable,
								Status:  apiv1.ConditionFalse,
								Message: "0/1 nodes are available: 1 Insufficient cpu.",
							},
						},
					},
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							makeContainerWithResources("400m", "128Mi"),
							makeContainerWithResources("400m", "128Mi"),
						},
					},
				},
			},
			nodePodPairs: []struct {
				node     apiv1.Node
				nodePods []apiv1.Pod
			}{
				{
					node: makeNodeWithResources("1000m", "1Gi"),
					nodePods: []apiv1.Pod{
						{
							Spec: apiv1.PodSpec{
								Containers: []apiv1.Container{
									makeContainerWithResources("100m", "128Mi"),
									makeContainerWithResources("100m", "128Mi"),
									makeContainerWithResources("100m", "128Mi"),
								},
							},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "The cluster is full",
			expectedDetails: "Checked 1 schedulable nodes and found none with sufficient CPU and memory.",
		},
		{
			it: "fails if a pod is unschedulable with a message that there are insufficient compute resources and it tries to schedule too much memory using a single container",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase:             apiv1.PodPending,
						ContainerStatuses: []apiv1.ContainerStatus{},
						Conditions: []apiv1.PodCondition{
							{
								Type:    apiv1.PodScheduled,
								Reason:  apiv1.PodReasonUnschedulable,
								Status:  apiv1.ConditionFalse,
								Message: "0/1 nodes are available: 1 Insufficient memory.",
							},
						},
					},
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							makeContainerWithResources("100m", "1024Mi"),
						},
					},
				},
			},
			nodePodPairs: []struct {
				node     apiv1.Node
				nodePods []apiv1.Pod
			}{
				{
					node: makeNodeWithResources("1000m", "1Gi"),
					nodePods: []apiv1.Pod{
						{
							Spec: apiv1.PodSpec{
								Containers: []apiv1.Container{
									makeContainerWithResources("400m", "128Mi"),
								},
							},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "The cluster is full",
			expectedDetails: "Checked 1 schedulable nodes and found none with sufficient CPU and memory.",
		},
		{
			it: "fails if a pod is unschedulable with a message that there are insufficient compute resources and it tries to schedule too much memory using a single container",
			pods: []apiv1.Pod{
				{
					Status: apiv1.PodStatus{
						Phase:             apiv1.PodPending,
						ContainerStatuses: []apiv1.ContainerStatus{},
						Conditions: []apiv1.PodCondition{
							{
								Type:    apiv1.PodScheduled,
								Reason:  apiv1.PodReasonUnschedulable,
								Status:  apiv1.ConditionFalse,
								Message: "0/1 nodes are available: 1 Insufficient memory.",
							},
						},
					},
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							makeContainerWithResources("100m", "512Mi"),
							makeContainerWithResources("100m", "512Mi"),
						},
					},
				},
			},
			nodePodPairs: []struct {
				node     apiv1.Node
				nodePods []apiv1.Pod
			}{
				{
					node: makeNodeWithResources("1000m", "1Gi"),
					nodePods: []apiv1.Pod{
						{
							Spec: apiv1.PodSpec{
								Containers: []apiv1.Container{
									makeContainerWithResources("100m", "128Mi"),
									makeContainerWithResources("100m", "128Mi"),
									makeContainerWithResources("100m", "128Mi"),
								},
							},
						},
					},
				},
			},
			expectedToFail:  true,
			expectedMessage: "The cluster is full",
			expectedDetails: "Checked 1 schedulable nodes and found none with sufficient CPU and memory.",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			nodeCount := 0

			k8s := &kubernetesInterfaceMock{
				CoreV1Func: func() typedapiv1.CoreV1Interface {
					return &apiv1InterfaceMock{
						NodesFunc: func() typedapiv1.NodeInterface {
							return &apiv1NodeInterfaceMock{
								ListFunc: func(ctx context.Context, opts v1.ListOptions) (*apiv1.NodeList, error) {
									nodes := []apiv1.Node{}
									for _, nodePodPair := range tc.nodePodPairs {
										nodes = append(nodes, nodePodPair.node)
									}
									return &apiv1.NodeList{Items: nodes}, nil
								},
							}
						},
						PodsFunc: func(namespace string) typedapiv1.PodInterface {
							return &apiv1PodInterfaceMock{
								ListFunc: func(ctx context.Context, opts v1.ListOptions) (*apiv1.PodList, error) {
									nodeCount++
									return &apiv1.PodList{Items: tc.nodePodPairs[nodeCount-1].nodePods}, nil
								},
							}
						},
					}
				},
			}
			ll := Livelint{
				k8s: k8s,
			}

			result := ll.checkIsClusterFull(tc.pods)

			is.Equal(result.HasFailed, tc.expectedToFail)   // HasFailed
			is.Equal(result.Message, tc.expectedMessage)    // Message
			is.Equal(result.Details[0], tc.expectedDetails) // Details (first element)
		})
	}
}
