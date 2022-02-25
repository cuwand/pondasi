package logger

// Collections is log's collection of function
type Logger interface {
	InfoInterface(data interface{})
	Info(message string)
	Error(message string)
	Debug(message string)
}
