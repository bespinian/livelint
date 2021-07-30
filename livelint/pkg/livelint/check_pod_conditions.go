package livelint

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) checkPodConditions(pod corev1.Pod) {
	allOk := true
	fmt.Printf("Checking Pod conditions of pod %s\n", pod.Name)
	for i := 0; i < len(pod.Status.Conditions); i++ {
		condition := pod.Status.Conditions[i]
		if condition.Status != corev1.ConditionTrue {
			fmt.Printf("    %s: %s, Reason: %s, Message: %s (%s)\n",
				condition.Type,
				condition.Status,
				condition.Reason,
				condition.Message,
				condition.LastTransitionTime.Format("2006-01-02T15:04:05Z07:00"),
			)
			allOk = false
		}
	}

	if allOk {
		fmt.Println("    Pod conditions are all ok")
	}
}
