package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

// getProblematicContainers returns all containers from a pod that
// are not in status "Running" and that are also not init containers
// that terminated successfully.
func (n *livelint) getProblematicContainers(pod corev1.Pod) []corev1.Container {
	nonStartedContainers := []corev1.Container{}

	for _, status := range pod.Status.InitContainerStatuses {
		if status.State.Running == nil &&
			status.State.Terminated != nil &&
			status.State.Terminated.ExitCode != 0 {
			for _, container := range pod.Spec.InitContainers {
				if container.Name == status.Name {
					nonStartedContainers = append(nonStartedContainers, container)
				}
			}
		}
	}

	for _, status := range pod.Status.ContainerStatuses {
		if status.State.Running == nil {
			for _, container := range pod.Spec.Containers {
				if container.Name == status.Name {
					nonStartedContainers = append(nonStartedContainers, container)
				}
			}
		}
	}

	return nonStartedContainers
}
