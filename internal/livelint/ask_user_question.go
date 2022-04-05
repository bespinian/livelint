package livelint

import (
	"fmt"
)

func askUserQuestion(question string) string {
	fmt.Println("")
	fmt.Println(question)

	var response string
	fmt.Scanln(&response)

	return response
}
