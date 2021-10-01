package livelint

func checkAreResourceQuotasHit() CheckResult {
	yes := askUserYesOrNo("Are you hitting the ResourceQuota limits?")

	if yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "You are hitting the ResourceQuota limits",
			Instructions: "Relax the ResourceQuota limits",
		}
	}

	return CheckResult{
		Message: "You are not hitting the ResourceQuota limits",
	}
}
