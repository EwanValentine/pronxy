package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"go.uber.org/zap"
)

// Proxy -
type Proxy struct {
	proxy       *httputil.ReverseProxy
	target      *url.URL
	logger      *zap.Logger
	preFilters  []http.Handler
	postFilters []http.Handler
}

// NewReverseProxy -
func NewReverseProxy(
	target *url.URL,
	logger *zap.Logger,
	opts ...ProxyOption,
) *Proxy {
	proxy := &Proxy{
		proxy:       httputil.NewSingleHostReverseProxy(target),
		target:      target,
		logger:      logger,
		preFilters:  []http.Handler{},
		postFilters: []http.Handler{},
	}

	// Loop through each option
	for _, opt := range opts {
		opt(proxy)
	}

	return proxy
}

// ServeHTTP -
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Run list of middlewares
	for _, filter := range p.preFilters {
		filter.ServeHTTP(w, r)
	}

	r.URL.Host = p.target.Host

	p.logger.Info("target", zap.String("target", r.URL.String()))

	p.proxy.ServeHTTP(w, r)

	for _, filter := range p.postFilters {
		filter.ServeHTTP(w, r)
	}
}
