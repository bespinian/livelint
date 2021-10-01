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
		color.New(color.FgRed).Add(color.Bold).Printf("✗ %s\n", r.Message)
	} else {
		color.New(color.FgGreen).Printf("✓ %s\n", r.Message)
	}

	if isVerbose {
		for _, d := range r.Details {
			fmt.Printf("  %s\n", d)
		}
	}

	if r.Instructions != "" {
		fmt.Println("")
		color.New(color.Bold).Printf("  → %s\n", r.Instructions)
		fmt.Println("")
	}
}
