package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) CheckAreThereRestartCyclingPods(pods []apiv1.Pod) CheckResult {
	reasonsToFail := []string{}
	for _, pod := range pods {
		isCycling, reason := n.isRestartCycling(pod)
		if isCycling {
			reasonsToFail = append(reasonsToFail, reason)
		}
	}

	if len(reasonsToFail) > 0 {
		msgTemplate := "There are %v Pods cycling between RUNNING and CRASHING"
		if len(reasonsToFail) == 1 {
			msgTemplate = "There is %v Pod cycling between RUNNING and CRASHING"
		}

		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf(msgTemplate, len(reasonsToFail)),
			Details:      reasonsToFail,
			Instructions: "Fix the respective liveness probe(s)",
		}

	}

	return CheckResult{
		Message: "No Pods are restart cycling",
	}
}

func (n *Livelint) isRestartCycling(pod apiv1.Pod) (bool, string) {
	events := n.getPodEvents(pod)

	var lastUnhealthyEvent apiv1.Event
	var lastUnhealthyMessage string
	unhealthyEventFound := false
	backoffEventsFound := false
	for _, event := range events {
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
		return true, lastUnhealthyMessage
	}

	return false, ""
}
