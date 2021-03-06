package livelint

func (n *Livelint) checkIsImageNameCorrect() CheckResult {
	yes := n.askUserYesOrNo("Is the name of the image correct?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "The name of the image is not correct",
			Instructions: "Fix the image name",
		}
	}

	return CheckResult{
		Message: "The name of the image is correct",
	}
}
