package livelint

import (
	"context"
	"fmt"
	"log"
	"strings"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n *Livelint) checkTargetPortMatchesContainerPort(pods []apiv1.Pod, serviceName string, namespace string) CheckResult {
	service, err := n.K8s.CoreV1().Services(namespace).Get(context.Background(), serviceName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("error getting service %s in namespace %s: %w", serviceName, namespace, err))
	}

	unexposedPorts := []string{}
	firstPod := pods[0]
	for _, servicePort := range service.Spec.Ports {
		portIsExposed := false
		for _, container := range firstPod.Spec.Containers {
			for _, containerPort := range container.Ports {
				if containerPort.ContainerPort == servicePort.TargetPort.IntVal &&
					containerPort.Protocol == servicePort.Protocol {
					portIsExposed = true
					break
				}
			}
			if portIsExposed {
				break
			}
		}
		if !portIsExposed {
			unexposedPorts = append(unexposedPorts, fmt.Sprintf("%s %d", servicePort.Protocol, servicePort.TargetPort.IntVal))
		}
	}

	if len(unexposedPorts) > 0 {
		portList := strings.Join(unexposedPorts, ", ")

		return CheckResult{
			HasFailed:    true,
			Message:      fmt.Sprintf("The targetPorts %s on the Service don't match the containerPort in the Pod", portList),
			Instructions: "Fix the Service targetPort or the containerPort",
		}
	}

	return CheckResult{
		Message:      "The targetPorts on the Service match the containerPorts in the Pod",
		Instructions: "The issue could be with the Kube Proxy",
	}
}
