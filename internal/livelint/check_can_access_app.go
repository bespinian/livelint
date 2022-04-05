package livelint

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

const connectionTimeoutSeconds = 2

func (n *Livelint) checkCanAccessApp(pods []apiv1.Pod) CheckResult {
	failureDetails := []string{}
	for _, pod := range pods {
		for _, container := range pod.Spec.Containers {
			for _, port := range container.Ports {
				if !n.portForwardAndCheck(pod, port.ContainerPort) {
					failureDetail := fmt.Sprintf("Pod %s, container %s has refused connection on port %d", pod.Name, container.Name, port.ContainerPort)
					failureDetails = append(failureDetails, failureDetail)
				}
			}
		}
	}

	checkResult := CheckResult{
		Message: "You can access the app",
	}
	if len(failureDetails) > 0 {
		checkResult = CheckResult{
			Message:   "One or more ports were not acessible",
			HasFailed: true,
			Details:   failureDetails,
		}
	}

	return checkResult
}

func (n *Livelint) portForwardAndCheck(pod apiv1.Pod, port int32) bool {
	connectionSuccessful := true

	// set up error handling used by port forwarding
	handleConnectionError := func(err error) {
		connectionSuccessful = false
	}
	runtime.ErrorHandlers = []func(error){handleConnectionError}

	// prepare the port forwarding
	roundTripper, upgrader, err := spdy.RoundTripperFor(n.config)
	if err != nil {
		panic(err)
	}
	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", pod.Namespace, pod.Name)
	hostIP := strings.TrimLeft(n.config.Host, "htps:/")
	serverURL := url.URL{Scheme: "https", Path: path, Host: hostIP}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: roundTripper}, http.MethodPost, &serverURL)
	stopChan, readyChan := make(chan struct{}, 1), make(chan struct{}, 1)
	out, errOut := new(bytes.Buffer), new(bytes.Buffer)
	ports := []string{fmt.Sprintf(":%d", port)}
	forwarder, err := portforward.New(dialer, ports, stopChan, readyChan, out, errOut)

	// start the port forwarding
	go func() {
		if err = forwarder.ForwardPorts(); err != nil {
			connectionSuccessful = false
		}
	}()

	// wait for port forwarding to be ready
	for range readyChan {
	}

	// send some traffic via the port forwarding
	forwardedPorts, _ := forwarder.GetPorts()
	_, err = net.Dial("tcp", fmt.Sprintf("localhost:%d", forwardedPorts[0].Local))
	if err != nil {
		return false
	}

	// wait for a certain time to see whether port forwarding error handling is called
	time.Sleep(connectionTimeoutSeconds * time.Second)

	return connectionSuccessful
}
