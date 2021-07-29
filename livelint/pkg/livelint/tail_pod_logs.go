package livelint

import (
	"bytes"
	"context"
	"io"

	v1 "k8s.io/api/core/v1"
)

// tailPodLogs returns the last log messages of a pod
func (n *livelint) tailPodLogs(namespace, podName, containerName string, messageCount int64, previous bool) (string, error) {
	var logs string
	podLogOptions := v1.PodLogOptions{
		Container: containerName,
		TailLines: &messageCount,
		Previous: previous,
	}
	req := n.k8s.CoreV1().Pods(namespace).GetLogs(podName, &podLogOptions)
	podLogs, err := req.Stream(context.Background())
    if err != nil {
		return logs, err
    }
    defer podLogs.Close()

    buf := new(bytes.Buffer)
    _, err = io.Copy(buf, podLogs)
	
    if err != nil {
		return logs, err
    }
    logs = buf.String()

	return logs, nil
}
