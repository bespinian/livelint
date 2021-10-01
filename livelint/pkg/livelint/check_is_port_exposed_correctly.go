package livelint

func checkIsPortExposedCorrectly() CheckResult {
	yes := askUserYesOrNo("Is the port exposed by container correct and listing on 0.0.0.0?")

	if yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "The port exposed by the container isn't correct or isn't listening on 0.0.0.0",
			Instructions: "Fix the app. It should listen on 0.0.0.0. Update the container port.",
		}
	}

	return CheckResult{
		Message:      "The port exposed by the container is correct and listening on 0.0.0.0",
		Instructions: "Unknown state",
	}
}
