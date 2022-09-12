package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func checkIsPodAssignedToNode(allPods []apiv1.Pod) CheckResult {
	for _, pod := range allPods {
		if len(pod.Spec.NodeName) == 0 {
			return CheckResult{
				HasFailed:    true,
				Message:      fmt.Sprintf("The Pod %s is not assigned to any Node", pod.Name),
				Instructions: "There is an issue with the Scheduler",
			}
		}
	}

	return CheckResult{
		Message:      "The Pod is assigned to a Node",
		Instructions: "There is an issue with the Kubelet",
	}
}
