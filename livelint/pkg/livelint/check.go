package livelint

import "fmt"

// Topo lists the topology of a connection path.
func (n *livelint) Check(namespace, deploymentName string) error {
	pendingPods, err := n.getPendingPods(namespace, deploymentName)

	if err != nil {
		return fmt.Errorf("error getting pending pods: %w", err)
	}

	if len(pendingPods) > 0 {
		fmt.Println("NOK: There are pending pods")
		} else {
		fmt.Println("OK: No pending pods")	
	}

	nonRunningPods, err := n.getNonRunningPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error getting non-running pods: %w", err)
	}

	if len(nonRunningPods) > 0 {
		fmt.Println("NOK: There are non running pods")

		} else {
		fmt.Println("OK: All pods running")	
	}

	allPods, err := n.getPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error getting pods: %w", err)
	}
	
	for i := 0; i < len(allPods); i++ {
		pod := allPods[i]
		nonStartedContainerNames := n.getNonStartedContainerNames(pod)

		if len(nonStartedContainerNames) > 0 {
			fmt.Printf("NOK: There are %d containers that are not started in pod %s\n", len(nonStartedContainerNames), pod.Name)
			fmt.Printf("Trying to print logs from the first non started container %s\n", nonStartedContainerNames[0])
			
			logs, err := n.checkContainerLogs(pod, nonStartedContainerNames[0])
			if (err != nil)  {
				fmt.Printf("error getting container logs: %s", err.Error())
			} else {
				fmt.Println(*logs)
				return nil
			}
		}
	}

	return nil
}
