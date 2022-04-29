package livelint

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// getIngressesFromServices gets a list of all ingresses which have a backend with a servicename and
// port matching one of the given services.
func (n *Livelint) getIngressesFromServices(namespace string, services []apiv1.Service) ([]netv1.Ingress, error) {
	ingresses, err := n.k8s.NetworkingV1().Ingresses(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error getting ingresses in namespace %q: %w", namespace, err)
	}

	matchingIngresses := []netv1.Ingress{}

	for _, ingress := range ingresses.Items {
		for _, rule := range ingress.Spec.Rules {
			for _, path := range rule.HTTP.Paths {
				for _, service := range services {
					if path.Backend.Service.Name == service.Name {
						for _, port := range service.Spec.Ports {
							if port.Port == path.Backend.Service.Port.Number {
								matchingIngresses = append(matchingIngresses, ingress)
							}
						}
					}
				}
			}
		}
	}

	return matchingIngresses, nil
}
