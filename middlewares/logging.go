package middlewares

import (
	"net/http"

	"go.uber.org/zap"
)

type LoggerMiddleware struct {
	logger *zap.Logger
}

func NewLoggerMiddleware(logger *zap.Logger) http.Handler {
	return &LoggerMiddleware{logger}
}

func (mw *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mw.logger.Info(
		"request received",
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
	)
}
