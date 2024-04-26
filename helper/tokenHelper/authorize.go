package tokenHelper

import (
	"github.com/cuwand/pondasi/database/redis"
	"github.com/cuwand/pondasi/helper/httpHelper"
	"github.com/cuwand/pondasi/logger"
	"net/http"
	"time"
)

type ConfigAuthorize struct {
	redisClient   redis.Redis
	cacheDuration time.Duration
	cacheKey      string
	jwkUrl        string
	httpConfig    httpHelper.HttpConfig
	logger        logger.Logger
}

func InitConfigAuthorize(jwkUrl, cacheKey string, redisClient redis.Redis, logger logger.Logger) ConfigAuthorize {
	httpConfig := httpHelper.HttpConfig{}

	return ConfigAuthorize{
		redisClient:   redisClient,
		cacheDuration: 10 * time.Minute,
		cacheKey:      cacheKey,
		jwkUrl:        jwkUrl,
		httpConfig:    httpConfig,
		logger:        logger,
	}
}

func (c ConfigAuthorize) FetchCacheableJWK() (*PublicKeysData, error) {
	var publicKeysData *PublicKeysData

	err := c.redisClient.Get(c.cacheKey, &publicKeysData)

	if err != nil {
		return nil, err
	}

	if publicKeysData == nil {
		publicKeysData, err = c.FetchJWK()

		if err != nil {
			return nil, err
		}

		err = c.redisClient.Set(c.cacheKey, *publicKeysData, c.cacheDuration)

		if err != nil {
			return nil, err
		}
	}

	return publicKeysData, nil
}

func (c ConfigAuthorize) FetchJWK() (*PublicKeysData, error) {
	response := &PublicKeysData{}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        c.jwkUrl,
		Method:     http.MethodGet,
		Result:     response,
		Client:     c.httpConfig.HttpClient,
		TimeoutReq: c.httpConfig.Timeout,
		Logger:     c.logger,
		Header: &http.Header{
			"Accept": []string{httpHelper.ContentTypeApplicationJson},
		},
	}); err != nil {
		return nil, err
	}

	return response, nil
}
