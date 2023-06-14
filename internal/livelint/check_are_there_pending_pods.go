package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

// checkAreTherePendingPods checks if there are pending pods with no further container status information.
func checkAreTherePendingPods(pods []apiv1.Pod) CheckResult {
	pendingPodCount := 0
	details := []string{}
	for _, pod := range pods {
		if pod.Status.Phase == apiv1.PodPending &&
			len(pod.Status.ContainerStatuses) == 0 {
			pendingPodCount++
			details = append(details, fmt.Sprintf("Pod %q is pending", pod.Name))
		}
	}

	if pendingPodCount > 0 {
		msgTemplate := "There are %v pending Pods"
		if pendingPodCount == 1 {
			msgTemplate = "There is %v pending Pod"
		}

		return CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf(msgTemplate, pendingPodCount),
			Details:   details,
		}
	}

	return CheckResult{
		Message: "There are no pending Pods",
	}
}
