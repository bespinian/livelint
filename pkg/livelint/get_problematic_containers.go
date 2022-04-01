package livelint

import (
	apiv1 "k8s.io/api/core/v1"
)

// getNonRunningContainers returns all containers from a pod that
// are not in status RUNNING and that are also not init containers
// that terminated successfully.
func getNonRunningContainers(pod apiv1.Pod) []apiv1.Container {
	nonStartedContainers := []apiv1.Container{}

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
