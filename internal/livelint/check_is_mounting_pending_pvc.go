package livelint

import (
	"context"
	"fmt"
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) checkIsMountingPendingPVC(pods []apiv1.Pod, namespace string) CheckResult {
	pvcs, err := n.k8s.CoreV1().PersistentVolumeClaims(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("error listing PVCs in namespace %s: %w", namespace, err))
	}

	for _, pod := range pods {
		for _, vol := range pod.Spec.Volumes {
			if vol.PersistentVolumeClaim != nil {
				pvcName := vol.PersistentVolumeClaim.ClaimName
				for _, pvc := range pvcs.Items {
					if pvc.Name != pvcName {
						continue
					}
					if pvc.Status.Phase == apiv1.ClaimPending {
						return CheckResult{
							HasFailed:    true,
							Message:      fmt.Sprintf("You are mounting a PENDING PersistentVolumeClaim %s in Pod %s", pvcName, pod.Name),
							Instructions: "Fix the PersistentVolumeClaim",
						}
					}
				}
			}
		}
	}

	return CheckResult{
		Message: "You are not mounting any PENDING PersistentVolumeClaims",
	}
}
