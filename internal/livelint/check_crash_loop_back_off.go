package livelint

import (
	apiv1 "k8s.io/api/core/v1"
)

func checkCrashLoopBackOff(pod apiv1.Pod, containerName string) CheckResult {
	for _, cs := range pod.Status.ContainerStatuses {
		if cs.Name != containerName {
			continue
		}

		if cs.State.Waiting != nil &&
			cs.State.Waiting.Reason == "CrashLoopBackOff" {
			return CheckResult{
				HasFailed: true,
				Message:   "The Pod status is CrashLoopBackOff",
				Details:   []string{cs.State.Waiting.Message},
			}
		}
	}

	return CheckResult{
		Message: "The Pod is not in a crash loop",
	}
}
