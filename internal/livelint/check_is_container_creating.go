package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func checkIsContainerCreating(pod apiv1.Pod) CheckResult {
	if pod.Status.Phase == apiv1.PodPending {
		for _, containerStatus := range append(pod.Status.ContainerStatuses, pod.Status.InitContainerStatuses...) {
			if containerStatus.State.Waiting != nil && containerStatus.State.Waiting.Reason == "ContainerCreating" {
				return CheckResult{
					HasFailed: true,
					Message:   fmt.Sprintf("Container %s of pod %s is in state ContainerCreating", containerStatus.Name, pod.Name),
				}
			}
		}
	}

	return CheckResult{
		Message: "No Container is in state ContainerCreating",
	}
}
