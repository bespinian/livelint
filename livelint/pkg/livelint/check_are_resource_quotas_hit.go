package livelint

import (
	"context"
	"fmt"
	"log"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) checkAreResourceQuotasHit(namespace string, deploymentName string) CheckResult {
	replicaSets, err := n.k8s.AppsV1().ReplicaSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("error when querying replica set for deployment %s in namespace %s: %w", deploymentName, namespace, err))
	}
	for _, rs := range replicaSets.Items {
		isPartOfDeployment := false
		for _, ownerRef := range rs.ObjectMeta.OwnerReferences {
			if ownerRef.Kind == "Deployment" && ownerRef.Name == deploymentName {
				isPartOfDeployment = true
			}
		}
		if !isPartOfDeployment {
			continue
		}

		if rs.Status.AvailableReplicas != *rs.Spec.Replicas {
			for _, condition := range rs.Status.Conditions {
				if condition.Type == appsv1.ReplicaSetReplicaFailure && strings.Contains(condition.Message, "exceeded quota:") {
					return CheckResult{
						HasFailed:    true,
						Message:      "You are hitting the resource quota limits",
						Instructions: "Investigate and relax the resource quota limits of the namespace or lower the resources required by some of your workloads in the namespace.",
						Details:      []string{condition.Message},
					}
				}
			}
		}
	}

	return CheckResult{
		Message: "You are not hitting the ResourceQuota limits",
	}
}
