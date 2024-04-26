package headerHelper

import (
	"fmt"
	"github.com/cuwand/pondasi/constant"
	"github.com/cuwand/pondasi/crypto/aes"
	"github.com/cuwand/pondasi/helper/envHelper"
	"github.com/cuwand/pondasi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGin(c *gin.Context, key string) string {
	return c.Request.Header.Get(key)
}

func GetGinAndValidate(c *gin.Context, key string) string {
	value := c.Request.Header.Get(key)

	if len(value) == 0 {
		panic(fmt.Sprintf("Header %s is required", key))
	}

	return value
}

func SetGin(c *gin.Context, key, value string) {
	c.Request.Header.Set(key, value)
}

func Get(header http.Header, key string) string {
	return header.Get(key)
}

func GetUserAudit(c *gin.Context) models.UserRequest {
	value := c.Request.Header.Get(constant.X_CORE_UA)
	auditKey := envHelper.GetEnvAndValidate("AUDIT_KEY")

	if len(value) == 0 {
		panic(fmt.Sprintf("Header %s is required", constant.X_CORE_UA))
	}

	usrReq := models.UserRequest{}

	aes.DecryptPayload(auditKey, auditKey, value, &usrReq)

	return usrReq
}

func GenerateUserAudit(user models.UserRequest) string {
	auditKey := envHelper.GetEnvAndValidate("AUDIT_KEY")
	
	return aes.EncryptPayload(auditKey, auditKey, user)
}
