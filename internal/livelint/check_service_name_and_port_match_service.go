package livelint

func (n *Livelint) checkServiceNameAndPortMatchService() CheckResult {
	yes := n.askUserYesOrNo("Are the serviceName and servicePort matching the Service?")

	if !yes {
		return CheckResult{
			HasFailed:    true,
			Message:      "The serviceName and servicePort are not matching the Service",
			Instructions: "Fix the ingress serviceName and servicePort",
		}
	}

	return CheckResult{
		Message:      "The serviceName and servicePort are matching the Service",
		Instructions: "The issue is specific to the Ingress controller. Consult the docs for your Ingress.",
	}
}
