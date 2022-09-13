package livelint

import (
	"fmt"
	"strings"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) checkIsImageTagValid(container apiv1.Container) CheckResult {
	parts := strings.Split(container.Image, ":")

	image := parts[0]
	tag := "latest"
	if len(parts) > 1 {
		tag = parts[1]
	}

	yes := n.askUserYesOrNo(fmt.Sprintf("Is the tag %q for the image %q valid? Does it exist?", tag, image))

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("The tag %s for the image %s is not valid or doesn't exist", tag, image),
			Instructions: fmt.Sprintf("Fix the incorrect tag %s", tag),
		}
	}

	return CheckResult{
		Message: "The image tag is valid and exists",
	}
}
