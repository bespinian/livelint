package livelint

import (
	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) CheckForgottenCMDInDockerfile(c apiv1.Container) CheckResult {
	resultSuccess := CheckResult{
		Message:      "Your container has defined a command",
		Instructions: "Unknown state",
	}
	if len(c.Command) > 0 {
		return resultSuccess
	}

	yes := n.ui.AskYesNo("Does your container have the correct CMD or ENTRYPOINT instruction in the Dockerfile?")
	if yes {
		return resultSuccess
	}

	return CheckResult{
		HasFailed:    true,
		Message:      "You forgot the CMD or ENTRYPOINT instruction in the Dockerfile",
		Instructions: "Add a CMD or ENTRYPOINT instruction to your Dockerfile that does not terminate immediately",
	}
}
