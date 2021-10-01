package livelint

func checkTargetPortMatchesContainerPort() CheckResult {
	yes := askUserYesOrNo("Is the targetPort on the Service matching the containerPort in the Pod?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "The targetPort on the Service doesn't match the containerPort in the Pod",
			Instructions: "Fix the Service targetPort and the containerPort",
		}
	}

	return CheckResult{
		Message:      "The targetPort on the Service matches the containerPort in the Pod",
		Instructions: "The issue could be with the Kube Proxy",
	}
}
