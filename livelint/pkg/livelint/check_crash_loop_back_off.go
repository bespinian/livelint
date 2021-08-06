package livelint

import (
	corev1 "k8s.io/api/core/v1"
)

func (n *livelint) checkCrashLoopBackOff(pod corev1.Pod, containerName string) (bool, string, string) {
	for _, cs := range pod.Status.ContainerStatuses {
		if cs.Name != containerName {
			continue
		}
		
		if cs.State.Waiting != nil &&
			cs.State.Waiting.Reason == "CrashLoopBackOff" {
			return true, cs.State.Waiting.Reason,
				cs.State.Waiting.Message
		}
	}
	return false, "", ""
}
