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
	n.tea.Send(contextMsg(fmt.Sprintf("Checking deployment %s in namespace %s", deploymentName, namespace)))

	allPods, err := n.getDeploymentPods(namespace, deploymentName)
	if err != nil {
		return fmt.Errorf("error getting Pods: %w", err)
	}

	// Are you hitting the ResourceQuota limits ?
	result := n.checkAreResourceQuotasHit(namespace, deploymentName)
	n.tea.Send(stepMsg(result))
	if result.HasFailed {
		return nil
	}

	// Is there any PENDING Pod?
	result = checkAreTherePendingPods(allPods)
	n.tea.Send(stepMsg(result))
	if result.HasFailed {

		// Is the cluster full?
		result = n.checkIsClusterFull(allPods)
		n.tea.Send(stepMsg(result))
		if result.HasFailed {
			return nil
		}

		// Are you mounting a PENDING PersistentVolumeClaim?
		result = n.checkIsMountingPendingPVC(allPods, namespace)
		n.tea.Send(stepMsg(result))
		if result.HasFailed {
			return nil
		}

		// Is the Pod assigned to the Node?
		result = n.checkIsPodAssignedToNode()
		n.tea.Send(stepMsg(result))

		return nil
	}

	// Are any Pods restart cycling
	result = n.checkAreThereRestartCyclingPods(allPods)
	n.tea.Send(stepMsg(result))
	if result.HasFailed {
		return nil
	}

	// Are the Pods RUNNING?
	result = checkAreAllPodsRunning(allPods)
	n.tea.Send(stepMsg(result))
	if result.HasFailed {
		for _, pod := range allPods {
			nonRunningContainers := getNonRunningContainers(pod)
			if len(nonRunningContainers) > 0 {

				// Is the Pod status ImagePullBackOff?
				result = checkImagePullErrors(pod, nonRunningContainers[0].Name)
				n.tea.Send(stepMsg(result))
				if result.HasFailed {

					// Is the name of the image correct?
					result = n.checkIsImageNameCorrect()
					n.tea.Send(stepMsg(result))
					if result.HasFailed {
						return nil
					}

					// Is the image tag valid? Does it exist?
					result = n.checkIsImageTagValid()
					n.tea.Send(stepMsg(result))
					if result.HasFailed {
						return nil
					}

					// Are you pulling images from a private registry?
					result = n.checkIsPullingFromPrivateRegistry()
					n.tea.Send(stepMsg(result))

					return nil
				}

				// Is the Pod status CrashLoopBackOff?
				result = checkCrashLoopBackOff(pod, nonRunningContainers[0].Name)
				n.tea.Send(stepMsg(result))
				if result.HasFailed {

					// Did you inspect the logs and fix the crashing app?
					result = n.checkDidInspectLogsAndFix()
					n.tea.Send(stepMsg(result))
					if result.HasFailed {
						// Can you see the logs for the app?
						result = n.checkContainerLogs(pod, nonRunningContainers[0].Name)
						n.tea.Send(stepMsg(result))
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
				n.tea.Send(stepMsg(result))
				if result.HasFailed {
					// Is there any container running?
					result = checkAreThereRunningContainers(pod)
					n.tea.Send(stepMsg(result))
					return nil
				}

				result = n.checkFailedMount(pod)
				n.tea.Send(stepMsg(result))

				// Is there any container running?
				result = checkAreThereRunningContainers(pod)
				n.tea.Send(stepMsg(result))

				return nil
			}
		}
	}

	// Are the Pods READY?
	result = checkAreAllPodsReady(allPods)
	n.tea.Send(stepMsg(result))
	if result.HasFailed {

		// Is the Readiness probe failing?
		result = n.checkReadinessProbe(allPods)
		n.tea.Send(stepMsg(result))

		return nil
	}

	// Can you access the app?
	result = n.checkCanAccessApp(allPods)
	n.tea.Send(stepMsg(result))
	if result.HasFailed {

		// Is the port exposed by container correct and listing on 0.0.0.0?
		result = n.checkIsPortExposedCorrectly()
		n.tea.Send(stepMsg(result))

		return nil
	}

	n.tea.Send(summaryMsg("Pods are running correctly"))

	allServices, partlyMatchingServices, err := n.getServices(namespace, deploymentName)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting services in namespace %s: %w", namespace, err))
	}

	allServices = append(allServices, partlyMatchingServices...)

	if len(allServices) < 1 {
		result = CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf("No services match the deployment's matchSelector."),
		}
		n.tea.Send(stepMsg(result))
		return nil
	}

	for _, service := range allServices {

		// Can you see a list of endpoints?
		result = n.checkCanSeeEndpoints(service.Name, namespace)
		n.tea.Send(stepMsg(result))
		if result.HasFailed {
			// Does the Pod have an IP address assigned?
			result = checkPodHasIPAddressAssigned(allPods)
			n.tea.Send(stepMsg(result))

			return nil
		}

		// Can you visit the app?
		result = n.checkCanVisitServiceApp(service)
		n.tea.Send(stepMsg(result))
		if result.HasFailed {

			// Is the targetPort on the Service matching the containerPort in the Pod?
			result = n.checkTargetPortMatchesContainerPort(allPods, service.Name, namespace)
			n.tea.Send(stepMsg(result))

			return nil
		}

		n.tea.Send(summaryMsg(fmt.Sprintf("The Service %s is running correctly", service.Name)))
	}

	ingresses, err := n.getIngressesFromServices(namespace, allServices)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting ingresses in namespace %s: %w", namespace, err))
	}

	if len(ingresses) < 1 {
		result = CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf("No ingresses match any of the previously detected services names and ports."),
		}
		n.tea.Send(stepMsg(result))
		return nil
	}

	for _, ingress := range ingresses {
		// Can you see a list of Backends?
		result = n.checkCanSeeBackends(ingress, namespace)
		n.tea.Send(stepMsg(result))
		if result.HasFailed {
			return nil
		}

		n.tea.Send((summaryMsg(fmt.Sprintf("The Ingress %s is set up correctly", ingress.Name))))
	}

	// Can you visit the app?
	result = n.checkCanVisitIngressApp()
	n.tea.Send(stepMsg(result))
	if result.HasFailed {
		return nil
	}

	n.tea.Send(summaryMsg("The Ingresses are running correctly"))

	// The app should be running. Can you visit it from the public internet?
	result = n.checkCanVisitPublicApp()
	n.tea.Send(stepMsg(result))
	if result.HasFailed {
		return nil
	}
	n.tea.Send(summaryMsg("All checks finished"))
	return nil
}
