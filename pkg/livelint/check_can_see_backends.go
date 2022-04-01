package livelint

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) checkCanSeeBackends(ingressName, namespace string) CheckResult {
	ingress, err := n.k8s.NetworkingV1().Ingresses(namespace).Get(context.Background(), ingressName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("error getting backends for ingress %s in namespace %s: %w", ingressName, namespace, err))
	}

	hasResourceBackends := false

	for _, rules := range ingress.Spec.Rules {
		if rules.HTTP == nil {
			continue
		}
		for _, path := range rules.HTTP.Paths {
			switch {
			case path.Backend.Service != nil:
				service, err := n.k8s.CoreV1().Services(namespace).Get(context.Background(), path.Backend.Service.Name, metav1.GetOptions{})
				if err != nil {
					log.Fatal(fmt.Errorf("error getting backends for ingress %s in namespace %s: %w", ingressName, namespace, err))
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
						Message:   fmt.Sprintf("No backends available for ingress path %s", path.Path),
					}
				}

				endpoint, err := n.k8s.CoreV1().Endpoints(namespace).Get(context.Background(), path.Backend.Service.Name, metav1.GetOptions{})
				if err != nil {
					log.Fatal(fmt.Errorf("error getting endpoints in namespace %s: %w", namespace, err))
				}

				if len(endpoint.Subsets) < 1 {
					return CheckResult{
						HasFailed: true,
						Message:   fmt.Sprintf("No backends available for ingress path %s, because service %s has no endpoints", path.Path, path.Backend.Service.Name),
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
			Message: "There are Backends available for this ingress",
			Details: []string{"Some paths have a resource backend instead of a service. Ensure this backend is setup correctly."},
		}
	}

	return CheckResult{
		Message: "There are Backends available for this ingress",
	}
}
