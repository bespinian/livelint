package livelint

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	apiv1 "k8s.io/api/core/v1"
)

const tailLineCount = 20

func (n *Livelint) checkContainerLogs(pod apiv1.Pod, containerName string) CheckResult {
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
		Details:      strings.Split(fmt.Sprintf("Logs for Pod %s:\n", pod.Name)+logs, "\n"),
		Instructions: "Fix the issue in the application",
	}
}

// tailPodLogs returns the last log messages of a pod.
func (n *Livelint) tailPodLogs(namespace, podName, containerName string, tailLines int64, previous bool) (string, error) {
	podLogOptions := apiv1.PodLogOptions{
		Container: containerName,
		TailLines: &tailLines,
		Previous:  previous,
	}
	req := n.k8s.CoreV1().Pods(namespace).GetLogs(podName, &podLogOptions)
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		return "", fmt.Errorf("error getting request stream: %w", err)
	}
	defer podLogs.Close()

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", fmt.Errorf("error copying logs buffer: %w", err)
	}

	return buf.String(), nil
}
