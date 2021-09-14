package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

// areAllPodsRunning checks if all Pods are RUNNING.
func areAllPodsRunning(allPods []corev1.Pod) bool {
	for i := 0; i < len(allPods); i++ {
		if allPods[i].Status.Phase != corev1.PodRunning {
			return false
		}
	}
	return true
}
