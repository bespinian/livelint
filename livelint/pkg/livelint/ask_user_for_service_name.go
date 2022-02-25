package livelint

import (
	"fmt"
)

func askUserForServiceName() string {
	fmt.Println("")
	fmt.Println("Which service should expose this deployment?")

	var input string
	fmt.Scanln(&input)

	return input
}
