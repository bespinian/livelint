package livelint

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *livelint) checkCanSeeEndpoints(serviceName string, namespace string) CheckResult {
	endpoint, err := n.k8s.CoreV1().Endpoints(namespace).Get(context.Background(), serviceName, v1.GetOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("error getting endpoint %s in namespace %s: %w", serviceName, namespace, err))
	}

	if len(endpoint.Subsets) < 1 {
		return CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf("No endpoints exists on the service %s", serviceName),
		}
	}

	return CheckResult{
		Message: fmt.Sprintf("Endpoints exists for service %s", serviceName),
	}
}
