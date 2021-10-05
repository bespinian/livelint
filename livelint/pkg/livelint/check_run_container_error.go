package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

func checkRunContainerError(pod corev1.Pod) CheckResult {
	if pod.Status.Phase == "RunContainerError" {
		return CheckResult{
			HasFailed:    true,
			Message:      "The Pod status is RunContainerError",
			Instructions: "The issue is likely to be with mounting volumes",
		}
	}

	return CheckResult{
		Message: "The Pod status is not RunContainerError",
	}
}
