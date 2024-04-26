package router

import (
	"fmt"
	"github.com/cuwand/pondasi/enum/environtmentEnums"
	"github.com/cuwand/pondasi/errors"
	"github.com/cuwand/pondasi/helper/dateHelper"
	"github.com/cuwand/pondasi/response"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

var startTime time.Time

func GetRouter() *gin.Engine {
	startTime = dateHelper.TimeNow()

	if environtmentEnums.Environment(os.Getenv("ENVIRONMENT")) == environtmentEnums.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		response.Error(c, errors.NotFound("Page Not Found"))
	})

	router.NoMethod(func(c *gin.Context) {
		response.Error(c, errors.MethodNotAllowed("Method Not Allowed"))
	})

	router.Use(CatchError())

	return router
}

func RegisterDefaultHome(router *gin.Engine) {
	appVersion, ok := os.LookupEnv("APP_VERSION")

	if !ok {
		appVersion = "0.0.0-PONDASI"
	}

	appName, ok := os.LookupEnv("APP_NAME")

	if !ok {
		appName = fmt.Sprintf("PONDASI-%v", time.Now().Format("02012006"))
	}

	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(200, fmt.Sprintf("%s | %s | %s", appName, appVersion, dateHelper.TimeToString(startTime)))
	})
}

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Error(c, errors.InternalServerError(fmt.Sprintf("Something when wrong - [%v]", err)))
			}
		}()
		c.Next()
	}
}
