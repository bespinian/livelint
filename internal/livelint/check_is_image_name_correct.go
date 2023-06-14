package livelint

import (
	"fmt"
	"strings"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) checkIsImageNameCorrect(container apiv1.Container) CheckResult {
	parts := strings.Split(container.Image, ":")

	image := parts[0]

	yes := n.ui.AskYesNo(fmt.Sprintf("Is image name %q correct for the %q container?", image, container.Name))

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("The image name %q is not correct", image),
			Details:      []string{fmt.Sprintf("Container %q uses incorrect image name %q", container.Name, image)},
			Instructions: "Fix the image name",
		}
	}

	return CheckResult{
		Message: "The image name is correct",
	}
}
