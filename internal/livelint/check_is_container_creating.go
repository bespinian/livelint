package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func checkIsContainerCreating(pod apiv1.Pod) CheckResult {
	creatingContainerCount := 0
	details := []string{}
	if pod.Status.Phase == apiv1.PodPending {
		containerStatuses := append(pod.Status.ContainerStatuses, pod.Status.InitContainerStatuses...)
		for _, containerStatus := range containerStatuses {
			if containerStatus.State.Waiting != nil && containerStatus.State.Waiting.Reason == "ContainerCreating" {
				creatingContainerCount++
				details = append(details, fmt.Sprintf("Container %q of Pod %q is in state ContainerCreating %s", containerStatus.Name, pod.Name, containerStatus.State.Waiting.Message))
			}
		}
	}

	if creatingContainerCount > 0 {
		msgTemplate := "There are %v containers still being created"
		if creatingContainerCount == 1 {
			msgTemplate = "There is %v container still being created"
		}

		return CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf(msgTemplate, creatingContainerCount),
			Details:   details,
		}
	}

	return CheckResult{
		Message: "No Container is still being created",
	}
}
