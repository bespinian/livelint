package livelint

// CheckResult is the result of a check.
type CheckResult struct {
	DoesPass bool   `json:"doesPass"`
	Message  string `json:"message"`
}
