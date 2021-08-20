package proc

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var podUidPattern = regexp.MustCompile(`\d+:.+:/kubepods/[^/]+/pod([^/]*)/`)
var containerIdPattern = regexp.MustCompile(`\d+:.+:/kubepods/[^/]+/pod[^/]+/([0-9a-f]{64})`)

func PidToPodUid(pid uint32) (string, error) {
	return scanProcCgroup(pid, podUidPattern)
}

func PidToContainerId(pid uint32) (string, error) {
	return scanProcCgroup(pid, containerIdPattern)
}

func scanProcCgroup(pid uint32, regexp *regexp.Regexp) (string, error) {

	f, err := os.Open(fmt.Sprintf("/proc/%d/cgroup", pid))
	if err != nil {
		// PID no longer exists
		return "", nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := regexp.FindStringSubmatch(line)
		if parts != nil {
			return parts[1], nil
		}
	}
	return "", nil

}
