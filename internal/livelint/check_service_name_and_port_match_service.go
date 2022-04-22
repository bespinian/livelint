package livelint

import (
	"context"
	"fmt"
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) checkServiceNameAndPortMatchService(ingressName, namespace string, service apiv1.Service) CheckResult {
	ingress, err := n.k8s.NetworkingV1().Ingresses(namespace).Get(context.Background(), ingressName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("error getting ingress %s in namespace %s: %w", ingressName, namespace, err))
	}

	for _, rule := range ingress.Spec.Rules {
		for _, path := range rule.HTTP.Paths {
			if path.Backend.Service.Name != service.Name {
				continue
			}
			for _, port := range service.Spec.Ports {
				if port.Port == path.Backend.Service.Port.Number {
					return CheckResult{
						Message:      "The serviceName and servicePort are matching the Service",
						Instructions: "The issue is specific to the Ingress controller. Consult the docs for your Ingress.",
					}
				}
			}
		}
	}

	return CheckResult{
		HasFailed:    true,
		Message:      "The serviceName and servicePort are not matching the Service",
		Instructions: "Fix the ingress serviceName and servicePort",
	}
}
