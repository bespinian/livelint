package livelint

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) checkServiceNameAndPortAreMatching(ingress netv1.Ingress, services []v1.Service, namespace string) CheckResult {

	for _, rule := range ingress.Spec.Rules {
		if rule.HTTP == nil {
			continue
		}

		for _, path := range rule.HTTP.Paths {
			switch {
			case path.Backend.Service != nil:
				//TODO compare here if service name matches any serviceName forwarded to the function as list parameter
				service, err := n.k8s.CoreV1().Services(namespace).Get(context.Background(), path.Backend.Service.Name, metav1.GetOptions{})
				if err != nil {
					log.Fatal(fmt.Errorf("error getting backends for Ingress %s in Namespace %s: %w", ingress.Name, namespace, err))
				}

				//TODO compare port from service forwarded to function and the one from ingress spec
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
						Message:   fmt.Sprintf("Service name and port mismatch for %s", ingress.Name),
					}
				}
			}
		}
	}

	return CheckResult{
		Message: fmt.Sprintf("Service name and port mismatch for %s", ingress.Name),
	}
}
