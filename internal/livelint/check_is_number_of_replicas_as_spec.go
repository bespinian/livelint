package livelint

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) CheckIsNumberOfReplicasCorrect(namespace string, deploymentName string) CheckResult {
	deployment, err := n.getDeployment(namespace, deploymentName)
	if err != nil {
		n.ui.DisplayError(fmt.Errorf("error getting Deployment %s in Namespace %s: %w", deploymentName, namespace, err))
	}

	replicaSets, err := n.K8s.AppsV1().ReplicaSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		n.ui.DisplayError(fmt.Errorf("error listing ReplicaSets for Namespace %s: %w", namespace, err))
	}

	for _, rs := range replicaSets.Items {
		if *rs.Spec.Replicas == 0 {
			continue
		}

		isOwnedByDeployment := false
		for _, ownerRef := range rs.ObjectMeta.OwnerReferences {
			if ownerRef.Kind == "Deployment" && ownerRef.Name == deploymentName {
				isOwnedByDeployment = true
				break
			}
		}
		if !isOwnedByDeployment {
			continue
		}

		desiredReplicas := deployment.Spec.Replicas
		actualReplicas := rs.Status.Replicas

		details := []string{
			fmt.Sprintf("ReplicaSet %q has %d replica(s).", rs.Name, actualReplicas),
			fmt.Sprintf("According to the spec of Deployment %q, it should have %v.", deployment.Name, *desiredReplicas),
		}

		if actualReplicas < *desiredReplicas {
			return CheckResult{
				HasFailed: true,
				Message:   "Number of replicas is lower than desired",
				Details:   details,
			}
		}

		if actualReplicas > *desiredReplicas {
			return CheckResult{
				HasFailed: true,
				Message:   "Cluster in intermediary state. Number of replicas is larger than desired. Re-run livelint once cluster is in stable state.",
				Details:   details,
			}
		}
	}

	return CheckResult{
		Message: "Number of replicas is as desired",
	}
}
