package livelint

import (
	"context"
	"fmt"

	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) getIngressClasses() (map[string]netv1.IngressClass, error) {
	ingressClassList, err := n.k8s.NetworkingV1().IngressClasses().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error getting ingress classes of cluster: %w", err)
	}
	result := make(map[string]netv1.IngressClass)
	for _, ingressClass := range ingressClassList.Items {
		result[ingressClass.Name] = ingressClass
	}
	return result, nil
}

func getDefaultIngressClass(ingressClasses map[string]netv1.IngressClass) (netv1.IngressClass, bool) {
	found := false
	var result netv1.IngressClass
	for _, ingressClass := range ingressClasses {
		if len(ingressClasses) == 1 || ingressClass.Annotations["ingressclass.kubernetes.io/is-default-class"] == "true" {
			result = ingressClass
			found = true
			break
		}
	}
	return result, found
}

func getIngressClassName(ingress netv1.Ingress) string {
	name := ""
	if ingress.Spec.IngressClassName != nil {
		name = *ingress.Spec.IngressClassName
	}
	legacyAnnotationValue := ingress.Annotations["kubernetes.io/ingress.class"]
	if legacyAnnotationValue != "" {
		name = legacyAnnotationValue
	}
	return name
}
