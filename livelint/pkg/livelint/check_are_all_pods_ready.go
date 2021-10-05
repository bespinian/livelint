package livelint

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// checkAreAllPodsReady sequentially checks pod conditions (pod scheduled, pod initialized,
// containers ready, pod ready) and breaks and prints the first one that is not ok.
func checkAreAllPodsReady(pods []corev1.Pod) CheckResult {
	nonReadyPods := []corev1.Pod{}
	for _, pod := range pods {
		sequentialConditions := [4]corev1.PodConditionType{
			corev1.PodScheduled,
			corev1.PodInitialized,
			corev1.ContainersReady,
			corev1.PodReady,
		}

		hasCondition, _ := hasPodCondition(pod, sequentialConditions[:])
		if hasCondition {
			nonReadyPods = append(nonReadyPods, pod)
		}
	}

	if len(nonReadyPods) > 0 {
		nonReadyPodNames := make([]string, 0, len(nonReadyPods))
		for _, pod := range nonReadyPods {
			nonReadyPodNames = append(nonReadyPodNames, pod.ObjectMeta.Name)
		}

		return CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf("There are %d Pods that aren't READY", len(nonReadyPods)),
			Details:   nonReadyPodNames,
		}
	}

	return CheckResult{
		Message: "All Pods are READY",
	}
}
