package livelint

import (
	"fmt"

	"github.com/fatih/color"
	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) checkPodConditions(pod corev1.Pod, isVerbose bool) {
	allOk := true
	if (isVerbose) {
		fmt.Printf("Checking Pod conditions of pod %s\n", pod.Name)
	}
	for i := 0; i < len(pod.Status.Conditions); i++ {
		condition := pod.Status.Conditions[i]
		if condition.Status != corev1.ConditionTrue {
			fmt.Printf("Pod %s with condition: %s: %s, Reason: %s, Message: %s (%s)\n",
				pod.Name,
				condition.Type,
				condition.Status,
				condition.Reason,
				condition.Message,
				condition.LastTransitionTime.Format("2006-01-02T15:04:05Z07:00"),
			)
			allOk = false
		}
	}

	if allOk && isVerbose {
		color.Green("Conditions for pod %s are all ok", pod.Name)
	}
}
