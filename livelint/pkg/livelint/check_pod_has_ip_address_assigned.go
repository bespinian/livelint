package livelint

func checkPodHasIPAddressAssigned() CheckResult {
	yes := askUserYesOrNo("Does the Pod have an IP address assigned?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "The Pod has no IP address assigned",
			Instructions: "There is an issue with the Controller manager",
		}
	}

	return CheckResult{
		Message:      "The Pod has an IP address assigned",
		Instructions: "There is an issue with the Kubelet",
	}
}
