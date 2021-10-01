package livelint

func checkIsSelectorMatchingPodLabel() CheckResult {
	yes := askUserYesOrNo("Is the Selector matching the right Pod label?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "The Selector is not matching the Pod label",
			Instructions: "Fix the Service selector. It has to match the Pod labels",
		}
	}

	return CheckResult{
		Message: "The Selector is matching the Pod label",
	}
}
