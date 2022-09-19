package livelint

import (
	"context"
	"fmt"
	"log"

	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) checkCanSeeBackends(ingress netv1.Ingress, namespace string) CheckResult {
	hasResourceBackends := false

	for _, rule := range ingress.Spec.Rules {
		if rule.HTTP == nil {
			continue
		}

		for _, path := range rule.HTTP.Paths {
			switch {
			case path.Backend.Service != nil:
				service, err := n.k8s.CoreV1().Services(namespace).Get(context.Background(), path.Backend.Service.Name, metav1.GetOptions{})
				if err != nil {
					log.Fatal(fmt.Errorf("error getting backends for Ingress %s in Namespace %s: %w", ingress.Name, namespace, err))
				}

				hasPortMatch := false
				for _, servicePort := range service.Spec.Ports {
					if servicePort.Port == path.Backend.Service.Port.Number {
						hasPortMatch = true
						break
					}
				}

				if !hasPortMatch {
					return CheckResult{
						HasFailed: true,
						Message:   fmt.Sprintf("No backends available for Ingress path %s", path.Path),
					}
				}

				endpoint, err := n.k8s.CoreV1().Endpoints(namespace).Get(context.Background(), path.Backend.Service.Name, metav1.GetOptions{})
				if err != nil {
					log.Fatal(fmt.Errorf("error getting endpoints in namespace %s: %w", namespace, err))
				}

				if len(endpoint.Subsets) < 1 {
					return CheckResult{
						HasFailed: true,
						Message:   fmt.Sprintf("No backends available for Ingress path %s, because service %s has no endpoints", path.Path, path.Backend.Service.Name),
					}
				}
			case path.Backend.Resource != nil:
				hasResourceBackends = true

			default:
				return CheckResult{
					HasFailed: true,
					Message:   fmt.Sprintf("Did not find backend for path %s", path.Path),
				}
			}
		}
	}

	if hasResourceBackends {
		return CheckResult{
			Message: fmt.Sprintf("There are Backends available for Ingress %s", ingress.Name),
			Details: []string{"Some paths have a resource backend instead of a service. Ensure this backend is set up correctly."},
		}
	}

	return CheckResult{
		Message: fmt.Sprintf("There are Backends available for Ingress %s", ingress.Name),
	}
}
