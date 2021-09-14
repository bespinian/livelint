package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

// areTherePendingPods checks if there are PENDING pods.
func areTherePendingPods(allPods []corev1.Pod) bool {
	for i := 0; i < len(allPods); i++ {
		if allPods[i].Status.Phase == corev1.PodPending {
			return true
		}
	}
	return false
}
