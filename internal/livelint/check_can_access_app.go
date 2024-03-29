package livelint

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ishidawataru/sctp"
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
				checkFunc := checkTCPConnection
				if port.Protocol == apiv1.ProtocolUDP {
					checkFunc = checkUDPConnection
				}
				if port.Protocol == apiv1.ProtocolSCTP {
					checkFunc = checkSCTPConnection
				}
				portForwardOk, connectionCheckMsg := n.canPortForward(pod, port.ContainerPort, checkFunc)
				if !portForwardOk {
					reason := fmt.Sprintf("container %s of Pod %s has refused %s connection on port %v: %s", container.Name, pod.Name, port.Protocol, port.ContainerPort, connectionCheckMsg)
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

func (n *Livelint) canPortForward(pod apiv1.Pod, port int32, checkFunc func(uint16) (bool, string)) (bool, string) {
	connectionSuccessful := true

	// set up error handling used by port forwarding
	handleConnectionError := func(err error) {
		connectionSuccessful = false
	}
	runtime.ErrorHandlers = append(runtime.ErrorHandlers, handleConnectionError)

	// prepare the port forwarding
	roundTripper, upgrader, err := spdy.RoundTripperFor(n.config)
	if err != nil {
		panic(err)
	}
	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", pod.Namespace, pod.Name)
	hostIP := strings.TrimLeft(n.config.Host, "htps:/")
	fmt.Println(hostIP)
	serverURL := url.URL{Scheme: "https", Path: path, Host: hostIP}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: roundTripper}, http.MethodPost, &serverURL)
	stopChan, readyChan := make(chan struct{}, 1), make(chan struct{}, 1)
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
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

	n.ui.StartSpinner()

	// send some traffic via the port forwarding
	forwardedPorts, _ := forwarder.GetPorts()
	checkOk, message := checkFunc(forwardedPorts[0].Local)
	if !checkOk {
		return false, message
	}

	// wait for a certain time to see whether port forwarding error handling is called
	time.Sleep(connectionTimeoutSeconds * time.Second)
	n.ui.StopSpinner()

	return connectionSuccessful, message
}

func checkTCPConnection(port uint16) (bool, string) {
	_, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return false, fmt.Sprintf("TCP connection error: %s", err)
	}
	return true, "TCP connection successful"
}

func checkUDPConnection(port uint16) (bool, string) {
	_, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return false, fmt.Sprintf("UDP connection error: %s", err)
	}
	return true, "UDP connection successful"
}

func checkSCTPConnection(port uint16) (bool, string) {
	ipAddr, err := net.ResolveIPAddr("ip", "localhost")
	if err != nil {
		return false, fmt.Sprintf("error resolving localhost: %s", err)
	}
	addr := sctp.SCTPAddr{
		IPAddrs: []net.IPAddr{*ipAddr},
		Port:    int(port),
	}
	_, err = sctp.DialSCTP("sctp", nil, &addr)
	if err != nil {
		return false, fmt.Sprintf("SCTP connection error: %s", err)
	}
	return true, "SCTP connection successful"
}
