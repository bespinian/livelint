package livelint

type CheckResult struct {
	HasFailed    bool
	Message      string
	Details      []string
	Instructions string
}
