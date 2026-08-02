package logger

import "context"

// Collections is log's collection of function
type Logger interface {
	InfoInterface(data interface{})
	Info(message string)
	Error(message string)
	ErrorWithContext(ctx context.Context, message string)
	Debug(message string)
	InfoHttp(message, path string, data interface{})
	InfoHttpWithContext(ctx context.Context, message, path string, data interface{})
}
