#define BPF_NO_PRESERVE_ACCESS_INDEX 0
#include "vmlinux.h"
#include "common.h"
#include "bpf_helpers.h"

char __license[] SEC("license") = "Dual MIT/GPL";

struct event_t {
        u32 pid;
        u32 tgid;
        int ec;
	char comm[16];
        u32 ppid;
        u32 ptgid;
	char pcomm[16];
};

struct {
        __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
} events SEC(".maps");

SEC("tracepoint/sched/sched_process_exit")
int bpf_prog(void* ctx) {
  struct event_t event;
  struct task_struct *task = (struct task_struct*)bpf_get_current_task();
  struct task_struct *parent;
  int exitcode;
  bpf_probe_read(&exitcode, sizeof(task->exit_code), &task->exit_code);
  event.ec = exitcode >> 8;
  bpf_probe_read(&event.comm, sizeof(task->comm), &task->comm);
  bpf_probe_read(&event.pid, sizeof(task->pid), &task->pid);
  bpf_probe_read(&event.tgid, sizeof(task->tgid), &task->tgid);
  bpf_probe_read(&parent, sizeof(task->parent), &task->parent);
  bpf_probe_read(&event.ppid, sizeof(parent->pid), &parent->pid);
  bpf_probe_read(&event.ptgid, sizeof(parent->tgid), &parent->tgid);
  bpf_probe_read(&event.pcomm, sizeof(parent->comm), &parent->comm);
  bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &event, sizeof(event));
  return 0;
}
