package livelint

import "fmt"

func (n *Livelint) checkIsPullingFromPrivateRegistry(image string) CheckResult {
	yes := n.askUserYesOrNo(fmt.Sprintf("Are you pulling the image %q from a private image registry?", image))

	if !yes {
		return CheckResult{
			Message:      "You are pulling from a public image registry",
			Instructions: "The issue could be with the CRI or Kubelet",
		}
	}

	return CheckResult{
		Message:      "You are pulling from a private image registry",
		Instructions: "Configure pulling images from a private registry",
	}
}
