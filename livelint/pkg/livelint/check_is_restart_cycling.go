package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) checkIsRestartCycling(pod corev1.Pod) CheckResult {
	podEvents := n.getPodEvents(pod)

	var lastUnhealthyEvent corev1.Event
	var lastUnhealthyMessage string
	unhealthyEventFound := false
	backoffEventsFound := false
	for _, event := range podEvents.Items {
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

	if backoffEventsFound && unhealthyEventFound {
		return CheckResult{
			HasFailed:    true,
			Message:      "The Pod is restarting frequently, cycling between Running and CrashLoopBackOff",
			Details:      []string{lastUnhealthyMessage},
			Instructions: "Fix the liveness probe",
		}
	}

	return CheckResult{
		Message: "The Pod is not restarting frequently, cycling between Running and CrashLoopBackOff",
	}
}
