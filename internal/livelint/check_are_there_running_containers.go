package livelint

import (
	apiv1 "k8s.io/api/core/v1"
)

func checkAreThereRunningContainers(pod apiv1.Pod) CheckResult {
	for _, pod := range append(pod.Status.ContainerStatuses, pod.Status.InitContainerStatuses...) {
		if pod.State.Running != nil {
			return CheckResult{
				Message:      "There are RUNNING containers",
				Instructions: "The issue is with the node-lifecycle controller",
			}
		}
	}

	return CheckResult{
		HasFailed:    true,
		Message:      "There are no RUNNING containers",
		Instructions: "Give it some more time or consult StackOverflow",
	}
}
