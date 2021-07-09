package livelint

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// getNonRunningPods returns all pods of a deployment that are not in status "RUNNING".
func (n *livelint) getNonRunningPods(namespace, deploymentName string) ([]corev1.Pod, error) {
	deployment, err := n.k8s.AppsV1().Deployments(namespace).Get(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("error getting deployment %q in namespace %q: %w", deploymentName, namespace, err)
	}

	matchLabels := deployment.Spec.Selector.MatchLabels
	options := metav1.ListOptions{
		LabelSelector: labels.Set(matchLabels).String(),
	}
	pods, err := n.k8s.CoreV1().Pods(namespace).List(context.Background(), options)

	nonRunningPods := []corev1.Pod{}
	for i := 0; i < len(pods.Items); i++ {
		if pods.Items[i].Status.Phase != corev1.PodRunning {
			nonRunningPods = append(nonRunningPods, pods.Items[i])
		}
	}

	return nonRunningPods, nil
}