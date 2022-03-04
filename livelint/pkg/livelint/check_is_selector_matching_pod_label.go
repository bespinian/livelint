package livelint

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *livelint) checkIsSelectorMatchingPodLabel(allPods []corev1.Pod, serviceName string, namespace string) CheckResult {
	service, err := n.k8s.CoreV1().Services(namespace).Get(context.Background(), serviceName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("error getting service %s in namespace %s: %w", serviceName, namespace, err))
	}

	hasMatchingPod := false
	for _, pod := range allPods {
		labelsMatch := true
		for svcLabelKey, svcLabelValue := range service.Spec.Selector {
			podLabelValue, hasLabel := pod.ObjectMeta.Labels[svcLabelKey]

			if !hasLabel || podLabelValue != svcLabelValue {
				labelsMatch = false
				break
			}
		}
		if labelsMatch {
			hasMatchingPod = true
			break
		}
	}

	if !hasMatchingPod {
		return CheckResult{
			HasFailed:    true,
			Message:      "The service Selector is not matching any Pod labels",
			Instructions: "Fix the Service selector. It has to match the Pod labels",
		}
	}

	return CheckResult{
		Message: "The Selector is matching the Pod label",
	}
}
