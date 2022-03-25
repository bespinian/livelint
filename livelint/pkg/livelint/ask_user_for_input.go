package livelint

import (
	"fmt"
)

func askUserForInput(question string) string {
	fmt.Println("")
	fmt.Println(question)

	var input string
	fmt.Scanln(&input)

	return input
}
