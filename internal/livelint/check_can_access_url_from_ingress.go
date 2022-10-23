package livelint

import (
	"context"
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
	result := CheckResult{HasFailed: false}
	checkHTTPConnection := func(port uint16) (bool, string) {
		localhostURL := getLocalhostURL(url, port)
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, localhostURL, nil)
		if err != nil {
			log.Fatal(fmt.Errorf("error creating HTTP GET request for URL %s: %w", localhostURL, err))
		}
		if url.Hostname() != "" && url.Hostname() != "localhost" {
			req.Host = url.Hostname()
		}
		resp, err := http.DefaultClient.Do(req)
		statusString := fmt.Sprintf("HTTP status code %d", resp.StatusCode)
		if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
			resp.Body.Close()
			return false, statusString
		}
		resp.Body.Close()
		return true, statusString
	}

	portForwardOk, connectionCheckMsg := n.canPortForward(ingressControllerPod, port, checkHTTPConnection)
	if !portForwardOk {
		result = CheckResult{
			HasFailed: true,
			Message:   fmt.Sprintf("App cannot be reached from ingress controller pod %s: %s", ingressControllerPod.Name, connectionCheckMsg),
		}
	}
	return result
}

func (n *Livelint) getForwardedPorts(ingress netv1.Ingress, ingressClasses map[string]netv1.IngressClass) (ForwardedPorts, error) {
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
		controllerServicesFullMatch, controllerServicesPartialMatch, err := n.getServicesForPod("", controllerPods[0])
		if err != nil {
			return result, fmt.Errorf("error getting ingress controller service for ingress %s: %w", ingress.Name, err)
		}

		// find the port which the services forward to for http
		for _, controllerService := range controllerServicesPartialMatch {
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

func (n *Livelint) checkCanAccessAppFromIngressController(ingress netv1.Ingress, ingressClasses map[string]netv1.IngressClass) CheckResult {
	result := CheckResult{HasFailed: false, Message: "You can visit the app from the corresponding ingress controller"}

	// find the ingress controller's pods
	controllerPods, err := n.getIngressControllerPods(ingress, ingressClasses)
	if err != nil {
		return CheckResult{HasFailed: true, Message: fmt.Sprintf("error getting controller pods for ingress %s", ingress.Name)}
	}

	forwardedPorts, err := n.getForwardedPorts(ingress, ingressClasses)
	if err != nil {
		return CheckResult{HasFailed: true, Message: fmt.Sprintf("error getting target ports of ingress controller for ingress %s", ingress.Name)}
	}
	urls := getUrlsFromIngress(ingress)

	for _, url := range urls {
		urlResult := n.checkCanAccessURLFromIngressControllerPods(url, forwardedPorts, controllerPods)
		if urlResult.HasFailed {
			return CheckResult{HasFailed: true, Message: urlResult.Message}
		}

	}
	return result
}

func getLocalhostURL(url url.URL, port uint16) string {
	return fmt.Sprintf("%s://localhost:%d%s", url.Scheme, port, url.Path)
}
