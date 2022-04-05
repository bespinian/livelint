package livelint

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// getPodsForDeployment returns all pods of a deployment.
func (n *Livelint) getPodsForDeployment(namespace, deploymentName string) ([]apiv1.Pod, error) {
	deployment, err := n.getDeployment(namespace, deploymentName)
	if err != nil {
		return nil, err
	}
	matchLabels := deployment.Spec.Selector.MatchLabels
	return n.getPods(namespace, matchLabels)
}

// getPodsForService returns all pods which are selected by a service.
func (n *Livelint) getPodsForService(service apiv1.Service) ([]apiv1.Pod, error) {
	matchLabels := service.Spec.Selector
	return n.getPods(service.Namespace, matchLabels)
}

// getPods returns all pods which are selected by a map of labels.
func (n *Livelint) getPods(namespace string, matchLabels map[string]string) ([]apiv1.Pod, error) {
	options := metav1.ListOptions{
		LabelSelector: labels.Set(matchLabels).String(),
	}
	pods, err := n.k8s.CoreV1().Pods(namespace).List(context.Background(), options)
	if err != nil {
		return []apiv1.Pod{}, fmt.Errorf("error listing pods: %w", err)
	}

	return pods.Items, nil
}
