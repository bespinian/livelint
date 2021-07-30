package livelint

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) checkContainerLogs(pod corev1.Pod, containerName string) (*string, error) {
	namespace := pod.Namespace
	logs, err := n.tailPodLogs(namespace, pod.Name, containerName, 20, false)
	if err != nil {
		logs, err = n.tailPodLogs(namespace, pod.Name, containerName, 20, true)
		if err != nil {
			return nil, fmt.Errorf("error getting logs: %w", err)
		}
	}

	return &logs, nil
}
