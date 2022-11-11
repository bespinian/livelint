package livelint

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) CheckIsNumberOfPodsMatching(namespace string, deploymentName string) CheckResult {
	deployment, err := n.getDeployment(namespace, deploymentName)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting deployment %s in namespace %s: %w", deploymentName, namespace, err))
	}

	replicaSets, err := n.k8s.AppsV1().ReplicaSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("error listing ReplicaSets for Namespace %s: %w", namespace, err))
	}
	for _, rs := range replicaSets.Items {
		deploymentPodsReplicas := deployment.Spec.Replicas
		definedReplicas := rs.Spec.Replicas
		isPartOfDeployment := false

		for _, ownerRef := range rs.ObjectMeta.OwnerReferences {
			if ownerRef.Kind == "Deployment" && ownerRef.Name == deploymentName {
				isPartOfDeployment = true
				break
			}
		}

		if !isPartOfDeployment {
			continue
		}

		if *deploymentPodsReplicas < *definedReplicas {
			return CheckResult{
				HasFailed: true,
				Message:   "Number of pods is less then desired",
			}
		} else if *deploymentPodsReplicas > *definedReplicas {
			return CheckResult{
				Message: "Number of pods is bigger the desired. Further checks will be run to find the issue.",
			}
		}
	}
	return CheckResult{
		Message: "Desired number of pods is running",
	}
}
