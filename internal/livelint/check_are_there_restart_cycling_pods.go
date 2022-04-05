package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) checkAreThereRestartCyclingPods(allPods []apiv1.Pod) CheckResult {
	failedChecks := []CheckResult{}
	for _, pod := range allPods {
		result := n.checkIsRestartCycling(pod)
		if result.HasFailed {
			failedChecks = append(failedChecks, result)
		}
	}
	if len(failedChecks) > 0 {
		cyclingPodReasons := make([]string, 0, len(failedChecks))
		for _, result := range failedChecks {
			for _, reason := range result.Details {
				cyclingPodReasons = append(cyclingPodReasons, reason)
			}
		}
		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("There are %d Pods which are cycling between running an crashing", len(failedChecks)),
			Details:      cyclingPodReasons,
			Instructions: "Fix the liveness probes",
		}

	}
	return CheckResult{
		Message: "No Pods are restart cycling",
	}
}
