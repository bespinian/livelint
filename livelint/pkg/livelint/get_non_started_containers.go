package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

// getNonStartedContainers returns all containers from a pod that are not in status "Running".
func (n *livelint) getNonStartedContainers(pod corev1.Pod) ([]corev1.Container) {
	nonStartedContainers := []corev1.Container{}
	containerStatuses := pod.Status.ContainerStatuses
	for i := 0; i < len(containerStatuses); i++ {
		status := containerStatuses[i]
		if status.State.Running == nil {
			for _, container := range pod.Spec.Containers {
				if (container.Name == status.Name) {
					nonStartedContainers = append(nonStartedContainers, container)
					break
				}
			}
		}
	}

	return nonStartedContainers
}