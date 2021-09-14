package livelint

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) checkInspectCrashingApp(pod corev1.Pod) bool {
	msg := fmt.Sprintf("Did you inspect the logs of %q and fix the crashing app", pod.Name)
	hasChecked := askUserYesOrNo(msg)
	return hasChecked
}
