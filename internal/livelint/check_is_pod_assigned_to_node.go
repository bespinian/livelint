package livelint

func (n *Livelint) checkIsPodAssignedToNode() CheckResult {
	yes := n.askUserYesOrNo("Run 'kubectl get pods -o wide'.\nIs the Pod assigned to the Node?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "The Pod is not assigned to the Node",
			Instructions: "There is an issue with the Scheduler",
		}
	}

	return CheckResult{
		Message:      "The Pod is assigned to the Node",
		Instructions: "There is an issue with the Kubelet",
	}
}
