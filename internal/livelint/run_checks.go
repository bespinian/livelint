package livelint

import (
	"fmt"
)

// RunChecks checks for potential issues with a deployment.
func (n *Livelint) RunChecks(namespace, deploymentName string, isVerbose bool) error {
	if namespace == "" {
		return fmt.Errorf("error running checks: %w", errNamespaceUndefined)
	}
	if deploymentName == "" {
		return errDeploymentNameUndefined
	}

	// nolint:govet
	n.ui.DisplayContext(fmt.Sprintf("Checking Deployment %q in Namespace %q", deploymentName, namespace))
	n.ui.DisplayCheckStart("Checking Deployment Pods")

	allPods, err := n.getDeploymentPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error running checks: %w", err)
	}

	// Is the number of running pods correct ?
	result := n.CheckIsNumberOfPodsMatching(namespace, deploymentName)
	n.ui.DisplayCheckResult(result)
	if result.HasFailed {
		// Are you hitting the ResourceQuota limits?
		result = n.CheckAreResourceQuotasHit(namespace, deploymentName)
		n.ui.DisplayCheckResult(result)
		if result.HasFailed {
			return nil
		}
	}

	// Is there any PENDING Pod?
	result = checkAreTherePendingPods(allPods)
	n.ui.DisplayCheckResult(result)
	if result.HasFailed {

		// Is the cluster full?
		result = n.CheckIsClusterFull(allPods)
		n.ui.DisplayCheckResult(result)
		if result.HasFailed {
			return nil
		}

		// Are you mounting a PENDING PersistentVolumeClaim?
		result = n.CheckIsMountingPendingPVC(allPods, namespace)
		n.ui.DisplayCheckResult(result)
		if result.HasFailed {
			return nil
		}

		// Is the Pod assigned to the Node?
		result = checkIsPodAssignedToNode(allPods)
		n.ui.DisplayCheckResult(result)

		return nil
	}

	// Are any Pods restart cycling
	result = n.CheckAreThereRestartCyclingPods(allPods)
	n.ui.DisplayCheckResult(result)
	if result.HasFailed {
		return nil
	}

	// Are the Pods RUNNING?
	result = checkAreAllPodsRunning(allPods)
	n.ui.DisplayCheckResult(result)
	if result.HasFailed {
		for _, pod := range allPods {
			nonRunningContainers := getNonRunningContainers(pod)
			if len(nonRunningContainers) > 0 {
				nonRunningContainer := nonRunningContainers[0]

				// Is the Pod status InvalidImageName?
				result = checkInvalidImageName(pod, nonRunningContainer)
				n.ui.DisplayCheckResult(result)
				if result.HasFailed {
					return nil
				}

				// Is the Pod status ImagePullBackOff?
				result = checkImagePullErrors(pod, nonRunningContainer)
				n.ui.DisplayCheckResult(result)
				if result.HasFailed {

					// Is the name of the image correct?
					result = n.checkIsImageNameCorrect(nonRunningContainer)
					n.ui.DisplayCheckResult(result)
					if result.HasFailed {
						return nil
					}

					// Is the image tag valid? Does it exist?
					result = n.checkDoesImageTagExist(nonRunningContainer)
					n.ui.DisplayCheckResult(result)
					if result.HasFailed {
						return nil
					}

					// Are you pulling images from a private registry?
					result = n.checkIsPullingFromPrivateRegistry(nonRunningContainer.Image)
					n.ui.DisplayCheckResult(result)

					return nil
				}

				// Is the Pod status CrashLoopBackOff?
				result = checkCrashLoopBackOff(pod, nonRunningContainer.Name)
				n.ui.DisplayCheckResult(result)
				if result.HasFailed {

					// Did you inspect the logs and fix the crashing app?
					result = n.checkDidInspectLogsAndFix()
					n.ui.DisplayCheckResult(result)
					if result.HasFailed {

						// Can you see the logs for the app?
						result = n.CheckContainerLogs(pod, nonRunningContainer.Name)
						n.ui.DisplayCheckResult(result)
						if !result.HasFailed {
							return nil
						}
					}

					// Did you forget the CMD instruction in the Dockerfile?
					result = n.CheckForgottenCMDInDockerfile(nonRunningContainer)
					n.ui.DisplayCheckResult(result)

					return nil
				}

				// Is the pod status RunContainerError?
				result = checkIsContainerCreating(pod)
				n.ui.DisplayCheckResult(result)
				if result.HasFailed {

					// Is there any container running?
					result = checkAreThereRunningContainers(pod)
					n.ui.DisplayCheckResult(result)
					return nil
				}

				// Are there failing volume mounts
				result = n.CheckFailedMount(pod)
				n.ui.DisplayCheckResult(result)

				// Is there any container running?
				result = checkAreThereRunningContainers(pod)
				n.ui.DisplayCheckResult(result)

				return nil
			}
		}
	}

	// Are the Pods READY?
	result = checkAreAllPodsReady(allPods)
	n.ui.DisplayCheckResult(result)
	if result.HasFailed {

		// Is the Readiness probe failing?
		result = n.CheckReadinessProbe(allPods)
		n.ui.DisplayCheckResult(result)

		return nil
	}

	// Can you access the app?
	result = n.checkCanAccessApp(allPods)
	n.ui.DisplayCheckResult(result)
	if result.HasFailed {

		// Is the port exposed by container correct and listening on 0.0.0.0?
		result = n.checkIsPortExposedCorrectly()
		n.ui.DisplayCheckResult(result)
		return nil
	}

	n.ui.DisplayCheckCompletion("Pods are running correctly", success)

	n.ui.DisplayCheckStart("Checking Services")

	allServices, partlyMatchingServices, err := n.getServices(namespace, deploymentName)
	if err != nil {
		n.ui.DisplayError(fmt.Errorf("error getting services in namespace %s: %w", namespace, err))
	}

	allServices = append(allServices, partlyMatchingServices...)

	if len(allServices) < 1 {
		result = CheckResult{
			HasFailed: true,
			Message:   "No services match the deployment's matchSelector.",
		}
		n.ui.DisplayCheckResult(result)
		return nil
	}

	for _, service := range allServices {

		// Can you see a list of endpoints?
		result = n.CheckCanSeeEndpoints(service.Name, namespace)
		n.ui.DisplayCheckResult(result)
		if result.HasFailed {
			// Is the selector matching the right pod label?
			result = n.checkIsSelectorMatchPodLabel(namespace, service.Name, allPods)
			n.ui.DisplayCheckResult(result)
			if result.HasFailed {
				// Does the Pod have an IP address assigned?
				result = checkPodHasIPAddressAssigned(allPods)
				n.ui.DisplayCheckResult(result)
			}
			return nil
		}

		// Can you visit the app?
		result = n.checkCanVisitServiceApp(service)
		n.ui.DisplayCheckResult(result)
		if result.HasFailed {

			// Is the targetPort on the Service matching the containerPort in the Pod?
			result = n.checkTargetPortMatchesContainerPort(allPods, service.Name, namespace)
			n.ui.DisplayCheckResult(result)

			return nil
		}
		n.ui.DisplayCheckCompletion(fmt.Sprintf("The Service %s is running correctly", service.Name), success)
	}

	n.ui.DisplayCheckStart("Checking Ingresses")

	ingresses, err := n.getIngressesFromServices(namespace, allServices)
	if err != nil {
		n.ui.DisplayError(fmt.Errorf("error getting ingresses in namespace %s: %w", namespace, err))
	}

	if len(ingresses) < 1 {
		result = CheckResult{
			HasFailed: true,
			Message:   "No Ingresses match any of the previously detected services names and ports",
		}
		n.ui.DisplayCheckResult(result)
		return nil
	}

	ingressClasses, err := n.getIngressClasses()
	if err != nil {
		return fmt.Errorf("error getting ingress classes: %w", err)
	}

	for _, ingress := range ingresses {
		// Can you see a list of Backends?
		result = n.CheckCanSeeBackends(ingress, namespace)
		n.ui.DisplayCheckResult(result)
		if result.HasFailed {
			return nil
		}

		result = checkHasValidIngressClass(ingress, ingressClasses)
		n.ui.DisplayCheckResult(result)
		if result.HasFailed {
			return nil
		}

		result = n.checkCanAccessAppFromIngressController(ingress, ingressClasses)
		n.ui.DisplayCheckResult(result)
		if result.HasFailed {
			return nil
		}

		n.ui.DisplayCheckCompletion(fmt.Sprintf("The Ingress %s is set up correctly", ingress.Name), success)
	}

	// Can you visit the app?
	n.ui.DisplayCheckStart("Checking App Connectivity")

	// The app should be running. Can you visit it from the public internet?
	result = n.CheckCanVisitPublicApp(namespace, allServices)
	n.ui.DisplayCheckResult(result)
	if result.HasFailed {
		return nil
	}

	n.ui.DisplayCheckCompletion("The Ingresses are running correctly", success)
	return nil
}
