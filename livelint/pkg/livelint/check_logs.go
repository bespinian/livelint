package livelint

import (
	"strings"

	corev1 "k8s.io/api/core/v1"
)

const tailLineCount = 20

func (n *Livelint) checkContainerLogs(pod corev1.Pod, containerName string) CheckResult {
	namespace := pod.Namespace

	logs, err := n.tailPodLogs(namespace, pod.Name, containerName, tailLineCount, false)
	if err != nil {
		logs, err = n.tailPodLogs(namespace, pod.Name, containerName, tailLineCount, true)
		if err != nil {
			return CheckResult{
				HasFailed: true,
				Message:   "You cannot see the logs for the app",
			}
		}
	}

	if logs == "" {
		return CheckResult{
			HasFailed: true,
			Message:   "You cannot see the logs for the app",
		}
	}

	return CheckResult{
		Message:      "You can see the logs for the app",
		Details:      strings.Split(logs, "\n"),
		Instructions: "Fix the issue in the application",
	}
}
