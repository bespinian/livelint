package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) checkCanVisitServiceApp(service apiv1.Service) CheckResult {
	failureDetails := []string{}

	pods, _ := n.getServicePods(service)
	for _, port := range service.Spec.Ports {
		for _, pod := range pods {
			checkFunc := checkTCPConnection
			if port.Protocol == apiv1.ProtocolUDP {
				checkFunc = checkUDPConnection
			}
			portForwardOk, connectionCheckMsg := n.canPortForward(pod, port.TargetPort.IntVal, checkFunc)
			if !portForwardOk {
				failureDetail := fmt.Sprintf("Pod %s has refused %s connection on port %d, forwarded from port %d: %s", pod.Name, port.Protocol, port.TargetPort.IntVal, port.Port, connectionCheckMsg)
				failureDetails = append(failureDetails, failureDetail)
			}
		}
	}

	if len(failureDetails) > 0 {
		return CheckResult{
			HasFailed: true,
			Message:   "One or more ports were not accessible",
			Details:   failureDetails,
		}
	}

	return CheckResult{
		Message: "You can access the service",
	}
}
