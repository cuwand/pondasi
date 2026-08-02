package logger

import (
	"context"

	"github.com/cuwand/pondasi/constant"
)

func GetAppLogger() Logger {
	return appLoggerConfig
}

func (log LogConfig) Info(message string) {
	log.logger.Info().Msg(message)
}

func (log LogConfig) InfoHttp(message, path string, data interface{}) {
	log.logger.Info().
		Str("path", path).
		Interface("response", data).
		Msg(message)
}

func GetTraceID(ctx context.Context) string {
	if v := ctx.Value(constant.X_CORE_TRACE_ID); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return "TRACE-ID-NOT-FOUND" // default kalau tidak ada
}

func (log LogConfig) InfoHttpWithContext(ctx context.Context, message, path string, data interface{}) {
	traceId := GetTraceID(ctx)

	log.logger.Info().
		Str("path", path).
		Str("trace_id", traceId).
		Interface("response", data).
		Msg(message)
}

func (log LogConfig) InfoInterface(data interface{}) {
	log.logger.Info().Interface("data", data)
}

func (log LogConfig) Error(message string) {
	log.logger.Error().Msg(message)
}

func (log LogConfig) ErrorWithContext(ctx context.Context, message string) {
	traceId := GetTraceID(ctx)

	log.logger.Error().
		Str("trace_id", traceId).
		Msg(message)
}

func (log LogConfig) Debug(message string) {
	log.logger.Debug().Msg(message)
}
