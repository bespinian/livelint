package livelint

func checkCanSeeBackends() CheckResult {
	yes := askUserYesOrNo("Run 'kubectl describe ingress <ingress-name>'.\nCan you see a list of Backends?")

	if !yes {
		return CheckResult{
			HasFailed: true,
			Message:   "You cannot see a list of Backends",
		}
	}

	return CheckResult{
		Message: "You can see a list of Backends",
	}
}
