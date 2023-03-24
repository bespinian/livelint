package livelint

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) getPodEvents(pod apiv1.Pod) []apiv1.Event {
	opts := metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.kind=Pod,involvedObject.name=%s,involvedObject.namespace=%s", pod.Name, pod.Namespace)}
	eventList, err := n.K8s.CoreV1().Events(pod.Namespace).List(context.Background(), opts)
	if err != nil {
		n.ui.DisplayError(fmt.Errorf("error querying events for pod %s in namespace %s: %w", pod.Name, pod.Namespace, err))
	}

	return eventList.Items
}
