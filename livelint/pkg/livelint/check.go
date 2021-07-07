package livelint

import "fmt"

// Topo lists the topology of a connection path.
func (n *livelint) Check(namespace, deploymentName string) error {
	pendingPods, err := n.getPendingPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error getting pending pods: %w", err)
	}

	if len(pendingPods) > 0 {
		fmt.Println("There are pending pods")
	}

	return nil
}
