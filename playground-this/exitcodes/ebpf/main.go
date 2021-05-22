package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang-11 SchedProcessExit ./bpf/sched_process_exit.c -- -I../headers

func main() {
}
