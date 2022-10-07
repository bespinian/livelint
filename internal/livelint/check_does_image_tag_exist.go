package livelint

import (
	"fmt"
	"strings"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) checkDoesImageTagExist(container apiv1.Container) CheckResult {
	parts := strings.Split(container.Image, ":")

	image := parts[0]
	tag := "latest"
	if len(parts) > 1 {
		tag = parts[1]
	}

	yes := n.askUserYesOrNo(fmt.Sprintf("Does the tag %q for the image %q exist in the container repository?", tag, image))

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("The tag %s for the image %s does not exist", tag, image),
			Instructions: fmt.Sprintf("Fix the configured tag %s of the image %s to the correct value", tag, image),
		}
	}

	return CheckResult{
		Message: "The image tag exists",
	}
}
