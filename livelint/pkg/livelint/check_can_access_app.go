package livelint

func checkCanAccessApp() CheckResult {
	yes := askUserYesOrNo("Run 'kubectl port-forward <pod-name> 8080:<pod-port>'.\nCan you access the app?")

	if !yes {
		return CheckResult{
			HasFailed: true,
			Message:   "You cannot access the app",
		}
	}

	return CheckResult{
		Message: "You can access the app",
	}
}
