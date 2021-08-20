package main

import (
	"log"

	"github.com/bespinian/k8s-app-benchmarks/exitcodes/pkg/ebpf"
	"golang.org/x/sys/unix"
)

func main() {
	ebpf := ebpf.New()
	for {
		select {
		case execEvent := <-ebpf.ExecEvents():
			log.Printf("Exec event: ppid: %d, ptgid: %d, pcomm: %s, pid: %d, tgid: %d, comm: %s, nspid: %d", execEvent.PPID, execEvent.PTGID, unix.ByteSliceToString(execEvent.PComm[:]), execEvent.PID, execEvent.TGID, unix.ByteSliceToString(execEvent.Comm[:]), execEvent.NSPID)
		case exitEvent := <-ebpf.ExitEvents():
			log.Printf("Exit event: ppid: %d, ptgid: %d, pcomm: %s, pid: %d, tgid: %d, exit code: %d, comm: %s, nspid: %d", exitEvent.PPID, exitEvent.PTGID, unix.ByteSliceToString(exitEvent.PComm[:]), exitEvent.PID, exitEvent.TGID, exitEvent.Ec, unix.ByteSliceToString(exitEvent.Comm[:]), exitEvent.NSPID)
		case <-ebpf.DoneEvents():
			return

		}
	}
}
