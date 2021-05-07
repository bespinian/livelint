package livelint

// Check is a check to be performed.
type Check interface {
	Run() (CheckResult, error)
}
