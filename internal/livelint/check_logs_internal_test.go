package livelint

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/matryer/is"
	apiv1 "k8s.io/api/core/v1"
	typedapiv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"
	fake "k8s.io/client-go/rest/fake"
)

type LogResponse struct {
	returnErr  bool
	statusCode int
	logLines   []string
}

func TestCheckContainerLogs(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it                    string
		pod                   apiv1.Pod
		logResponse           LogResponse
		previousLogResponse   LogResponse
		expectRequestPrevious bool
		expectedToFail        bool
		expectedMessage       string
	}{
		{
			it:  "succeeds if it can pull logs for a running container",
			pod: apiv1.Pod{},
			logResponse: LogResponse{
				returnErr:  false,
				statusCode: 200,
				logLines:   []string{"hello", "world"},
			},
			expectRequestPrevious: false,
			expectedToFail:        false,
			expectedMessage:       "You can see the logs for the app",
		},
		{
			it:  "succeeds if it can pull logs for a previously running container",
			pod: apiv1.Pod{},
			logResponse: LogResponse{
				returnErr:  true,
				statusCode: 200,
				logLines:   []string{},
			},
			previousLogResponse: LogResponse{
				returnErr:  false,
				statusCode: 200,
				logLines:   []string{"hello", "world"},
			},
			expectRequestPrevious: true,
			expectedToFail:        false,
			expectedMessage:       "You can see the logs for the app",
		},
		{
			it:  "fails if the logs are empty",
			pod: apiv1.Pod{},
			logResponse: LogResponse{
				returnErr:  false,
				statusCode: 200,
				logLines:   []string{""},
			},
			expectRequestPrevious: false,
			expectedToFail:        true,
			expectedMessage:       "You cannot see the logs for the app",
		},
		{
			it:  "fails if the logs cannot be retrieved from a running or previously running container",
			pod: apiv1.Pod{},
			logResponse: LogResponse{
				returnErr:  true,
				statusCode: 200,
				logLines:   []string{""},
			},
			previousLogResponse: LogResponse{
				returnErr:  true,
				statusCode: 200,
				logLines:   []string{""},
			},
			expectRequestPrevious: true,
			expectedToFail:        true,
			expectedMessage:       "You cannot see the logs for the app",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			requestedPrevious := false

			k8s := &kubernetesInterfaceMock{
				CoreV1Func: func() typedapiv1.CoreV1Interface {
					return &apiv1InterfaceMock{
						PodsFunc: func(namespace string) typedapiv1.PodInterface {
							return &apiv1PodInterfaceMock{
								GetLogsFunc: func(name string, opts *apiv1.PodLogOptions) *restclient.Request {
									header := http.Header{}
									header.Set("Content-Type", "text/plain")
									logResponse := tc.logResponse
									if opts.Previous {
										requestedPrevious = true
										logResponse = tc.previousLogResponse
									}

									client := fake.CreateHTTPClient(func(r *http.Request) (*http.Response, error) {
										if logResponse.returnErr {
											return nil, http.ErrServerClosed
										}
										return &http.Response{
											StatusCode: logResponse.statusCode,
											Header:     header,
											Body:       ioutil.NopCloser(bytes.NewReader([]byte(strings.Join(logResponse.logLines, "\n")))),
										}, nil
									})
									req := restclient.RESTClient{
										Client: client,
									}
									return req.Get()
								},
							}
						},
					}
				},
			}
			ll := Livelint{
				k8s: k8s,
			}
			result := ll.checkContainerLogs(tc.pod, "CONTAINER_NAME")

			is.Equal(requestedPrevious, tc.expectRequestPrevious)
			is.Equal(result.HasFailed, tc.expectedToFail) // HasFailed
			if !result.HasFailed /* && !tc.expectedToFail && requestedPrevious == tc.expectRequestPrevious */ {
				var logResponse LogResponse
				if tc.expectRequestPrevious {
					logResponse = tc.previousLogResponse
				} else {
					logResponse = tc.logResponse
				}
				is.Equal(len(result.Details), len(logResponse.logLines))
				for i, line := range logResponse.logLines {
					is.Equal(result.Details[i], line)
				}
			}
			is.Equal(result.Message, tc.expectedMessage) // Message
		})
	}
}
