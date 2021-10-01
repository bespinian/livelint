package livelint

func checkIsClusterFull() CheckResult {
	yes := askUserYesOrNo("Run 'kubectl describe pod <pod-name>'.\nIs the cluster full?")

	if yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "The cluster is full",
			Instructions: "Provision a bigger cluster",
		}
	}

	return CheckResult{
		Message: "The cluster is not full",
	}
}
