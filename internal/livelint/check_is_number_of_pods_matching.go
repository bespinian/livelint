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
		runningReplicas := rs.Status.Replicas
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

		if *deploymentPodsReplicas > runningReplicas {
			return CheckResult{
				HasFailed:  true,
				HasWarning: false,
				Message:    "Number of pods is lower then expected",
			}
		}

		if *deploymentPodsReplicas < runningReplicas {
			return CheckResult{
				HasWarning: true,
				HasFailed:  false,
				Message:    "Number of pods is bigger the desired. Further checks will be run to find the issue.",
			}
		}
	}

	return CheckResult{
		Message: "Desired number of pods is running",
	}
}
