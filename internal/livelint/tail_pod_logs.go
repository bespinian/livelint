package livelint

import (
	"bytes"
	"context"
	"fmt"
	"io"

	apiv1 "k8s.io/api/core/v1"
)

// tailPodLogs returns the last log messages of a pod.
func (n *Livelint) tailPodLogs(namespace, podName, containerName string, tailLines int64, previous bool) (string, error) {
	var logs string
	podLogOptions := apiv1.PodLogOptions{
		Container: containerName,
		TailLines: &tailLines,
		Previous:  previous,
	}
	req := n.k8s.CoreV1().Pods(namespace).GetLogs(podName, &podLogOptions)
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		return logs, fmt.Errorf("error getting request stream: %w", err)
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return logs, fmt.Errorf("error copying logs buffer: %w", err)
	}

	logs = buf.String()

	return logs, nil
}
