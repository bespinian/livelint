package livelint

func (n *Livelint) checkIsPortExposedCorrectly() CheckResult {
	yes := n.askUserYesOrNo("Is the port exposed by container correct and listening on 0.0.0.0?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "The port exposed by the container is not correct or is not listening on 0.0.0.0",
			Instructions: "Fix the app. It should listen on 0.0.0.0. Update the container port.",
		}
	}

	return CheckResult{
		Message:      "The port exposed by the container is correct and listening on 0.0.0.0",
		Instructions: "Unknown state",
	}
}
