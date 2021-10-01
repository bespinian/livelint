package livelint

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// checkAreTherePendingPods checks if there are PENDING pods with no further container status information.
func checkAreTherePendingPods(allPods []corev1.Pod) CheckResult {
	pendingPods := []corev1.Pod{}
	for _, pod := range allPods {
		if pod.Status.Phase == corev1.PodPending &&
			len(pod.Status.ContainerStatuses) < 1 {
			pendingPods = append(pendingPods, pod)
		}
	}

	if len(pendingPods) > 0 {
		pendingPodNames := make([]string, 0, len(pendingPods))
		for _, pod := range pendingPods {
			pendingPodNames = append(pendingPodNames, pod.ObjectMeta.Name)
		}

		return CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf("There are %d PENDING Pods", len(pendingPods)),
			Details:   pendingPodNames,
		}
	}

	return CheckResult{
		Message: "There are no PENDING Pods",
	}
}
