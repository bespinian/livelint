package livelint

func (n *Livelint) checkIsPullingFromPrivateRegistry() CheckResult {
	yes := n.askUserYesOrNo("Are you pulling from a private image registry?")

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
