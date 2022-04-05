package livelint

func checkDidInspectLogsAndFix() CheckResult {
	yes := askUserYesOrNo("Did you inspect the logs fix the crashing app?")

	if yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "You didn't inspect the logs and fix the crashing app",
			Instructions: "Fix crashing app",
		}
	}

	return CheckResult{
		Message: "You inspected the logs and fixed the crashing app",
	}
}
