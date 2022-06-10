package livelint

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func getTargetPort(service apiv1.Service, servicePort int32) (intstr.IntOrString, bool) {
	var targetPort intstr.IntOrString
	found := false
	for _, port := range service.Spec.Ports {
		if port.Port == servicePort {
			targetPort = port.TargetPort
			found = true
			break
		}
	}
	return targetPort, found
}

func getPortNumber(pod apiv1.Pod, port intstr.IntOrString) (int32, bool) {
	var portNumber int32
	found := false
	if port.Type == intstr.Int {
		portNumber = port.IntVal
	} else {
		for _, container := range pod.Spec.Containers {
			for _, containerPort := range container.Ports {
				if containerPort.Name == port.StrVal {
					found = true
					portNumber = containerPort.ContainerPort
					break
				}
			}
			if found {
				break
			}
		}
	}
	return portNumber, found
}

func getTargetPortNumber(service apiv1.Service, pod apiv1.Pod, servicePort int32) (int32, bool) {
	var httpTargetPort int32
	httpTargetPortObject, foundObject := getTargetPort(service, servicePort)
	foundPortNumber := false
	if foundObject {
		httpTargetPort, foundPortNumber = getPortNumber(pod, httpTargetPortObject)
	}
	return httpTargetPort, foundPortNumber
}
