package livelint

import (
	"context"
	"fmt"
	"log"
	"net/http"

	apiv1 "k8s.io/api/core/v1"
)

func (n *Livelint) checkCanVisitPublicApp(namespace string, services []apiv1.Service) CheckResult {
	ingresses, err := n.getIngressesFromServices(namespace, services)
	if err != nil || len(ingresses) < 1 {
		return CheckResult{
			HasFailed:    true,
			Message:      "You cannot visit the app from the public internet",
			Instructions: "No matching ingresses were found",
		}
	}
	result := CheckResult{
		Message: "You can visit the app from the public internet",
	}
	for _, ingress := range ingresses {
		urls := getUrlsFromIngress(ingress)
		for _, url := range urls {
			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url.String(), nil)
			if err != nil {
				log.Fatal(fmt.Errorf("error when creating http request for url %s: %w", url.String(), err))
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return CheckResult{
					HasFailed: true,
					Message:   fmt.Sprintf("You cannot visit %s from public internet: %s", url.String(), err.Error()),
				}
			}
			defer resp.Body.Close()
			if resp.StatusCode < 200 || resp.StatusCode >= 500 {
				return CheckResult{
					HasFailed:    true,
					Message:      "You cannot visit the app from the public internet",
					Instructions: fmt.Sprintf("HTTP request to URL %s returned with response code %v", url.String(), resp.StatusCode),
				}
			}
			result.Details = append(result.Details, fmt.Sprintf("Response code was %v", resp.StatusCode))
		}
	}
	return result
}
