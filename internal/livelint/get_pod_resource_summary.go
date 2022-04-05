package livelint

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// getPodResourceSummary returns summary of a pods cpu and memory requests.
// Returns (cpu requests, memory requests).
func getPodResourceSummary(pod apiv1.Pod) (*resource.Quantity, *resource.Quantity) {
	podTotalCPURequests := resource.NewDecimalQuantity(*resource.Zero.AsDec(), resource.DecimalSI)
	podTotalMemoryRequests := resource.NewDecimalQuantity(*resource.Zero.AsDec(), resource.BinarySI)
	for _, container := range pod.Spec.Containers {
		cpuRequests := container.Resources.Requests.Cpu()
		if cpuRequests != nil {
			podTotalCPURequests.Add((*cpuRequests))
		}
		memoryRequests := container.Resources.Requests.Memory()
		if memoryRequests != nil {
			podTotalMemoryRequests.Add((*memoryRequests))
		}
	}
	return podTotalCPURequests, podTotalMemoryRequests
}
