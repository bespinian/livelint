package livelint

import (
	tea "github.com/charmbracelet/bubbletea"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Livelint struct {
	k8s    kubernetes.Interface
	config *rest.Config
	ui     *tea.Program
}

// New creates a livelint application.
func New(k8s kubernetes.Interface, config *rest.Config) *Livelint {
	l := &Livelint{
		k8s:    k8s,
		config: config,
		ui:     tea.NewProgram(initialModel()),
	}
	return l
}
