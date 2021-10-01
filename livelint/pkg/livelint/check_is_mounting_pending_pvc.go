package livelint

func checkIsMountingPendingPVC() CheckResult {
	yes := askUserYesOrNo("Are you mounting a PENDING PersistentVolumeClaim?")

	if yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "You are mounting a PENDING PersistentVolumeClaim",
			Instructions: "Fix the PersistentVolumeClaim",
		}
	}

	return CheckResult{
		Message: "You are not mounting any PENDING PersistentVolumeClaims",
	}
}
