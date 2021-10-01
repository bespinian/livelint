package livelint

func checkCanSeeEndpoints() CheckResult {
	yes := askUserYesOrNo("Run 'kubectl describe service <service-name>'.\nCan you see a list of endpoints?")

	if !yes {
		return CheckResult{
			HasFailed: true,
			Message:   "You cannot see a list of endpoints",
		}
	}

	return CheckResult{
		Message: "You can see a list of endpoints",
	}
}
