package main

import (
	"log"

	"github.com/bespinian/k8s-app-benchmarks/exitcodes/pkg/ebpf"
	"github.com/bespinian/k8s-app-benchmarks/exitcodes/pkg/proc"
	"golang.org/x/sys/unix"
)

func handleExec(execEvent ebpf.ExecEvent) {
	podName, _ := proc.PidToPodName(execEvent.PID)
	log.Printf("Exec event: pid: %d, comm: %s, pod name: %s", execEvent.PID, unix.ByteSliceToString(execEvent.Comm[:]), podName)
}

func handleExit(exitEvent ebpf.ExitEvent) {
	podName, _ := proc.PidToPodName(exitEvent.PID)
	log.Printf("Exit event: pid: %d, comm: %s, pod name: %s", exitEvent.PID, unix.ByteSliceToString(exitEvent.Comm[:]), podName)
}

func handleDone() {
	log.Printf(" ... exiting")
}

func main() {
	ebpf.Run(handleExec, handleExit, handleDone)
}
