#define BPF_NO_PRESERVE_ACCESS_INDEX 0
#include "vmlinux.h"
#include "common.h"
#include "bpf_helpers.h"

char __license[] SEC("license") = "Dual MIT/GPL";

struct event_t {
        u32 pid;
        int ec;
};

struct {
        __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
} events SEC(".maps");

SEC("tracepoint/sched/sched_process_exit")
int bpf_prog(void* ctx) {
  struct event_t event;
  event.pid = bpf_get_current_pid_tgid();
  struct task_struct *task = (struct task_struct*)bpf_get_current_task();
  int exitcode;
  bpf_probe_read(&exitcode, sizeof(task->exit_code), &task->exit_code);
  event.ec = exitcode >> 8;
  bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &event, sizeof(event));
  return 0;
}
