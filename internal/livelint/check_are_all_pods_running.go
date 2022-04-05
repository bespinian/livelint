package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

// checkAreAllPodsRunning checks if all Pods are in phase RUNNING and all their containers are in state RUNNING.
func checkAreAllPodsRunning(pods []apiv1.Pod) CheckResult {
	nonRunningPods := []apiv1.Pod{}
	for _, pod := range pods {
		if pod.Status.Phase != apiv1.PodRunning {
			nonRunningPods = append(nonRunningPods, pod)
		}
		for _, cs := range pod.Status.ContainerStatuses {
			if cs.State.Running == nil {
				nonRunningPods = append(nonRunningPods, pod)
			}
		}
	}

	if len(nonRunningPods) > 0 {
		nonRunningPodNames := make([]string, 0, len(nonRunningPods))
		for _, pod := range nonRunningPods {
			nonRunningPodNames = append(nonRunningPodNames, pod.ObjectMeta.Name)
		}

		msgTemplate := "There are %v Pods that are not RUNNING"
		if len(nonRunningPods) == 1 {
			msgTemplate = "There is %v Pod that isn't RUNNING"
		}

		return CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf(msgTemplate, len(nonRunningPods)),
			Details:   nonRunningPodNames,
		}
	}

	return CheckResult{
		Message: "All Pods are RUNNING",
	}
}
