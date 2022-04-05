package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

// checkAreAllPodsRunning checks if all Pods are in phase RUNNING and all their containers are in state RUNNING.
func checkAreAllPodsRunning(allPods []apiv1.Pod) CheckResult {
	nonRunningPods := []apiv1.Pod{}
	for _, pod := range allPods {
		if pod.Status.Phase != apiv1.PodRunning {
			nonRunningPods = append(nonRunningPods, pod)
		}
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.State.Running == nil {
				nonRunningPods = append(nonRunningPods, pod)
			}
		}
	}

	if len(nonRunningPods) > 0 {
		nonRunningPodNames := make([]string, 0, len(nonRunningPods))
		for _, pod := range nonRunningPods {
			nonRunningPodNames = append(nonRunningPodNames, pod.ObjectMeta.Name)
		}

		return CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf("There are %d Pods that are not RUNNING", len(nonRunningPods)),
			Details:   nonRunningPodNames,
		}
	}

	return CheckResult{
		Message: "All Pods are RUNNING",
	}
}
