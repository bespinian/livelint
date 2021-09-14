package exitcodes

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/bespinian/k8s-app-benchmarks/exitcodes/pkg/ebpf"
	"github.com/bespinian/k8s-app-benchmarks/exitcodes/pkg/proc"
	"golang.org/x/sys/unix"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type Exitcodes interface {
	HandleExec(execEvent ebpf.ExecEvent)
	HandleExit(exitEvent ebpf.ExitEvent)
	HandleDone()
}

type exitcodes struct {
	k8s kubernetes.Interface
}

// New creates a exitcodes application.
func New(k8s kubernetes.Interface) Exitcodes {
	e := &exitcodes{
		k8s: k8s,
	}
	return e
}

func (e *exitcodes) getPodData(podUID string) (v1.Pod, error) {
	pods, err := e.k8s.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	var resultPod v1.Pod
	podFound := false
	if err != nil {
		return resultPod, err
	}
	for _, pod := range pods.Items {
		if pod.UID == types.UID(podUID) {
			resultPod = pod
			podFound = true
		}
	}
	if !podFound {
		return resultPod, fmt.Errorf("Pod with UID %s not found!", podUID)
	}
	return resultPod, nil
}

func (e *exitcodes) getContainerData(pod v1.Pod, containerId string) (v1.Container, error) {
	var resultContainer v1.Container
	var resultContainerName string
	containerStatusFound := false
	containerFound := false
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if strings.Contains(containerStatus.ContainerID, containerId) {
			resultContainerName = containerStatus.Name
			containerStatusFound = true
			break
		}
	}
	if !containerStatusFound {
		return resultContainer, fmt.Errorf("Container status for ID %s not found in pod %s!", containerId, pod.Name)
	}
	for _, container := range pod.Spec.Containers {
		if container.Name == resultContainerName {
			resultContainer = container
			containerFound = true
		}
	}
	if !containerFound {
		return resultContainer, fmt.Errorf("Container with name %s not found in pod %s!", resultContainerName, pod.Name)
	}
	return resultContainer, nil
}

func (e *exitcodes) HandleExec(execEvent ebpf.ExecEvent) {
	podUID, _ := proc.PidToPodUid(execEvent.PID)
	containerId, _ := proc.PidToContainerId(execEvent.PID)
	pod, err := e.getPodData(podUID)
	var podName, podNamespace string
	var containerName string
	if err != nil {
		log.Printf("Pod lookup in k8s failed with Pod UID %s", podUID)
	} else {
		podName = pod.Name
		podNamespace = pod.Namespace
		container, err := e.getContainerData(pod, containerId)
		if err != nil {
			log.Printf("Container lookup in k8s failed with container ID %s", containerId)
		} else {
			containerName = container.Name
		}
	}
	log.Printf("Exec event: pid: %d, comm: %s, pod name: %s, container name: %s, namespace: %s", execEvent.PID, unix.ByteSliceToString(execEvent.Comm[:]), podName, containerName, podNamespace)
}

func (e *exitcodes) HandleExit(exitEvent ebpf.ExitEvent) {
	podUID, _ := proc.PidToPodUid(exitEvent.PID)
	containerId, _ := proc.PidToContainerId(exitEvent.PID)
	log.Printf("Exit event: pid: %d, comm: %s, pod UID: %s, container ID: %s", exitEvent.PID, unix.ByteSliceToString(exitEvent.Comm[:]), podUID, containerId)
}

func (e *exitcodes) HandleDone() {
	log.Printf(" ... exiting")
}
