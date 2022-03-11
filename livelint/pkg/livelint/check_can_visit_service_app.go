package livelint

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
)

func (n *Livelint) checkCanVisitServiceApp(service v1.Service) CheckResult {
	failureDetails := []string{}
	pods, _ := n.getPodsForService(service)
	for _, port := range service.Spec.Ports {
		for _, pod := range pods {
			if !n.portForwardAndCheck(pod, port.TargetPort.IntVal) {
				failureDetail := fmt.Sprintf("Pod %s has refused connection on port %d, forwarded from port %d", pod.Name, port.TargetPort.IntVal, port.Port)
				failureDetails = append(failureDetails, failureDetail)
			}
		}
	}

	checkResult := CheckResult{
		Message: "You can access the service",
	}
	if len(failureDetails) > 0 {
		checkResult = CheckResult{
			Message:   "One or more ports were not acessible",
			HasFailed: true,
			Details:   failureDetails,
		}
	}

	return checkResult
}
