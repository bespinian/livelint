package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

// areTherePendingPods checks if there are PENDING pods with no further container status information.
func areTherePendingPods(allPods []corev1.Pod) bool {
	for _, pod := range allPods {
		if pod.Status.Phase == corev1.PodPending &&
			len(pod.Status.ContainerStatuses) < 1 {
			return true
		}
	}
	return false
}
