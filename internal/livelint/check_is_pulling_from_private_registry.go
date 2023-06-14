package livelint

import "fmt"

func (n *Livelint) checkIsPullingFromPrivateRegistry(image string) CheckResult {
	yes := n.ui.AskYesNo(fmt.Sprintf("Are you pulling the image %q from a private image registry?", image))

	if !yes {
		return CheckResult{
			Message:      "You are pulling from a public image registry",
			Instructions: "The issue could be with the CRI or Kubelet",
		}
	}

	return CheckResult{
		HasFailed:    true,
		Message:      "You are pulling from a private image registry",
		Details:      []string{fmt.Sprintf("Registry for image %q requires an image pull secret", image)},
		Instructions: "Configure pulling images from a private registry",
	}
}
