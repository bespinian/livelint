package livelint

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	apiv1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
)

const (
	httpPortNumber  = 80
	httpsPortNumber = 443
)

type ForwardedPorts struct {
	HTTP       int32
	HTTPFound  bool
	HTTPS      int32
	HTTPSFound bool
}

func (n *Livelint) checkCanAccessAppFromIngressControllerPod(ingressControllerPod apiv1.Pod, port int32, url url.URL) CheckResult {
	result := CheckResult{HasFailed: true}

	checkHTTPConnection := func(port uint16) bool {
		localhostURL := getLocalhostURL(url, port)
		req, err := http.NewRequest("GET", localhostURL, nil)
		if err != nil {
			log.Fatal(fmt.Errorf("error creating http GET request for url %s: %w", localhostURL, err))
		}
		if url.Hostname() != "" && url.Hostname() != "localhost" {
			req.Header.Add("Host", url.Hostname())
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil || (resp.StatusCode >= 200 && resp.StatusCode < 300) {
			return false
		}
		return true
	}

	if !n.canPortForward(ingressControllerPod, port, checkHTTPConnection) {
		result = CheckResult{HasFailed: true}
	}
	return result
}

func (n *Livelint) getForwardedPorts(namespace string, ingress netv1.Ingress, ingressClasses map[string]netv1.IngressClass) (ForwardedPorts, error) {
	var result ForwardedPorts
	// find the ingress controller's pods
	controllerPods, err := n.getIngressControllerPods(ingress, ingressClasses)
	if err != nil {
		return result, fmt.Errorf("error getting controller pods for ingress %s: %w", ingress.Name, err)
	}
	var httpsPort int32
	var httpsFound bool
	var httpPort int32
	var httpFound bool
	if len(controllerPods) > 0 {
		// find the infress controllers services
		controllerServicesFullMatch, _, err := n.getServicesForPod(namespace, controllerPods[0])
		if err != nil {
			return result, fmt.Errorf("error getting ingress controller service for ingress %s: %w", ingress.Name, err)
		}

		// find the port which the services forward to for http
		for _, controllerService := range controllerServicesFullMatch {
			httpPort, httpFound = getTargetPortNumber(controllerService, controllerPods[0], httpPortNumber)
			if httpFound {
				break
			}
		}

		// find the port which the services forwards to for https
		for _, controllerService := range controllerServicesFullMatch {
			httpsPort, httpsFound = getTargetPortNumber(controllerService, controllerPods[0], httpsPortNumber)
			if httpsFound {
				break
			}
		}
	}
	return ForwardedPorts{HTTP: httpPort, HTTPFound: httpFound, HTTPS: httpsPort, HTTPSFound: httpsFound}, nil
}

func (n *Livelint) checkCanAccessURLFromIngressControllerPods(url url.URL, forwardedPorts ForwardedPorts, controllerPods []apiv1.Pod) CheckResult {
	result := CheckResult{HasFailed: false}
	var port int32
	failed := false
	if url.Scheme == "https" {
		port = forwardedPorts.HTTPS
	} else {
		port = forwardedPorts.HTTP
	}
	var podResult CheckResult
	for _, controllerPod := range controllerPods {
		podResult = n.checkCanAccessAppFromIngressControllerPod(controllerPod, port, url)
		if podResult.HasFailed {
			failed = true
			break
		}
	}
	if failed {
		result = CheckResult{HasFailed: true}
	}
	return result
}

func (n *Livelint) checkCanAccessAppFromIngressController(namespace string, ingress netv1.Ingress, ingressClasses map[string]netv1.IngressClass) (CheckResult, error) {
	result := CheckResult{HasFailed: false}

	// find the ingress controller's pods
	controllerPods, err := n.getIngressControllerPods(ingress, ingressClasses)
	if err != nil {
		return result, fmt.Errorf("error getting controller pods for ingress %s: %w", ingress.Name, err)
	}

	forwardedPorts, err := n.getForwardedPorts(namespace, ingress, ingressClasses)
	if err != nil {
		return result, fmt.Errorf("error getting target ports of ingress controller for ingress %s: %w", ingress.Name, err)
	}
	urls := getUrlsFromIngress(ingress)

	failed := false
	for _, url := range urls {
		urlResult := n.checkCanAccessURLFromIngressControllerPods(url, forwardedPorts, controllerPods)
		if urlResult.HasFailed {
			failed = true
			break
		}

	}
	if failed {
		result = CheckResult{HasFailed: true}
	}
	return result, nil
}

func getLocalhostURL(url url.URL, port uint16) string {
	return fmt.Sprintf("%s://localhost:%d/%s", url.Scheme, port, url.Path)
}
