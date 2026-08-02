package impl

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type OtelMiddlewareFactory struct{}

func NewOtelMiddlewareFactory() *OtelMiddlewareFactory {
	return &OtelMiddlewareFactory{}
}

func (o *OtelMiddlewareFactory) Gin(serviceName string) gin.HandlerFunc {
	return otelgin.Middleware(serviceName)
}
