package livelint

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) checkReadinessProbe(nonReadyPods []corev1.Pod) bool {
	red := color.New(color.FgRed)

	for _, pod := range nonReadyPods {
		podEventList := n.getPodEvents(pod)

		for _, event := range podEventList.Items {
			if event.Reason == "Unhealthy" &&
				strings.HasPrefix(event.Message, "Readiness probe failed:") &&
				event.Count > 3 {
				if event.LastTimestamp.After(time.Now().Add(time.Minute * -5)) {
					red.Printf("Pod %s had a failing readiness probe within the last 5 minutes. Verify that the readiness probes are set correctly and that the probe within the container is working.\n", pod.Name)
					fmt.Printf("Event message: %s\n", event.Message)
					fmt.Println("Get more information using the following command")
					fmt.Printf("    kubectl get events --field-selector involvedObject.kind=Pod,involvedObject.name=%s,involvedObject.namespace=%s\n", pod.Name, pod.Namespace)
					return true
				}
			}
		}
	}
	return false
}
