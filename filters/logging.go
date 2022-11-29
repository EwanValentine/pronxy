package filters

import (
	"net/http"

	"go.uber.org/zap"
)

// LoggerMiddleware is a filter, that logs out information about the request
// Use as a pre-filter, to see what requests are coming in etc...
type LoggerMiddleware struct {
	logger *zap.Logger
}

// NewLoggerMiddleware -
func NewLoggerMiddleware(logger *zap.Logger) http.Handler {
	return &LoggerMiddleware{logger}
}

// ServeHTTP -
func (mw *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mw.logger.Info(
		"request received",
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
	)
}
