package livelint

import (
	"net/http"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Livelint struct {
	K8s    kubernetes.Interface
	config *rest.Config
	HTTP   *http.Client
	ui     UserInteraction
}

// New creates a livelint application.
func New(k8s kubernetes.Interface, config *rest.Config, ui UserInteraction) *Livelint {
	l := &Livelint{
		K8s:    k8s,
		config: config,
		HTTP:   http.DefaultClient,
		ui:     ui,
	}
	return l
}
