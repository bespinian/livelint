package livelint

import (
	"context"
	"fmt"
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Check if service's selector is matching at least one pod's label.
func (n *Livelint) checkIsSelectorMatchPodLabel(namespace, serviceName string, pods []apiv1.Pod) CheckResult {
	service, err := n.K8s.CoreV1().Services(namespace).Get(context.Background(), serviceName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("error getting service %s in namespace %s: %w", serviceName, namespace, err))
	}
	serviceSelector := service.Spec.Selector
	if len(serviceSelector) == 0 {
		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("No selector is set for service %s", service),
			Instructions: "Fix the Service selector",
		}
	}
	var matchingPods []bool
	for _, pod := range pods {
		if isSubmap(serviceSelector, pod.Labels) {
			matchingPods = append(matchingPods, true)
		}
	}
	if len(matchingPods) != len(pods) {
		return CheckResult{
			HasFailed:    true,
			Message:      "Not all pods in the deployment has label that matches service's selector.",
			Instructions: "Fix pods labels",
		}
	}

	return CheckResult{
		Message:      "The Service's selector match Pods labels",
		Instructions: "The issue could be that Pod's missing assigned IP address",
	}
}
