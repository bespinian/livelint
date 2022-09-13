package livelint

import (
	"fmt"
	"strings"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) checkIsImageNameCorrect(container apiv1.Container) CheckResult {
	parts := strings.Split(container.Image, ":")

	image := parts[0]

	yes := n.askUserYesOrNo(fmt.Sprintf("Is the name of the image %q correct for the container %q?", image, container.Name))

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("The name of the image %s is not correct for the container %s", image, container.Name),
			Instructions: fmt.Sprintf("Fix the image name %s", image),
		}
	}

	return CheckResult{
		Message: "The name of the image is correct",
	}
}
