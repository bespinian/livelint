package ebpf

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/perf"
	"golang.org/x/sys/unix"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang-11 SchedProcess ./bpf/sched_process.c -- -I./headers

type ExecEvent struct {
	PID   uint32
	TGID  uint32
	Comm  [16]byte
	PPID  uint32
	PTGID uint32
	PComm [16]byte
	NSPID uint32
}
type ExitEvent struct {
	PID   uint32
	TGID  uint32
	Ec    int32
	Comm  [16]byte
	PPID  uint32
	PTGID uint32
	PComm [16]byte
	NSPID uint32
}

func Run(onExec func(e ExecEvent), onExit func(e ExitEvent), onDone func()) {
	// Subscribe to signals for terminating the program.
	stopper := make(chan os.Signal, 1)
	signal.Notify(stopper, os.Interrupt, syscall.SIGTERM)

	// Increase the rlimit of the current process to provide sufficient space
	// for locking memory for the eBPF map.
	if err := unix.Setrlimit(unix.RLIMIT_MEMLOCK, &unix.Rlimit{
		Cur: unix.RLIM_INFINITY,
		Max: unix.RLIM_INFINITY,
	}); err != nil {
		log.Fatalf("failed to set temporary rlimit: %v", err)
	}

	log.Println("Loading objects to kernel..")
	// Load pre-compiled programs and maps into the kernel.
	objs := SchedProcessObjects{}
	if err := LoadSchedProcessObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	// Open readers from userspace into the event arrays
	// created earlier.
	rdExec, err := perf.NewReader(objs.ExecEvents, os.Getpagesize())
	if err != nil {
		log.Fatalf("creating exec event reader: %s", err)
	}
	defer rdExec.Close()

	rdExit, err := perf.NewReader(objs.ExitEvents, os.Getpagesize())
	if err != nil {
		log.Fatalf("creating exit event reader: %s", err)
	}
	defer rdExit.Close()

	// Close the readers when the process receives a signal, which will exit
	// the read loops.
	go func() {
		<-stopper
		rdExec.Close()
		rdExit.Close()
	}()

	log.Println("Setting up exec event tracepoint..")
	tpExec, err := link.Tracepoint("sched", "sched_process_exec", objs.BpfProcessExec)
	if err != nil {
		log.Fatalf("opening exec tracepoint: %s", err)
	}
	defer tpExec.Close()

	log.Println("Setting up exit event tracepoint..")
	tpExit, err := link.Tracepoint("sched", "sched_process_exit", objs.BpfProcessExit)
	if err != nil {
		log.Fatalf("opening exit tracepoint: %s", err)
	}
	defer tpExit.Close()

	log.Println("Setting up exec event callback..")
	go func() {
		for {
			record, err := rdExec.Read()
			var execEvent ExecEvent
			if err != nil {
				if perf.IsClosed(err) {
					log.Println("Received signal, exiting exec read loop ...")
					onDone()
					return
				}
				log.Fatalf("reading from exec reader: %s", err)
			}

			// Parse the perf event entry into an ExecEvent struct.
			if err := binary.Read(bytes.NewBuffer(record.RawSample), binary.LittleEndian, &execEvent); err != nil {
				log.Printf("parsing exec event: %s", err)
				continue
			}
			onExec(execEvent)
		}
	}()

	log.Println("Setting up exit event callback..")
	go func() {
		for {
			record, err := rdExit.Read()
			var exitEvent ExitEvent
			if err != nil {
				if perf.IsClosed(err) {
					log.Println("Received signal, exiting exit read loop ...")
					onDone()
					return
				}
				log.Fatalf("reading from exit reader: %s", err)
			}

			// Parse the perf event entry into an ExitEvent struct.
			if err := binary.Read(bytes.NewBuffer(record.RawSample), binary.LittleEndian, &exitEvent); err != nil {
				log.Printf("parsing exit event: %s", err)
				continue
			}
			onExit(exitEvent)
		}
	}()
}
