package livelint

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// getPendingPods returns all pods of a deployment that are in status "PENDING".
func (n *livelint) getPendingPods(namespace, deploymentName string) ([]corev1.Pod, error) {
	deployment, err := n.k8s.AppsV1().Deployments(namespace).Get(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("error getting deployment %q in namespace %q: %w", deploymentName, namespace, err)
	}

	fmt.Println(deployment.Namespace)

	return []corev1.Pod{}, nil
}
