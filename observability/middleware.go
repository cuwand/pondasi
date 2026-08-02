package observability

import "github.com/gin-gonic/gin"

// MiddlewareFactory adalah kontrak pembungkus middleware
type MiddlewareFactory interface {
	Gin(serviceName string) gin.HandlerFunc
}

// default = noop
var middlewareFactory MiddlewareFactory

// SetMiddlewareFactory inject implementation (otel / vendor lain)
func SetMiddlewareFactory(f MiddlewareFactory) {
	if f != nil {
		middlewareFactory = f
	}
}

// Middleware dipakai di main.go
func Middleware(serviceName string) gin.HandlerFunc {
	return middlewareFactory.Gin(serviceName)
}
