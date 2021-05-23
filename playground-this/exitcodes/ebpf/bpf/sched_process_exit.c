#include "common.h"
#include "bpf_helpers.h"

char __license[] SEC("license") = "Dual MIT/GPL";

struct event_t {
        u32 pid;
};

struct {
        __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
} events SEC(".maps");

SEC("tracepoint/sched/sched_process_exit")
int bpf_prog(void *ctx) {
  struct event_t event;
  event.pid = bpf_get_current_pid_tgid();
  bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &event, sizeof(event));
  return 0;
}
