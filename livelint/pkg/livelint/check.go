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

	if len(nonRunningPods) > 0 {
		fmt.Println("NOK: There are non running pods")
		} else {
		fmt.Println("OK: All pods running")	
	}

	return nil
}
