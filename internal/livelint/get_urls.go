package livelint

import (
	"net/url"
	"regexp"

	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/net"
)

const wildcardDomain = `(?m)\*\.(.+)`

func getUrlsFromIngress(ingress netv1.Ingress) []url.URL {
	wildcardDomainRegex := regexp.MustCompile(wildcardDomain)
	urls := []url.URL{}
	for _, ingressRule := range ingress.Spec.Rules {
		var host string
		switch {
		case ingressRule.Host == "":
			host = "localhost"
		case wildcardDomainRegex.MatchString(ingressRule.Host):
			host = wildcardDomainRegex.FindStringSubmatch(ingressRule.Host)[1]
		default:
			host = ingressRule.Host
		}
		var scheme string
		var port int
		if hasTLSCertificate(ingressRule.Host, ingress) {
			scheme = "https"
			port = 443
		} else {
			scheme = "http"
			port = 80
		}
		for _, path := range ingressRule.HTTP.Paths {
			urls = append(urls, *net.FormatURL(scheme, host, port, path.Path))
		}
	}
	return urls
}

func hasTLSCertificate(host string, ingress netv1.Ingress) bool {
	for _, tlsConfig := range ingress.Spec.TLS {
		for _, tlsHost := range tlsConfig.Hosts {
			if host == tlsHost {
				return true
			}
		}
	}
	return false
}
