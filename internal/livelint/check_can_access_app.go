package livelint

import (
	"bytes"
	"fmt"
	"log"
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
	reasonsToFail := []string{}
	for _, pod := range pods {
		for _, container := range pod.Spec.Containers {
			for _, port := range container.Ports {
				if !n.canPortForward(pod, port.ContainerPort, checkTCPConnection) {
					reason := fmt.Sprintf("container %s of Pod %s has refused connection on port %v", container.Name, pod.Name, port.ContainerPort)
					reasonsToFail = append(reasonsToFail, reason)
				}
			}
		}
	}

	if len(reasonsToFail) > 0 {
		return CheckResult{
			HasFailed: true,
			Message:   "One or more ports were not accessible",
			Details:   reasonsToFail,
		}
	}

	return CheckResult{
		Message: "You can access the app",
	}
}

func (n *Livelint) canPortForward(pod apiv1.Pod, port int32, check func(uint16) bool) bool {
	log.Print("Checking port forwarding")
	connectionSuccessful := true

	// set up error handling used by port forwarding
	handleConnectionError := func(err error) {
		log.Printf("Connection error %s", err)
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
	if err != nil {
		log.Printf("Error crating port forwarder %s", err)
	}

	// start the port forwarding
	go func() {
		if err = forwarder.ForwardPorts(); err != nil {
			log.Printf("Error forwarding port %s", err)
			connectionSuccessful = false
		}
	}()

	// wait for port forwarding to be ready
	for range readyChan {
	}

	n.tea.Send(showSpinnerMsg{showing: true})

	// perform the check function against the forwarded port
	forwardedPorts, _ := forwarder.GetPorts()
	if !check(forwardedPorts[0].Local) {
		return false
	}

	// wait for a certain time to see whether port forwarding error handling is called
	time.Sleep(connectionTimeoutSeconds * time.Second)
	n.tea.Send(showSpinnerMsg{showing: false})

	return connectionSuccessful
}

func checkTCPConnection(port uint16) bool {
	_, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	log.Print(fmt.Errorf("Error sending tcp packet on port %d %w", port, err))
	return err != nil
}
