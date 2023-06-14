package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

// checkAreAllPodsReady sequentially checks pod conditions (pod scheduled, pod initialized,
// containers ready, pod ready) and breaks and prints the first one that is not ok.
func checkAreAllPodsReady(pods []apiv1.Pod) CheckResult {
	nonReadyPods := []apiv1.Pod{}
	for _, pod := range pods {
		sequentialConditions := [4]apiv1.PodConditionType{
			apiv1.PodScheduled,
			apiv1.PodInitialized,
			apiv1.ContainersReady,
			apiv1.PodReady,
		}

		hasCondition, _ := hasPodCondition(pod, sequentialConditions[:])
		if hasCondition {
			nonReadyPods = append(nonReadyPods, pod)
		}
	}

	if len(nonReadyPods) > 0 {
		nonReadyPodNames := make([]string, 0, len(nonReadyPods))
		for _, pod := range nonReadyPods {
			nonReadyPodNames = append(nonReadyPodNames, pod.Name)
		}

		msgTemplate := "There are %v Pods that are not READY"
		if len(nonReadyPods) == 1 {
			msgTemplate = "There is %v Pod that is not READY"
		}

		return CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf(msgTemplate, len(nonReadyPods)),
			Details:   nonReadyPodNames,
		}
	}

	return CheckResult{
		Message: "All Pods are READY",
	}
}
