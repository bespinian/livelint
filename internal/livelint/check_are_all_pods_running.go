package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

// checkAreAllPodsRunning checks if all Pods and all their containers are running.
func checkAreAllPodsRunning(pods []apiv1.Pod) CheckResult {
	notRunningPodCount := 0
	details := []string{}
	for _, pod := range pods {
		state := ""
		if pod.Status.Phase != apiv1.PodRunning {
			state = string(pod.Status.Phase)
		}
		for _, cs := range pod.Status.ContainerStatuses {
			if cs.State.Running == nil {
				if cs.State.Waiting != nil {
					state = cs.State.Waiting.Reason
					break
				}
			}
		}
		if state != "" {
			notRunningPodCount++
			details = append(details, fmt.Sprintf("Pod %q is in status %s", pod.Name, state))
		}
	}

	if notRunningPodCount > 0 {
		msgTemplate := "There are %v Pods that are not running"
		if notRunningPodCount == 1 {
			msgTemplate = "There is %v Pod that is not running"
		}

		return CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf(msgTemplate, notRunningPodCount),
			Details:   details,
		}
	}

	return CheckResult{
		Message: "All Pods are running",
	}
}
