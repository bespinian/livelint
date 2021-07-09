package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

// getNonStartedContainerNames returns all containers from a pod that are not in status "Started".
func (n *livelint) getNonStartedContainerNames(pod corev1.Pod) ([]string) {
	nonStartedContainers := []string{}
	if pod.Status.Phase == corev1.PodRunning {
		containerStatuses := pod.Status.ContainerStatuses
		for j := 0; j < len(containerStatuses); j++ {
			status := containerStatuses[j]
			if status.State.Running == nil {
				nonStartedContainers = append(nonStartedContainers, status.Name)
			}
		}
	}

	return nonStartedContainers
}