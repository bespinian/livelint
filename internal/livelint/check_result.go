package livelint

type CheckResult struct {
	HasFailed    bool
	HasWarning   bool
	Message      string
	Details      []string
	Instructions string
}
