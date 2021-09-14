package livelint

import (
	"fmt"

	"github.com/fatih/color"
)

// Check checks for potential issues with the deployment.
func (n *livelint) Check(namespace, deploymentName string, isVerbose bool) error {
	bold := color.New(color.Bold)
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	fmt.Println("")

	allPods, err := n.getPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error getting Pods: %w", err)
	}

	// Is there any PENDING Pod?
	hasPendingPods := areTherePendingPods(allPods)
	if hasPendingPods {
		clusterIsFull := askUserYesOrNo("Is the cluster full?")
		if clusterIsFull {
			bold.Println("Provision a bigger cluster")
			return nil
		}

		resourceQuotasAreHit := askUserYesOrNo("Are you hitting the ResourceQuota limits?")
		if resourceQuotasAreHit {
			bold.Println("Relax the ResourceQuota limits")
			return nil
		}

		isMountingPendingPVC := askUserYesOrNo("Are you mounting a PENDING PersistentVolumeClaim?")
		if isMountingPendingPVC {
			bold.Println("Fix the PersistentVolumeClaim")
			return nil
		}

		isPodAssignedToNode := askUserYesOrNo("Is the Pod assigned to the Node?")
		if isPodAssignedToNode {
			bold.Println("There is an issue with the Kubelet")
			return nil
		}

		bold.Println("There is an issue with the Scheduler")
		return nil
	}

	if isVerbose {
		green.Println("OK: No PENDING pods")
	}

	// Are the Pods RUNNING?
	hasOnlyRunningPods := areAllPodsRunning(allPods)
	if !hasOnlyRunningPods {
		for _, pod := range allPods {
			n.checkPodConditions(pod, isVerbose)

			problematicContainers := n.getProblematicContainers(pod)
			if len(problematicContainers) > 0 {
				boldRed.Printf("NOK: There are %d containers that are not started (and are not successfully terminated init containers) in pod %s\n", len(problematicContainers), pod.Name)

				if isVerbose {
					fmt.Printf("Trying to print logs from the first problematic %s\n", problematicContainers[0].Name)
				}

				logs, err := n.checkContainerLogs(pod, problematicContainers[0].Name)
				if err == nil {
					fmt.Println("App Logs:")
					fmt.Println("")
					fmt.Println(*logs)
					fmt.Println("")
					bold.Println("Fix the issue in the application")
					return nil
				}

				if isVerbose {
					fmt.Println("Could not get container logs")
				}

				for _, container := range problematicContainers {
					hasImagePullError, reason, message := n.checkImagePullErrors(pod, container.Name)
					if hasImagePullError {
						boldRed.Printf("NOK: Container %s has error pulling image (%s): %s\n",
							container.Name,
							container.Image,
							reason,
						)
						fmt.Println(message)
						fmt.Println("Verify that the image name, tag and registry are correct and that credentials are correct.")
					}

					isInCrashLoopBackOff, reason, message := n.checkCrashLoopBackOff(pod, container.Name)

					if isInCrashLoopBackOff {
						boldRed.Printf("NOK: Container %s is in CrashLoopBackOff: %s\n",
							container.Name,
							reason,
						)
						fmt.Println(message)
					}

					// Did you forget the CMD instruction in the Dockerfile?
					forgotCMD := n.checkForgottenCMDInDockerfile(pod)
					if forgotCMD {
						bold.Println("Fix the Dockerfile")
						return nil
					}

					// Is the Pod restarting frequently? Cycling between Running and CrashLoopBackOff?
					isRestartCycling, message := n.isRestartCycling(namespace, pod)
					if isRestartCycling {
						boldRed.Printf("NOK: Pod %s seems to be unhealthy. The last message was %s\n", pod.Name, message)
						bold.Println("Fix the liveness probe")
						return nil
					}

					bold.Println("Unknown state")
					return nil
				}

				// Is the pod status RunContainerError?
				if pod.Status.Phase == "RunContainerError" {
					bold.Println("The issue is likely to be with mounting volumes")
					return nil
				}
			}
		}
	}

	if isVerbose {
		green.Println("OK: All pods RUNNING")

	}

	podsAreReady := askUserYesOrNo("Are the Pods READY?")
	if !podsAreReady {
		readinessProbeIsFailing := askUserYesOrNo("Is the Readiness Probe Failing?")
		if readinessProbeIsFailing {
			bold.Println("Fix the readiness probe")
			return nil
		}

		bold.Println("Unknown state")
		return nil
	}

	// TODO: Can you access the app?
	// if no {
	//	TODO: Is the port exposed by container correct and listing on 0.0.0.0?
	//}
	// TODO: Can you see a list of endpoints?
	// if no {
	//	TODO: Is the Selector matching the Pod label?
	//	if yes {
	//		TODO: Does the Pod have an IP address assigned?
	//	}
	// }
	// TODO: Can you visit the app?
	// if no {
	//	TODO: Is the targetPort on the Service matching the containerPort in the Pod?
	// }
	// TODO: Can you see a list of Backends?
	// if no {
	//	TODO: Are the serviceName and servicePort matching the Service?
	//}
	// TODO: Can you visit the app?
	// if yes {
	//	TODO: The app should be working. Can you visit it from the public internet?
	//}

	fmt.Println("All checks finished")
	return nil
}
