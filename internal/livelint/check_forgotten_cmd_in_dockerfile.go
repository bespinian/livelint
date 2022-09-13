package livelint

func (n *Livelint) checkForgottenCMDInDockerfile() CheckResult {
	yes := n.askUserYesOrNo("Did you forget the CMD instruction in the Dockerfile?")

	if yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "You forgot the CMD instruction in the Dockerfile",
			Instructions: "Add a CMD instruction to your Dockerfile",
		}
	}

	return CheckResult{
		Message:      "Your Dockerfile has a CMD instruction",
		Instructions: "Unknown state",
	}
}
