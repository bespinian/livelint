package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func checkImagePullErrors(pod apiv1.Pod, containerName string) CheckResult {
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Name != containerName {
			continue
		}

		if containerStatus.State.Waiting != nil &&
			(containerStatus.State.Waiting.Reason == "ErrImagePull" ||
				containerStatus.State.Waiting.Reason == "ImagePullBackOff") {
			return CheckResult{
				HasFailed: true,
				Message:   fmt.Sprintf("The Pod is in status %s", containerStatus.State.Waiting.Reason),
				Details:   []string{containerStatus.State.Waiting.Message},
			}
		}
	}

	return CheckResult{
		Message: "The Pod is not in status ImagePullBackOff",
	}
}
