package livelint

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *livelint) getPodEvents(pod corev1.Pod) *corev1.EventList {
	podEventListOptions := metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.kind=Pod,involvedObject.name=%s,involvedObject.namespace=%s", pod.Name, pod.Namespace)}
	podEventList, err := n.k8s.CoreV1().Events(pod.Namespace).List(context.Background(), podEventListOptions)
	if err != nil {
		log.Fatal(fmt.Errorf("error when querying events for pod %s in namespace %s: %w", pod.Name, pod.Namespace, err))
	}

	return podEventList
}
