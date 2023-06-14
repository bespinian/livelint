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

	yes := n.ui.AskYesNo(fmt.Sprintf("Does the tag %q for the image %q exist in the container repository?", tag, image))
	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("The tag %q for the image %q does not exist", tag, image),
			Details:      []string{fmt.Sprintf("Container %q with image %q uses incorrect tag %q", container.Name, image, tag)},
			Instructions: "Fix the tag",
		}
	}

	return CheckResult{
		Message: "The image tag exists",
	}
}
