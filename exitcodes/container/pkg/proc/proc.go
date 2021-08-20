package proc

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func PidToPodName(pid uint32) (string, error) {
	kubePattern := regexp.MustCompile(`\d+:.+:/kubepods/[^/]+/pod[^/]+/([0-9a-f]{64})`)
	f, err := os.Open(fmt.Sprintf("/proc/%d/cgroup", pid))
	if err != nil {
		// PID no longer exists
		return "", nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := kubePattern.FindStringSubmatch(line)
		if parts != nil {
			return parts[1], nil
		}
	}
	return "", nil
}
