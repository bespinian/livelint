package livelint

import (
	"fmt"

	"github.com/fatih/color"
	corev1 "k8s.io/api/core/v1"
)

// Sequentially checks pod conditions (pod scheduled, pod initialized,
// containers ready, pod ready) and breaks and prints the first one that is not ok
func (n *livelint) checkPodConditions(pod corev1.Pod, isVerbose bool) {
	allOK := true

	if isVerbose {
		fmt.Printf("Checking Pod conditions of pod %s\n", pod.Name)
	}

	sequentialConditions := [4]corev1.PodConditionType{corev1.PodScheduled,
		corev1.PodInitialized,
		corev1.ContainersReady,
		corev1.PodReady}

	for _, sequentialCondition := range sequentialConditions {
		for _, podCondition := range pod.Status.Conditions {
			if sequentialCondition == podCondition.Type && podCondition.Status != corev1.ConditionTrue {
				fmt.Printf("Pod %s with condition: %s: %s, Reason: %s, Message: %s (%s)\n",
					pod.Name,
					podCondition.Type,
					podCondition.Status,
					podCondition.Reason,
					podCondition.Message,
					podCondition.LastTransitionTime.Format("2006-01-02T15:04:05Z07:00"),
				)
				allOK = false
			}
		}
		if !allOK {
			break
		}
	}

	if allOK && isVerbose {
		color.Green("Conditions for Pod %s are all OK", pod.Name)
	}
}
