package promql

import "github.com/k8s-app-benchmarks/livelint/pkg/livelint"

type check struct {
	query string
}

func NewCheck(query string) livelint.Check {
	return &check{
		query: query,
	}
}

func (c *check) Run() (livelint.CheckResult, error) {
	result := livelint.CheckResult{
		DoesPass: false,
		Message:  "nope!",
	}

	return result, nil
}
