package livelint

import "k8s.io/client-go/kubernetes"

// Livelint represents a livelint application.
type Livelint interface {
	RunChecks(namespace, deploymentName string, isVerbose bool) error
}

type livelint struct {
	k8s kubernetes.Interface
}

// New creates a livelint application.
func New(k8s kubernetes.Interface) Livelint {
	l := &livelint{
		k8s: k8s,
	}
	return l
}
