package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/cuwand/pondasi/helper/contextHelper"
	"github.com/cuwand/pondasi/logger"
	"github.com/gin-gonic/gin"
)

func RequestCORSAllowable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "os-name, app-version, device-id, app-id, os-version, "+
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Inject trace ID ke context bawaan Go
		c.Request = contextHelper.SetTraceIdRequestGin(c)

		data := map[string]interface{}{
			"method": c.Request.Method,
			"path":   c.FullPath(),
			"ip":     c.ClientIP(),
		}

		// Deteksi content-type
		ct := c.GetHeader("Content-Type")

		switch {
		case strings.Contains(ct, "application/json"):
			var body map[string]interface{}
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			_ = json.Unmarshal(bodyBytes, &body)

			// Reset body supaya handler masih bisa baca
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			data["body"] = body

		case c.Request.Method == http.MethodGet:
			query := map[string]string{}
			for k, v := range c.Request.URL.Query() {
				if len(v) > 0 {
					query[k] = v[0]
				}
			}
			data["query"] = query

		default:
			// Hanya tampilkan tipe request tanpa dump data
			data["body_type"] = ct
		}

		logger.GetAppLogger().InfoHttpWithContext(
			c.Request.Context(),
			"[REQUEST]",
			c.FullPath(),
			data,
		)

		c.Next()
	}
}
