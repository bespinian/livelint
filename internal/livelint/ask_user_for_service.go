package livelint

import apiv1 "k8s.io/api/core/v1"

func (n *Livelint) askUserForService(namespace, deploymentName string) apiv1.Service {
	exactlyMatchingServices, partlyMatchingServices, _ := n.getServices(namespace, deploymentName)
	services := []apiv1.Service{}
	services = append(services, exactlyMatchingServices...)
	services = append(services, partlyMatchingServices...)
	serviceNames := getServiceNames(services)
	chosenIndex := make(chan int)
	n.tea.Send(listChoiceMsg{title: "Please choose the service to check", items: serviceNames, choice: chosenIndex})
	i := <-chosenIndex
	return services[i]
}
