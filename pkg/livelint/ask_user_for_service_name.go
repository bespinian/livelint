package livelint

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func askUserForServiceName(exactlyMatching, partlyMatching []apiv1.Service) apiv1.Service {
	fmt.Println("")
	fmt.Println("Matching services:")
	for i, service := range exactlyMatching {
		fmt.Printf("%v. %v \n", i+1, service.Name)
	}
	fmt.Println("")
	fmt.Println("Partly matching services:")
	for j, service := range partlyMatching {
		fmt.Printf("%v. %v \n", len(exactlyMatching)+j+1, service.Name)
	}
	fmt.Println("")
	fmt.Println("Which service would you like to check?")

	var input int
	var service apiv1.Service
	fmt.Scanln(&input)

	if 1 <= input && input <= len(exactlyMatching) {
		service = exactlyMatching[input-1]
	} else if len(exactlyMatching) < input && input <= len(partlyMatching) {
		service = partlyMatching[len(exactlyMatching)-input-1]
	}
	return service
}
