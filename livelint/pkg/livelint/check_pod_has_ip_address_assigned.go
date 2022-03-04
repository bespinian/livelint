package livelint

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func checkPodHasIPAddressAssigned(allPods []corev1.Pod) CheckResult {
	for _, pod := range allPods {
		if len(pod.Status.PodIP) <= 0 {
			return CheckResult{
				HasFailed:    true,
				Message:      fmt.Sprintf("The Pod %s has no IP address assigned", pod.Name),
				Instructions: "There is an issue with the Controller manager",
			}
		}
	}

	return CheckResult{
		Message:      "The Pods have an IP address assigned",
		Instructions: "There is an issue with the Kubelet",
	}
}
