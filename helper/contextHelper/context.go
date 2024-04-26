package contextHelper

import (
	"context"
	"encoding/json"
	"github.com/cuwand/pondasi/constant"
	"github.com/cuwand/pondasi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetValueGin(c *gin.Context, key, value string) *http.Request {
	contextParent := c.Request.Context()

	c.Request = c.Request.WithContext(context.WithValue(contextParent, key, value))

	return c.Request
}

func SetUserRequestGin(c *gin.Context, request models.UserRequest) *http.Request {
	contextParent := c.Request.Context()

	requestMarshaled, _ := json.Marshal(request)

	c.Request = c.Request.WithContext(context.WithValue(contextParent, constant.X_CORE_UR, requestMarshaled))

	return c.Request
}

func GetUserRequestGin(c *gin.Context) models.UserRequest {
	contextParent := c.Request.Context()

	userRequest := &models.UserRequest{}

	value := contextParent.Value(constant.X_CORE_UR).([]byte)

	err := json.Unmarshal(value, &userRequest)

	if err != nil {
		panic("user request not valid")
	}

	if userRequest == nil {
		panic("user request value not found")
	}

	return *userRequest
}

func GetValueStringAndValidateGin(c *gin.Context, key string) string {
	contextParent := c.Request.Context()

	value := contextParent.Value(key).(string)

	if len(value) == 0 {
		panic("value not found")
	}

	return value
}

func SetValue(contextParent context.Context, key, value string) context.Context {
	return context.WithValue(contextParent, key, value)
}

func GetValueStringAndValidate(context context.Context, key string) string {
	value := context.Value(key).(string)

	if len(value) == 0 {
		panic("value not found")
	}

	return value
}
