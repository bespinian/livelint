package livelint

func checkAreThereRunningContainers() CheckResult {
	yes := askUserYesOrNo("Is there any container RUNNING?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "There are no containers RUNNING",
			Instructions: "Consult StackOverflow",
		}
	}

	return CheckResult{
		Message:      "There are containers RUNNING",
		Instructions: "The issue is with the node-lifecycle controller",
	}
}
