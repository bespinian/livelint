package livelint

import (
	"context"
	"fmt"
	"strings"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) checkIsClusterFull(pods []apiv1.Pod) CheckResult {
	var podWithInsufficientResources *apiv1.Pod
	for i, pod := range pods {
		if pod.Status.Phase == apiv1.PodPending &&
			len(pod.Status.ContainerStatuses) < 1 {
			for _, condition := range pod.Status.Conditions {
				if condition.Type == apiv1.PodScheduled &&
					condition.Reason == apiv1.PodReasonUnschedulable &&
					condition.Status != apiv1.ConditionTrue &&
					strings.Contains(condition.Message, "nodes are available") &&
					(strings.Contains(condition.Message, "Insufficient cpu") ||
						strings.Contains(condition.Message, "Insufficient memory")) {
					podWithInsufficientResources = &pods[i]
					break
				}
			}
		}
		if podWithInsufficientResources != nil {
			break
		}
	}

	if podWithInsufficientResources == nil {
		return CheckResult{
			Message: "The cluster is not full",
			Details: []string{"Did not detect any pod with insufficient CPU or memory."},
		}
	}

	podCPURequests, podMemoryRequests := getPodResourceSummary(*podWithInsufficientResources)

	nodes, err := n.k8s.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return CheckResult{
			HasFailed: true,
			Message:   "Pod may have insufficient cpu or memory. Node resource check failed. Verify and resize pod requested resources or provisiona bigger cluster",
			Details:   []string{err.Error()},
		}
	}

	schedulableNodeCount := 0
	for _, node := range nodes.Items {
		if node.Spec.Unschedulable {
			continue
		}
		schedulableNodeCount++

		listOptions := metav1.ListOptions{FieldSelector: fmt.Sprintf("status.phase!=Succeeded,status.phase!=Failed,spec.nodeName=%s", node.Name)}
		nodePods, err := n.k8s.CoreV1().Pods("").List(context.Background(), listOptions)
		if err != nil {
			return CheckResult{
				HasFailed: true,
				Message:   "Pod may have insufficient cpu or memory. Node resource check failed. Verify and resize pod requested resources or provision a bigger cluster",
				Details:   []string{err.Error()},
			}
		}
		nodeTotalCPURequests := resource.NewDecimalQuantity(*resource.Zero.AsDec(), resource.DecimalSI)
		nodeTotalMemoryRequests := resource.NewDecimalQuantity(*resource.Zero.AsDec(), resource.BinarySI)
		for _, nodePod := range nodePods.Items {
			nodePodCPURequests, nodePodMemoryRequests := getPodResourceSummary(nodePod)
			nodeTotalCPURequests.Add(*nodePodCPURequests)
			nodeTotalMemoryRequests.Add(*nodePodMemoryRequests)
		}

		nodeAllocatableCPU := node.Status.Allocatable.Cpu().DeepCopy()
		nodeAllocatableCPU.Sub(*nodeTotalCPURequests)
		nodeAllocatableMemory := node.Status.Allocatable.Memory().DeepCopy()
		nodeAllocatableMemory.Sub(*nodeTotalMemoryRequests)
		if nodeAllocatableCPU.Cmp(*podCPURequests) > 0 &&
			nodeAllocatableMemory.Cmp(*podMemoryRequests) > 0 {
			return CheckResult{
				Message: "The cluster is not full",
				Details: []string{fmt.Sprintf("Found node %s with sufficient CPU and memory. (There may be other constraints on these nodes that prohibit a pod from being scheduled here.)", node.Name)},
			}
		}
	}

	return CheckResult{
		HasFailed:    true,
		Message:      "The cluster is full",
		Details:      []string{fmt.Sprintf("Checked %d schedulable nodes and found none with sufficient CPU and memory.", schedulableNodeCount)},
		Instructions: "Provision a bigger cluster",
	}
}
