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
						return nil
					}

					// Did you forget the CMD instruction in the Dockerfile?
					result = n.checkForgottenCMDInDockerfile()
					if result.HasFailed {
						return nil
					}

					return nil
				}

				// Is the pod status RunContainerError?
				result = checkRunContainerError(pod)
				n.tea.Send(stepMsg(result))
				if result.HasFailed {

					// Is there any container running?
					result = n.checkAreThereRunningContainers()
					n.tea.Send(stepMsg(result))

					return nil
				}

				// Can you see the logs for the app?
				result = n.checkContainerLogs(pod, nonRunningContainers[0].Name)
				n.tea.Send(stepMsg(result))
				if !result.HasFailed {
					return nil
				}

				result = n.checkFailedMount(pod)
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

	service := n.askUserForService(namespace, deploymentName)
	// Can you see a list of endpoints?
	result = n.checkCanSeeEndpoints(service.Name, namespace)
	n.tea.Send(stepMsg(result))
	if result.HasFailed {

		// Is the Selector matching the Pod label?
		result = n.checkIsSelectorMatchingPodLabel(allPods, service.Name, namespace)
		n.tea.Send(stepMsg(result))
		if result.HasFailed {
			return nil
		}

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

	n.tea.Send(summaryMsg("The Service is running correctly"))

	ingressName := n.askUserQuestion("Which ingress should expose this deployment?")

	// Can you see a list of Backends?
	result = n.checkCanSeeBackends(ingressName, namespace)
	n.tea.Send(stepMsg(result))
	if result.HasFailed {

		// Are the serviceName and servicePort matching the Service?
		result = n.checkServiceNameAndPortMatchService()
		n.tea.Send(stepMsg(result))

		return nil
	}

	// Can you visit the app?
	result = n.checkCanVisitIngressApp()
	n.tea.Send(stepMsg(result))
	if result.HasFailed {
		return nil
	}

	n.tea.Send(summaryMsg("The Ingress is running correctly"))

	// The app should be running. Can you visit it from the public internet?
	result = n.checkCanVisitPublicApp()
	n.tea.Send(stepMsg(result))
	if result.HasFailed {
		return nil
	}

	n.tea.Send(summaryMsg("All checks finished"))
	return nil
}
