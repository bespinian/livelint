package livelint

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Livelint represents a livelint application.
type Livelint interface {
	RunChecks(namespace, deploymentName string, isVerbose bool) error
}

type livelint struct {
	k8s    kubernetes.Interface
	config *rest.Config
}

// New creates a livelint application.
func New(k8s kubernetes.Interface, config *rest.Config) Livelint {
	l := &livelint{
		k8s:    k8s,
		config: config,
	}
	return l
}
