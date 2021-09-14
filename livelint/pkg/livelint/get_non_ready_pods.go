package livelint

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// Sequentially checks pod conditions (pod scheduled, pod initialized,
// containers ready, pod ready) and breaks and prints the first one that is not ok.
func (n *livelint) getNonReadyPods(allPods []corev1.Pod, isVerbose bool) []corev1.Pod {
	nonReadyPods := []corev1.Pod{}
	for _, pod := range allPods {
		if isVerbose {
			fmt.Printf("Checking Pod conditions of pod %s\n", pod.Name)
		}

		sequentialConditions := [4]corev1.PodConditionType{
			corev1.PodScheduled,
			corev1.PodInitialized,
			corev1.ContainersReady,
			corev1.PodReady,
		}

		hasCondition, message := hasPodCondition(pod, sequentialConditions[:], isVerbose)
		if hasCondition {
			fmt.Println(message)
			nonReadyPods = append(nonReadyPods, pod)
		}
	}
	return nonReadyPods
}
