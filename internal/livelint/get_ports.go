package livelint

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func getTargetPort(service apiv1.Service, servicePort int32) (intstr.IntOrString, bool) {
	var targetPort intstr.IntOrString
	for _, port := range service.Spec.Ports {
		if port.Port == servicePort {
			targetPort = port.TargetPort
			return targetPort, true
		}
	}
	return targetPort, false
}

func getPortNumber(pod apiv1.Pod, port intstr.IntOrString) (int32, bool) {
	var portNumber int32
	if port.Type == intstr.Int {
		portNumber = port.IntVal
	} else {
		for _, container := range pod.Spec.Containers {
			for _, containerPort := range container.Ports {
				if containerPort.Name == port.StrVal {
					portNumber = containerPort.ContainerPort
					return portNumber, true
				}
			}
		}
	}
	return portNumber, false
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
