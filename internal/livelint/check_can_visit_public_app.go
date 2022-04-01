package livelint

func (n *Livelint) checkCanVisitPublicApp() CheckResult {
	yes := n.askUserYesOrNo("The app should be working. Can you visit it from the public internet?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "You cannot visit the app from the public internet",
			Instructions: "The issue is likely to be with the infrastructure and how the cluster is exposed",
		}
	}

	return CheckResult{
		Message: "You can visit the app from the public internet",
	}
}
