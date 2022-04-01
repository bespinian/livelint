package livelint

func (n *Livelint) checkIsImageTagValid() CheckResult {
	yes := n.askUserYesOrNo("Is the image tag valid? Does it exist?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "The image tag is not valid or doesn't exist",
			Instructions: "Fix the tag",
		}
	}

	return CheckResult{
		Message: "The image tag is valid and exists",
	}
}
