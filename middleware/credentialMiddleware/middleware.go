package credentialMiddleware

import (
	"fmt"
	"github.com/cariapo/cservice/credential"
	"github.com/cariapo/cservice/credential/model/request"
	"github.com/cuwand/pondasi/constant"
	"github.com/cuwand/pondasi/errors"
	"github.com/cuwand/pondasi/helper/contextHelper"
	"github.com/cuwand/pondasi/helper/headerHelper"
	"github.com/cuwand/pondasi/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type CredentialMiddleware interface {
	Auth(c *gin.Context)
	Authorize(c *gin.Context)
}

type implementMiddleware struct {
	credentialOutbound credential.CredentialOutbound
}

func ImplementMiddleware(credential credential.CredentialOutbound) CredentialMiddleware {
	return implementMiddleware{
		credentialOutbound: credential,
	}
}

var (
	accessToken  string
	refreshToken string
)

func (i implementMiddleware) Authorize(c *gin.Context) {

	scope := strings.Split(c.FullPath(), "/")[2]

	authorized, err := i.credentialOutbound.Authorize(headerHelper.GetGinAndValidate(c, "Authorization"))

	if err != nil {
		response.Error(c, err)
		return
	}

	for x := range authorized.Data.Authorities {
		for y := range authorized.Data.Authorities[x].Scopes {
			if authorized.Data.Authorities[x].Scopes[y].Scope == scope {
				c.Request = contextHelper.SetValueGin(c, constant.X_CLIENT_ID, authorized.Data.Id)
				c.Next()
				return
			}
		}
	}

	response.ErrorWithMessage(c, http.StatusUnauthorized,
		fmt.Sprintf("Aplikasi tidak memiliki akses %s", strings.ToUpper(scope)),
		fmt.Sprintf("Application not have access %s", strings.ToUpper(scope)), errors.DefaultErrorCode)
	return
}

func (i implementMiddleware) Auth(c *gin.Context) {

	if len(accessToken) == 0 {
		err := i.doClientCredentials(c)

		if err != nil {
			response.Error(c, err)
			return
		}
	}

	err := i.doAuthorize(c)

	if err != nil {
		response.Error(c, err)
		return
	}

	c.Next()
}

func (i implementMiddleware) doAuthorize(c *gin.Context) error {
	authorized, err := i.credentialOutbound.Authorize(accessToken)

	if err != nil {
		return i.doRefreshToken(c)
	}

	c.Request = contextHelper.SetValueGin(c, constant.X_CLIENT_ID, authorized.Data.Id)

	return nil
}

func (i implementMiddleware) doClientCredentials(c *gin.Context) error {
	username, password, ok := c.Request.BasicAuth()

	if !ok {
		return errors.UnauthorizedError("Unauthorized")
	}

	oauth, err := i.credentialOutbound.Oauth(request.StoreToken{
		GrantType: "client_credentials",
		Username:  username,
		Password:  password,
	})

	if err != nil {
		return errors.UnauthorizedError("Unauthorized, Auth Failed")
	}

	accessToken = oauth.GetAccessToken()
	refreshToken = oauth.GetRefreshToken()

	return nil
}

func (i implementMiddleware) doRefreshToken(c *gin.Context) error {
	username, password, ok := c.Request.BasicAuth()

	if !ok {
		return errors.UnauthorizedError("Unauthorized")
	}

	oauth, err := i.credentialOutbound.Oauth(request.StoreToken{
		GrantType:    "refresh_token",
		Username:     username,
		Password:     password,
		RefreshToken: refreshToken,
	})

	if err != nil {
		return i.doClientCredentials(c)
	}

	accessToken = oauth.GetAccessToken()
	refreshToken = oauth.GetRefreshToken()

	return nil
}
