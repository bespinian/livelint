package livelint

import (
	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) CheckFailedMount(pod apiv1.Pod) CheckResult {
	events := n.getPodEvents(pod)

	for _, event := range events {
		if event.Reason == "FailedMount" {
			return CheckResult{
				HasFailed:    true,
				Message:      "The Pod is unable to mount a volume",
				Details:      []string{event.Message},
				Instructions: "Check if the volume exists, the referenced name is correct, and the pod has access to it",
			}
		}
	}

	return CheckResult{
		Message: "There are no issues mounting volumes",
	}
}
