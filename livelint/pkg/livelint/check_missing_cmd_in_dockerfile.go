package livelint

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) checkForgottenCMDInDockerfile(pod corev1.Pod) bool {
	msg := fmt.Sprintf("Did you forget the CMD instruction of %q in the Dockerfile?", pod.Name)
	hasChecked := askUserYesOrNo(msg)
	return hasChecked
}
