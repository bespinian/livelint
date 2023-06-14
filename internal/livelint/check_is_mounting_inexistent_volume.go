package livelint

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) checkIsMountingInexistentVolumeSrc(pod apiv1.Pod, namespace string) CheckResult {
	secrets, err := n.K8s.CoreV1().Secrets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		n.ui.DisplayError(fmt.Errorf("error listing Secrets for Namespace %s: %w", namespace, err))
	}
	configMaps, err := n.K8s.CoreV1().ConfigMaps(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		n.ui.DisplayError(fmt.Errorf("error listing ConfigMaps for Namespace %s: %w", namespace, err))
	}

	for _, vol := range pod.Spec.Volumes {
		if vol.Secret != nil {
			secretName := vol.Secret.SecretName
			found := false
			for _, secret := range secrets.Items {
				if secret.Name == secretName {
					found = true
					break
				}
			}
			if !found {
				return CheckResult{
					HasFailed:    true,
					Message:      fmt.Sprintf("The source Secret of volume %q doesn't exist", vol.Name),
					Details:      []string{fmt.Sprintf("The volume %q uses a source Secret %q that doesn't exist", vol.Name, secretName)},
					Instructions: "Create the missing Secret",
				}
			}
		}
		if vol.ConfigMap != nil {
			configMapName := vol.ConfigMap.Name
			found := false
			for _, configMap := range configMaps.Items {
				if configMap.Name == configMapName {
					found = true
					break
				}
			}
			if !found {
				return CheckResult{
					HasFailed:    true,
					Message:      fmt.Sprintf("The source ConfigMap of volume %q doesn't exist", vol.Name),
					Details:      []string{fmt.Sprintf("The volume %q uses a source ConfigMap %q that doesn't exist", vol.Name, configMapName)},
					Instructions: "Create the missing ConfigMap",
				}
			}
		}
	}

	return CheckResult{
		Message: "All volume sources exist",
	}
}
