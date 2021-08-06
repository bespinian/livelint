package livelint

import (
	"fmt"

	"github.com/fatih/color"
)

// Topo lists the topology of a connection path.
func (n *livelint) Check(namespace, deploymentName string, isVerbose bool) error {
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	fmt.Println("")

	pendingPods, err := n.getPendingPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error getting pending pods: %w", err)
	}
	if len(pendingPods) > 0 {
		boldRed.Println(("NOK: There are pending pods"))
	} else if isVerbose {
		green.Println("OK: No pending pods")
	}

	nonRunningPods, err := n.getNonRunningPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error getting non-running pods: %w", err)
	}

	if len(nonRunningPods) > 0 {
		boldRed.Println(("NOK: There are non running pods"))
	} else if isVerbose {
		green.Println("OK: All pods running")
	}

	allPods, err := n.getPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error getting pods: %w", err)
	}

	for _, pod := range allPods {
		n.checkPodConditions(pod, isVerbose)

		nonStartedContainers := n.getNonStartedContainers(pod)

		if len(nonStartedContainers) > 0 {
			boldRed.Printf("NOK: There are %d containers that are not started in pod %s\n", len(nonStartedContainers), pod.Name)
			if isVerbose {
				fmt.Printf("Trying to print logs from the first non started container %s\n", nonStartedContainers[0].Name)
			}

			logs, err := n.checkContainerLogs(pod, nonStartedContainers[0].Name)
			if err == nil {
				fmt.Println(*logs)
				return nil
			}
			fmt.Println("Could not get container logs")

			for _, container := range nonStartedContainers {
				hasImagePullError, reason, message := n.checkImagePullErrors(pod, container.Name)
				if hasImagePullError {
					fmt.Printf("Container %s has error pulling image (%s): %s\n",
						container.Name,
						container.Image,
						reason,
					)
					fmt.Println(message)
					fmt.Println("Verify that the image name, tag and registry are correct and that credentials are correct.")
				}
			}
		}
	}

	fmt.Println("All checks finished")
	return nil
}
