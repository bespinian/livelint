package livelint

import (
	apiv1 "k8s.io/api/core/v1"
)

func checkAreThereRunningContainers(pod apiv1.Pod) CheckResult {
	for _, pod := range append(pod.Status.ContainerStatuses, pod.Status.InitContainerStatuses...) {
		if pod.State.Running != nil {
			return CheckResult{
				Message:      "There are containers RUNNING",
				Instructions: "The issue is with the node-lifecycle controller",
			}
		}
	}

	return CheckResult{
		HasFailed:    true,
		Message:      "There are no containers RUNNING",
		Instructions: "Consult StackOverflow",
	}
}
