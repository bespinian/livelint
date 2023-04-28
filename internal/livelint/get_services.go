package livelint

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// isSubmap checks whether all elements of map sourceMatchLabels are also elements of the map targetMatchLabels (i.e. the source labels select the same or more pods).
func isSubmap(sourceMatchLabels, targetMatchLabels map[string]string) bool {
	result := true
	for key, value := range sourceMatchLabels {
		if targetMatchLabels[key] != value {
			result = false
			break
		}
	}
	return result
}

func (n *Livelint) getServices(namespace, deploymentName string) ([]apiv1.Service, []apiv1.Service, error) {
	deployment, err := n.getDeployment(namespace, deploymentName)
	if err != nil {
		return nil, nil, err
	}

	deploymentMatchLabels := deployment.Spec.Selector.MatchLabels
	return n.getServicesForSelector(namespace, deploymentMatchLabels)
}

func (n *Livelint) getServicesForPod(namespace string, pod apiv1.Pod) ([]apiv1.Service, []apiv1.Service, error) {
	podMatchLabels := pod.Labels
	return n.getServicesForSelector(namespace, podMatchLabels)
}

// getServicesForSelector gets a list of all services which select pods with a given set of labels. It also returns
// a list of all services which potentially match a superset of pods. This is useful for warning the user.
func (n *Livelint) getServicesForSelector(namespace string, matchLabels map[string]string) ([]apiv1.Service, []apiv1.Service, error) {
	allServices, err := n.K8s.CoreV1().Services(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("error listing services in namespace %q: %w", namespace, err)
	}
	exactlyMatchingServices := []apiv1.Service{}
	supersetMatchingServices := []apiv1.Service{}
	var serviceLabelsSubsetDeploymentLabels bool
	var deploymentLabelsSubsetServiceLabels bool
	for _, service := range allServices.Items {
		serviceMatchLabels := service.Spec.Selector
		deploymentLabelsSubsetServiceLabels = isSubmap(matchLabels, serviceMatchLabels)
		serviceLabelsSubsetDeploymentLabels = isSubmap(serviceMatchLabels, matchLabels)
		if serviceLabelsSubsetDeploymentLabels && !deploymentLabelsSubsetServiceLabels {
			supersetMatchingServices = append(supersetMatchingServices, service)
		} else if serviceLabelsSubsetDeploymentLabels && deploymentLabelsSubsetServiceLabels {
			exactlyMatchingServices = append(exactlyMatchingServices, service)
		}
	}
	return exactlyMatchingServices, supersetMatchingServices, nil
}
