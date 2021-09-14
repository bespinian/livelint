package livelint

import (
	"fmt"
	"strings"
)

func askUserYesOrNo(msg string) bool {
	fmt.Printf("%s (y/N)\n", msg)

	var input string
	fmt.Scanln(&input)

	normalizedInput := strings.ToLower(input)
	if normalizedInput == "y" || normalizedInput == "yes" {
		return true
	}

	return false
}
