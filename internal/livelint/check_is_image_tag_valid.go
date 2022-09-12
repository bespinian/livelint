package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) checkIsImageTagValid(container apiv1.Container) CheckResult {
	yes := n.askUserYesOrNo(fmt.Sprintf("Is the image tag %q for the container %q valid? Does it exist?", container.Image, container.Name))

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("The image tag %q for the container %q is not valid or doesn't exist", container.Image, container.Name),
			Instructions: "Fix the tag",
		}
	}

	return CheckResult{
		Message: "The image tag is valid and exists",
	}
}
