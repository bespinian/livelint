package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) checkFailedMount(pod corev1.Pod) CheckResult {
	events := n.getPodEvents(pod)
	for _, event := range events.Items {
		if event.Reason == "FailedMount" {
			return CheckResult{
				HasFailed:    true,
				Message:      "The Pod is unable to mount a volume",
				Details:      []string{event.Message},
				Instructions: "Check whether the volume exists, the referenced name is correct, and that the pod has access to it.",
			}
		}
	}

	return CheckResult{
		Message: "There appear to be no issues mounting volumes.",
	}
}