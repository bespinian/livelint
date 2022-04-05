package livelint

import (
	"fmt"

	"github.com/fatih/color"
)

type CheckResult struct {
	HasFailed    bool
	Message      string
	Details      []string
	Instructions string
}

func (r CheckResult) PrettyPrint(isVerbose bool) {
	if r.HasFailed {
		redBold := color.New(color.FgRed).Add(color.Bold)
		redBold.Printf("✖ %s\n", r.Message)
	} else {
		green := color.New(color.FgGreen)
		green.Printf("✔ %s\n", r.Message)
	}

	if isVerbose {
		for _, d := range r.Details {
			fmt.Printf("  %s\n", d)
		}
	}

	if r.Instructions != "" {
		bold := color.New(color.Bold)
		fmt.Println("")
		bold.Printf("  → %s\n", r.Instructions)
		fmt.Println("")
	}
}
