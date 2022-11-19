package livelint

import (
	"net/http"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Livelint struct {
	k8s    kubernetes.Interface
	config *rest.Config
	http   *http.Client
	ui     UserInteraction
}

// New creates a livelint application.
func New(k8s kubernetes.Interface, config *rest.Config, ui UserInteraction) *Livelint {
	l := &Livelint{
		k8s:    k8s,
		config: config,
		http:   http.DefaultClient,
		ui:     ui,
	}
	return l
}
