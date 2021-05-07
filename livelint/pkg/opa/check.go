package opa

import "github.com/k8s-app-benchmarks/livelint/pkg/livelint"

type check struct {
	rego string
}

func NewCheck(rego string) livelint.Check {
	return &check{
		rego: rego,
	}
}

func (c *check) Run() (livelint.CheckResult, error) {
	result := livelint.CheckResult{
		DoesPass: true,
		Message:  "looks good",
	}

	return result, nil
}
