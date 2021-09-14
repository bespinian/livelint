package livelint

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *livelint) getPodEvents(namespace string, pod corev1.Pod) *corev1.EventList {
	podEventListOptions := metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.kind=Pod,involvedObject.name=%s,involvedObject.namespace=%s", pod.Name, namespace)}
	podEventList, err := n.k8s.CoreV1().Events(namespace).List(context.Background(), podEventListOptions)
	if err != nil {
		log.Fatal(fmt.Errorf("error when querying events for pod %s in namespace %s: %w", pod.Name, namespace, err))
	}
	return podEventList
}

func (n *livelint) isRestartCycling(namespace string, pod corev1.Pod) (bool, string) {
	podEventList := n.getPodEvents(namespace, pod)
	var lastUnhealthyEvent corev1.Event
	unhealthyEventFound := false
	for _, event := range podEventList.Items {
		if event.Reason == "Unhealthy" {
			if event.LastTimestamp.After(lastUnhealthyEvent.LastTimestamp.Time) {
				lastUnhealthyEvent = event
				unhealthyEventFound = true
			}
		}
	}
	if unhealthyEventFound {
		return true, lastUnhealthyEvent.Message
	}
	return false, ""
}
