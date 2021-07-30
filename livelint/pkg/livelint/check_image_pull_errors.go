package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) checkImagePullErrors(pod corev1.Pod, containerName string) (bool, string, string) {
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Name != containerName {
			continue
		}

		if containerStatus.State.Waiting != nil &&
			containerStatus.State.Waiting.Reason == "ErrImagePull" ||
			containerStatus.State.Waiting.Reason == "ImagePullBackOff" {
			return true, containerStatus.State.Waiting.Reason,
				containerStatus.State.Waiting.Message
		}
	}
	return false, "", ""
}
