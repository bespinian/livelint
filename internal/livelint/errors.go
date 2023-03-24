package livelint

import (
	"errors"
)

var (
	errNamespaceUndefined      = errors.New("no namespace defined")
	errDeploymentNameUndefined = errors.New("no deployment defined")
)
