package livelint

import (
	"fmt"
	"strings"
	"time"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) CheckReadinessProbe(pods []apiv1.Pod) CheckResult {
	failingProbesCount := 0
	details := []string{}
	for _, pod := range pods {
		events := n.getPodEvents(pod)

		for _, event := range events {
			if event.Reason == "Unhealthy" &&
				strings.HasPrefix(event.Message, "Readiness probe failed:") &&
				event.Count > 3 {
				if event.LastTimestamp.After(time.Now().Add(time.Minute * -5)) {
					failingProbesCount++
					details = append(details,
						fmt.Sprintf("Pod %s had a failing readiness probe within the last 5 minutes. Verify that the readiness probes are set correctly and that the probe within the container is working.", pod.Name),
						fmt.Sprintf("Event message: %s", event.Message),
						"  → Get more information using the following command:",
						fmt.Sprintf("    kubectl get events --field-selector involvedObject.kind=Pod,involvedObject.name=%s,involvedObject.namespace=%s", pod.Name, pod.Namespace),
						"",
					)
				}
			}
		}
	}

	if failingProbesCount > 0 {
		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("There are %d failing ReadinessProbes", failingProbesCount),
			Details:      details,
			Instructions: "Fix the Readiness probe(s)",
		}
	}

	return CheckResult{
		Message:      "There are no failing Readiness probes",
		Instructions: "Unknown state",
	}
}
