package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) isRestartCycling(pod corev1.Pod) (bool, bool, string) {
	podEventList := n.getPodEvents(pod)
	var lastUnhealthyEvent corev1.Event
	var lastUnhealthyMessage string
	unhealthyEventFound := false
	backoffEventsFound := false
	for _, event := range podEventList.Items {
		if event.Reason == "BackOff" && event.Count > 5 {
			backoffEventsFound = true
		}

		if event.Reason == "Unhealthy" {
			if event.LastTimestamp.After(lastUnhealthyEvent.LastTimestamp.Time) {
				lastUnhealthyEvent = event
				lastUnhealthyMessage = event.Message
				unhealthyEventFound = true
			}
		}
	}

	return backoffEventsFound, unhealthyEventFound, lastUnhealthyMessage
}
