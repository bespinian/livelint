package livelint

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) CheckIsNumberOfPodsMatching(namespace string, deploymentName string) CheckResult {
	deployment, err := n.getDeployment(namespace, deploymentName)
	if err != nil {
		n.ui.DisplayError(fmt.Errorf("error getting deployment %s in namespace %s: %w", deploymentName, namespace, err))
	}

	replicaSets, err := n.K8s.AppsV1().ReplicaSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		n.ui.DisplayError(fmt.Errorf("error listing ReplicaSets for Namespace %s: %w", namespace, err))
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
				HasFailed: true,
				Message:   "Number of pods is lower then expected",
			}
		}

		if *deploymentPodsReplicas < runningReplicas {
			return CheckResult{
				HasFailed: true,
				Message:   "Your cluster is in intermediary state. Number of pods is bigger then expected. Rerun livelint once cluster is in stable state.",
			}
		}
	}

	return CheckResult{
		Message: "Desired number of pods is running",
	}
}
