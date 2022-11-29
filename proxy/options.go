package proxy

import "net/http"

type ProxyOption func(*Proxy)

func WithPreFilters(filters ...http.Handler) ProxyOption {
	return func(p *Proxy) {
		p.preFilters = filters
	}
}

func WithPostFilters(filters ...http.Handler) ProxyOption {
	return func(p *Proxy) {
		p.postFilters = filters
	}
}
