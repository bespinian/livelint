package livelint

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Livelint struct {
	k8s    kubernetes.Interface
	config *rest.Config
}

// New creates a livelint application.
func New(k8s kubernetes.Interface, config *rest.Config) *Livelint {
	l := &Livelint{
		k8s:    k8s,
		config: config,
	}
	return l
}
