package livelint

func checkCanVisitIngressApp() CheckResult {
	yes := askUserYesOrNo("Run 'kubectl port-forward <ingress-pod-name> 8080:<ingress-port>'.\nCan you visit the app?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "You cannot visit the app from the Ingress",
			Instructions: "The issue is specific to the Ingress controller. Consult the docs for your Ingress.",
		}
	}

	return CheckResult{
		Message: "You can visit the app from the Ingress",
	}
}
