package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func checkInvalidImageName(pod apiv1.Pod, container apiv1.Container) CheckResult {
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Name != container.Name {
			continue
		}

		if containerStatus.State.Waiting != nil &&
			containerStatus.State.Waiting.Reason == "InvalidImageName" {
			return CheckResult{
				HasFailed:    true,
				Message:      fmt.Sprintf("The image name %q is invalid", container.Image),
				Details:      []string{containerStatus.State.Waiting.Message},
				Instructions: "Fix the image",
			}
		}
	}

	return CheckResult{
		Message: "All image names are valid",
	}
}
