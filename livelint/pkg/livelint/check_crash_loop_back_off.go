package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

func checkCrashLoopBackOff(pod corev1.Pod, containerName string) CheckResult {
	for _, cs := range pod.Status.ContainerStatuses {
		if cs.Name != containerName {
			continue
		}

		if cs.State.Waiting != nil &&
			cs.State.Waiting.Reason == "CrashLoopBackOff" {
			return CheckResult{
				HasFailed: true,
				Message:   "The Pod status is CrashLoopBackOff",
			}
		}
	}

	return CheckResult{
		Message: "The Pod status is not CrashLoopBackOff",
	}
}
