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
			nonRunningContainers := getNonRunningContainers(pod)
			if len(nonRunningContainers) > 0 {

				// Can you see the logs for the app?
				result = n.checkContainerLogs(pod, nonRunningContainers[0].Name)
				result.PrettyPrint(isVerbose)
				if !result.HasFailed {
					return nil
				}

				// Is the Pod status ImagePullBackOff?
				result = checkImagePullErrors(pod, nonRunningContainers[0].Name)
				result.PrettyPrint(isVerbose)
				if result.HasFailed {

					// Is the name of the image correct?
					result = checkIsImageNameCorrect()
					result.PrettyPrint(isVerbose)
					if result.HasFailed {
						return nil
					}

					// Is the image tag valid? Does it exist?
					result = checkIsImageTagValid()
					result.PrettyPrint(isVerbose)
					if result.HasFailed {
						return nil
					}

					// Are you pulling images from a private registry?
					result = checkIsPullingFromPrivateRegistry()
					result.PrettyPrint(isVerbose)

					return nil
				}

				// Is the Pod status CrashLoopBackOff?
				result = checkCrashLoopBackOff(pod, nonRunningContainers[0].Name)
				result.PrettyPrint(isVerbose)
				if result.HasFailed {

					// Did you inspect the logs and fix the crashing app?
					result = checkDidInspectLogsAndFix()
					result.PrettyPrint(isVerbose)
					if result.HasFailed {
						return nil
					}

					// Did you forget the CMD instruction in the Dockerfile?
					result = checkForgottenCMDInDockerfile()
					if result.HasFailed {
						return nil
					}

					// Is the Pod restarting frequently? Cycling between Running and CrashLoopBackOff?
					result = n.checkIsRestartCycling(pod)
					result.PrettyPrint(isVerbose)

					return nil
				}

				// Is the pod status RunContainerError?
				result = checkRunContainerError(pod)
				result.PrettyPrint(isVerbose)
				if result.HasFailed {

					// Is there any container running?
					result = checkAreThereRunningContainers()
					result.PrettyPrint(isVerbose)

					return nil
				}

				return nil
			}
		}
	}

	// Are the Pods READY?
	result = checkAreAllPodsReady(allPods)
	result.PrettyPrint(isVerbose)
	if result.HasFailed {

		// Is the Readiness probe failing?
		result = n.checkReadinessProbe(allPods)
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
