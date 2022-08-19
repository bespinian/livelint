package livelint

import (
	"fmt"
	"log"
)

func (n *Livelint) Start() {
	err := n.tea.Start()
	if err != nil {
		log.Fatal(fmt.Errorf("error starting ui: %w", err))
	}
}

// RunChecks checks for potential issues with the deployment.
func (n *Livelint) RunChecks(namespace, deploymentName string, isVerbose bool) error {
	// nolint:govet
	statusMsg := initalizeStatus(fmt.Sprintf("Checking deployment %s in namespace %s", deploymentName, namespace))
	n.tea.Send(statusMsg)

	statusMsg.StartCheck("Checking that pods are running correctly")
	n.tea.Send(statusMsg)

	allPods, err := n.getDeploymentPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error getting Pods: %w", err)
	}

	// Are you hitting the ResourceQuota limits ?
	result := n.checkAreResourceQuotasHit(namespace, deploymentName)

	statusMsg.AddCheckResult(result)
	n.tea.Send(statusMsg)

	if result.HasFailed {
		return nil
	}

	// Is there any PENDING Pod?
	result = checkAreTherePendingPods(allPods)

	statusMsg.AddCheckResult(result)
	n.tea.Send(statusMsg)

	if result.HasFailed {

		// Is the cluster full?
		result = n.checkIsClusterFull(allPods)
		statusMsg.AddCheckResult(result)
		n.tea.Send(statusMsg)
		if result.HasFailed {
			return nil
		}

		// Are you mounting a PENDING PersistentVolumeClaim?
		result = n.checkIsMountingPendingPVC(allPods, namespace)
		statusMsg.AddCheckResult(result)
		n.tea.Send(statusMsg)
		if result.HasFailed {
			return nil
		}

		// Is the Pod assigned to the Node?
		result = checkIsPodAssignedToNode(allPods)
		statusMsg.AddCheckResult(result)
		n.tea.Send(statusMsg)

		return nil
	}

	// Are any Pods restart cycling
	result = n.checkAreThereRestartCyclingPods(allPods)
	statusMsg.AddCheckResult(result)
	n.tea.Send(statusMsg)
	if result.HasFailed {
		return nil
	}

	// Are the Pods RUNNING?
	result = checkAreAllPodsRunning(allPods)
	statusMsg.AddCheckResult(result)
	n.tea.Send(statusMsg)
	if result.HasFailed {
		for _, pod := range allPods {
			nonRunningContainers := getNonRunningContainers(pod)
			if len(nonRunningContainers) > 0 {

				// Is the Pod status ImagePullBackOff?
				result = checkImagePullErrors(pod, nonRunningContainers[0].Name)
				statusMsg.AddCheckResult(result)
				n.tea.Send(statusMsg)
				if result.HasFailed {

					// Is the name of the image correct?
					result = n.checkIsImageNameCorrect()
					statusMsg.AddCheckResult(result)
					n.tea.Send(statusMsg)
					if result.HasFailed {
						return nil
					}

					// Is the image tag valid? Does it exist?
					result = n.checkIsImageTagValid()
					statusMsg.AddCheckResult(result)
					n.tea.Send(statusMsg)
					if result.HasFailed {
						return nil
					}

					// Are you pulling images from a private registry?
					result = n.checkIsPullingFromPrivateRegistry()
					statusMsg.AddCheckResult(result)
					n.tea.Send(statusMsg)

					return nil
				}

				// Is the Pod status CrashLoopBackOff?
				result = checkCrashLoopBackOff(pod, nonRunningContainers[0].Name)
				statusMsg.AddCheckResult(result)
				n.tea.Send(statusMsg)
				if result.HasFailed {

					// Did you inspect the logs and fix the crashing app?
					result = n.checkDidInspectLogsAndFix()
					statusMsg.AddCheckResult(result)
					n.tea.Send(statusMsg)
					if result.HasFailed {
						// Can you see the logs for the app?
						result = n.checkContainerLogs(pod, nonRunningContainers[0].Name)
						statusMsg.AddCheckResult(result)
						n.tea.Send(statusMsg)
						if !result.HasFailed {
							return nil
						}
					}

					// Did you forget the CMD instruction in the Dockerfile?
					result = n.checkForgottenCMDInDockerfile()
					if result.HasFailed {
						return nil
					}

					return nil
				}

				// Is the pod status RunContainerError?
				result = checkIsContainerCreating(pod)
				statusMsg.AddCheckResult(result)
				n.tea.Send(statusMsg)
				if result.HasFailed {
					// Is there any container running?
					result = checkAreThereRunningContainers(pod)
					statusMsg.AddCheckResult(result)
					n.tea.Send(statusMsg)
					return nil
				}

				result = n.checkFailedMount(pod)
				statusMsg.AddCheckResult(result)
				n.tea.Send(statusMsg)

				// Is there any container running?
				result = checkAreThereRunningContainers(pod)
				statusMsg.AddCheckResult(result)
				n.tea.Send(statusMsg)

				return nil
			}
		}
	}

	// Are the Pods READY?
	result = checkAreAllPodsReady(allPods)
	statusMsg.AddCheckResult(result)
	n.tea.Send(statusMsg)
	if result.HasFailed {

		// Is the Readiness probe failing?
		result = n.checkReadinessProbe(allPods)
		statusMsg.AddCheckResult(result)
		n.tea.Send(statusMsg)

		return nil
	}

	// Can you access the app?
	result = n.checkCanAccessApp(allPods)
	statusMsg.AddCheckResult(result)
	n.tea.Send(statusMsg)
	if result.HasFailed {

		// Is the port exposed by container correct and listing on 0.0.0.0?
		result = n.checkIsPortExposedCorrectly()
		statusMsg.AddCheckResult(result)
		n.tea.Send(statusMsg)

		return nil
	}

	statusMsg.CompleteCheck(summaryMsg{text: "Pods are running correctly", kind: success})
	n.tea.Send(statusMsg)

	statusMsg.StartCheck("Checking service")
	n.tea.Send(statusMsg)

	allServices, partlyMatchingServices, err := n.getServices(namespace, deploymentName)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting services in namespace %s: %w", namespace, err))
	}

	allServices = append(allServices, partlyMatchingServices...)

	if len(allServices) < 1 {
		result = CheckResult{
			HasFailed: true,
			Message:   "No services match the deployment's matchSelector.",
		}
		statusMsg.AddCheckResult(result)
		n.tea.Send(statusMsg)
		return nil
	}

	for _, service := range allServices {

		// Can you see a list of endpoints?
		result = n.checkCanSeeEndpoints(service.Name, namespace)
		statusMsg.AddCheckResult(result)
		n.tea.Send(statusMsg)
		if result.HasFailed {
			// Is the selector matching the right pod label?
			result = n.checkIsSelectorMatchPodLabel(namespace, service.Name, allPods)
			statusMsg.AddCheckResult(result)
			n.tea.Send(statusMsg)
			if result.HasFailed {
				// Does the Pod have an IP address assigned?
				result = checkPodHasIPAddressAssigned(allPods)
				statusMsg.AddCheckResult(result)
				n.tea.Send(statusMsg)
			}
			return nil
		}

		// Can you visit the app?
		result = n.checkCanVisitServiceApp(service)
		statusMsg.AddCheckResult(result)
		n.tea.Send(statusMsg)
		if result.HasFailed {

			// Is the targetPort on the Service matching the containerPort in the Pod?
			result = n.checkTargetPortMatchesContainerPort(allPods, service.Name, namespace)
			statusMsg.AddCheckResult(result)
			n.tea.Send(statusMsg)

			return nil
		}

		statusMsg.CompleteCheck(summaryMsg{text: fmt.Sprintf("The Service %s is running correctly", service.Name), kind: success})
		n.tea.Send(statusMsg)
	}

	statusMsg.StartCheck("Checking ingress")
	n.tea.Send(statusMsg)
	ingresses, err := n.getIngressesFromServices(namespace, allServices)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting ingresses in namespace %s: %w", namespace, err))
	}

	if len(ingresses) < 1 {
		result = CheckResult{
			HasFailed: true,
			Message:   "No ingresses match any of the previously detected services names and ports.",
		}
		statusMsg.AddCheckResult(result)
		n.tea.Send(statusMsg)
		return nil
	}

	for _, ingress := range ingresses {
		// Can you see a list of Backends?
		result = n.checkCanSeeBackends(ingress, namespace)
		statusMsg.AddCheckResult(result)
		n.tea.Send(statusMsg)
		if result.HasFailed {
			return nil
		}

		statusMsg.CompleteCheck(summaryMsg{text: fmt.Sprintf("The Ingress %s is set up correctly", ingress.Name), kind: success})
		n.tea.Send(statusMsg)
	}

	// Can you visit the app?
	statusMsg.StartCheck("Checking app connectivity")
	n.tea.Send(statusMsg)
	result = n.checkCanVisitIngressApp()
	statusMsg.AddCheckResult(result)
	n.tea.Send(statusMsg)
	if result.HasFailed {
		return nil
	}

	// The app should be running. Can you visit it from the public internet?
	result = n.checkCanVisitPublicApp()
	statusMsg.AddCheckResult(result)
	n.tea.Send(statusMsg)
	if result.HasFailed {
		return nil
	}
	statusMsg.CompleteCheck(summaryMsg{text: "The Ingresses are running correctly", kind: success})
	n.tea.Send(statusMsg)
	return nil
}
