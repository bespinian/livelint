package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

// areAllPodsRunning checks if all Pods are in phase Running and all their containers are in state Running.
func areAllPodsRunning(allPods []corev1.Pod) bool {
	for _, pod := range allPods {
		if pod.Status.Phase != corev1.PodRunning {
			return false
		}
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.State.Running == nil {
				return false
			}
		}
	}
	return true
}
