package main

import (
	"log"

	"github.com/bespinian/k8s-app-benchmarks/exitcodes/pkg/ebpf"
	"github.com/bespinian/k8s-app-benchmarks/exitcodes/pkg/proc"
	"golang.org/x/sys/unix"
)

func handleExec(execEvent ebpf.ExecEvent) {
	podUID, _ := proc.PidToPodUid(execEvent.PID)
	containerId, _ := proc.PidToContainerId(execEvent.PID)
	log.Printf("Exec event: pid: %d, comm: %s, pod UID: %s, container ID: %s", execEvent.PID, unix.ByteSliceToString(execEvent.Comm[:]), podUID, containerId)
}

func handleExit(exitEvent ebpf.ExitEvent) {
	podUID, _ := proc.PidToPodUid(exitEvent.PID)
	containerId, _ := proc.PidToContainerId(exitEvent.PID)
	log.Printf("Exit event: pid: %d, comm: %s, pod UID: %s, container ID: %s", exitEvent.PID, unix.ByteSliceToString(exitEvent.Comm[:]), podUID, containerId)
}

func handleDone() {
	log.Printf(" ... exiting")
}

func main() {
	ebpf.Run(handleExec, handleExit, handleDone)
}
