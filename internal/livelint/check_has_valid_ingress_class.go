package livelint

import (
	"fmt"

	netv1 "k8s.io/api/networking/v1"
)

func checkHasValidIngressClass(ingress netv1.Ingress, ingressClasses map[string]netv1.IngressClass) CheckResult {
	result := CheckResult{
		HasFailed: false,
		Message:   fmt.Sprintf("Ingress %s has a valid ingress class", ingress.Name),
	}

	ingressClassName := getIngressClassName(ingress)
	_, defaultIngressClassFound := getDefaultIngressClass(ingressClasses)

	if ingressClassName == "" {
		if !defaultIngressClassFound {
			result = CheckResult{
				HasFailed: true,
				Message:   fmt.Sprintf("Ingress %s does not specify an ingress class and there is no default ingress class in the cluster", ingress.Name),
			}
		}
	} else if _, ok := ingressClasses[ingressClassName]; !ok {
		result = CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf("Ingress %s declares ingress class name %s but there is no such ingress class in the cluster", ingress.Name, ingressClassName),
		}
	}

	return result
}
