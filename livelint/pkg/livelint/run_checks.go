package livelint

import (
	"fmt"

	"github.com/fatih/color"
)

// RunChecks checks for potential issues with the deployment.
func (n *livelint) RunChecks(namespace, deploymentName string, isVerbose bool) error {
	greenBold := color.New(color.FgGreen).Add(color.Bold)

	fmt.Println("")
	fmt.Printf("Checking deployment %s in namespace %s...\n", deploymentName, namespace)
	fmt.Println("")

	allPods, err := n.getPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error getting Pods: %w", err)
	}

	// Is there any PENDING Pod?
	result := checkAreTherePendingPods(allPods)
	result.PrettyPrint(isVerbose)
	if result.HasFailed {

		// Is the cluster full?
		result = checkIsClusterFull()
		result.PrettyPrint(isVerbose)
		if result.HasFailed {
			return nil
		}

		// Are you hitting the ResourceQuota limits ?
		result = checkAreResourceQuotasHit()
		result.PrettyPrint(isVerbose)
		if result.HasFailed {
			return nil
		}

		// Are you mounting a PENDING PersistentVolumeClaim?
		result = checkIsMountingPendingPVC()
		result.PrettyPrint(isVerbose)
		if result.HasFailed {
			return nil
		}

		// Is the Pod assigned to the Node?
		result = checkIsPodAssignedToNode()
		result.PrettyPrint(isVerbose)

		return nil
	}

	// Are the Pods RUNNING?
	result = checkAreAllPodsRunning(allPods)
	result.PrettyPrint(isVerbose)
	if result.HasFailed {
		for _, pod := range allPods {
			problematicContainers := n.getProblematicContainers(pod)
			if len(problematicContainers) > 0 {
				logs, err := n.checkContainerLogs(pod, problematicContainers[0].Name)
				if err == nil {
					fmt.Println("App Logs:")
					fmt.Println("")
					fmt.Println(*logs)
					fmt.Println("")
					fmt.Println("Fix the issue in the application")
					return nil
				}

				for _, container := range problematicContainers {
					hasImagePullError, _, message := n.checkImagePullErrors(pod, container.Name)
					if hasImagePullError {
						fmt.Println(message)
						fmt.Println("Verify that the image name, tag and registry are correct and that credentials are correct.")
					}

					isInCrashLoopBackOff, _, message := n.checkCrashLoopBackOff(pod, container.Name)

					if isInCrashLoopBackOff {
						fmt.Println(message)
					}

					// Did you forget the CMD instruction in the Dockerfile?
					result = checkForgottenCMDInDockerfile()
					if result.HasFailed {
						return nil
					}

					// Is the Pod restarting frequently? Cycling between Running and CrashLoopBackOff?
					isBackingOff, hasUnhealthyEvents, _ := n.isRestartCycling(pod)
					if isBackingOff {
						if hasUnhealthyEvents {
							return nil
						}

						return nil
					}

					return nil
				}

				// Is the pod status RunContainerError?
				if pod.Status.Phase == "RunContainerError" {
					return nil
				}
			}
		}
	}

	// Are the Pods READY?
	result = checkAreAllPodsReady(allPods)
	result.PrettyPrint(isVerbose)
	if result.HasFailed {

		// Is the Readiness probe failing?
		result := n.checkReadinessProbe(allPods)
		result.PrettyPrint(isVerbose)

		return nil
	}

	// Can you access the app?
	result = checkCanAccessApp()
	if result.HasFailed {

		// Is the port exposed by container correct and listing on 0.0.0.0?
		result = checkIsPortExposedCorrectly()
		result.PrettyPrint(isVerbose)

		return nil
	}

	fmt.Println("")
	greenBold.Println("Pods are running correctly")
	fmt.Println("")

	// Can you see a list of endpoints?
	result = checkCanSeeEndpoints()
	result.PrettyPrint(isVerbose)
	if result.HasFailed {

		// Is the Selector matching the Pod label?
		result = checkIsSelectorMatchingPodLabel()
		result.PrettyPrint(isVerbose)
		if result.HasFailed {
			return nil
		}

		// Does the Pod have an IP address assigned?
		result = checkPodHasIPAddressAssigned()
		result.PrettyPrint(isVerbose)

		return nil
	}

	// Can you visit the app?
	result = checkCanVisitServiceApp()
	result.PrettyPrint(isVerbose)
	if result.HasFailed {

		// Is the targetPort on the Service matching the containerPort in the Pod?
		result = checkTargetPortMatchesContainerPort()
		result.PrettyPrint(isVerbose)

		return nil
	}

	fmt.Println("")
	greenBold.Println("The Service is running correctly")
	fmt.Println("")

	// Can you see a list of Backends?
	result = checkCanSeeBackends()
	result.PrettyPrint(isVerbose)
	if result.HasFailed {

		// Are the serviceName and servicePort matching the Service?
		result = checkServiceNameAndPortMatchService()
		result.PrettyPrint(isVerbose)

		return nil
	}

	// Can you visit the app?
	result = checkCanVisitIngressApp()
	result.PrettyPrint(isVerbose)
	if result.HasFailed {
		return nil
	}

	fmt.Println("")
	greenBold.Println("The Ingress is running correctly")
	fmt.Println("")

	// The app should be running. Can you visit it from the public internet?
	result = checkCanVisitPublicApp()
	result.PrettyPrint(isVerbose)
	if result.HasFailed {
		return nil
	}

	fmt.Println("")
	fmt.Println("All checks finished")
	return nil
}
