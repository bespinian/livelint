package livelint

import (
	"fmt"
	"strings"

	apiv1 "k8s.io/api/core/v1"
)

func checkCreateContainerConfigErrors(pod apiv1.Pod, container apiv1.Container) CheckResult {
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Name != container.Name {
			continue
		}

		if containerStatus.State.Waiting != nil &&
			containerStatus.State.Waiting.Reason == "CreateContainerConfigError" {
			missingResource := "ConfigMap"
			if strings.Contains(containerStatus.State.Waiting.Message, "secret") {
				missingResource = "Secret"
			}

			return CheckResult{
				HasFailed:    true,
				Message:      fmt.Sprintf("A Pod is in status %s", containerStatus.State.Waiting.Reason),
				Details:      []string{containerStatus.State.Waiting.Message},
				Instructions: fmt.Sprintf("Create the missing %s", missingResource),
			}
		}
	}

	return CheckResult{
		Message: "There are no container config errors",
	}
}
