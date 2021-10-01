package livelint

func checkCanVisitServiceApp() CheckResult {
	yes := askUserYesOrNo("Run 'kubectl port-forward service/<service-name> 8080:<service-port>'.\nCan you visit the app?")

	if !yes {
		return CheckResult{
			HasFailed: true,
			Message:   "You cannot visit the app from the Service",
		}
	}

	return CheckResult{
		Message: "You can visit the app from the Service",
	}
}
