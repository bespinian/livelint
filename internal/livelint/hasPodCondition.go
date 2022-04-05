package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

// Sequentially checks the given pod conditions.
// It breaks and returns the first one that is not ok.
func hasPodCondition(pod apiv1.Pod, conditionsToCheck []apiv1.PodConditionType) (bool, string) {
	for _, conditionToCheck := range conditionsToCheck {
		for _, podCondition := range pod.Status.Conditions {
			if conditionToCheck == podCondition.Type && podCondition.Status != apiv1.ConditionTrue {
				message := fmt.Sprintf("Pod %s with condition: %s: %s, Reason: %s, Message: %s (%s)\n",
					pod.Name,
					podCondition.Type,
					podCondition.Status,
					podCondition.Reason,
					podCondition.Message,
					podCondition.LastTransitionTime.Format("2006-01-02T15:04:05Z07:00"),
				)
				return true, message
			}
		}
	}

	return false, ""
}
