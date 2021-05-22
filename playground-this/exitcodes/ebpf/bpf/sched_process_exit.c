#include "common.h"
#include "bpf_helpers.h"

char __license[] SEC("license") = "Dual MIT/GPL";

SEC("tracepoint/syscalls/sys_process_exit")
int bpf_prog(void *ctx) {
  char msg[] = "Process exited";
  bpf_trace_printk(msg, sizeof(msg));
  return 0;
}
