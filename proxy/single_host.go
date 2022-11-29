package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/pkg/errors"
)

// NewSingleHostProxy -
func NewSingleHostProxy(target string) (*httputil.ReverseProxy, error) {
	u, err := url.Parse(target)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing target URL")
	}

	return httputil.NewSingleHostReverseProxy(u), nil
}

// ProxyRequestHandler -
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		proxy.ServeHTTP(w, req)
	}
}
