package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) checkIsImageNameCorrect(container apiv1.Container) CheckResult {
	yes := n.askUserYesOrNo(fmt.Sprintf("Is the name of the image %q correct for the container %q?", container.Image, container.Name))

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("The name of the image %q is not correct for the container %q", container.Image, container.Name),
			Instructions: "Fix the image name",
		}
	}

	return CheckResult{
		Message: "The name of the image is correct",
	}
}
