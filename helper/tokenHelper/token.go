package tokenHelper

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/MicahParks/keyfunc"
	"github.com/cuwand/pondasi/enum/grandTypeEnums"
	"github.com/cuwand/pondasi/errors"
	"github.com/cuwand/pondasi/helper/dateHelper"
	jwt "github.com/golang-jwt/jwt/v4"
	"math/big"
	"os"
	"time"
)

const (
	DefaultPrivateFileRSA = "./configs/resources/private_key.pem"
	DefaultPublicFileRSA  = "./configs/resources/public_key.pub"
	DefaultKeyIdentifier  = "236c88d15b514cd3b73bd6b4c1fbe177"
)

type RefreshTokenRequest struct {
	Id          string
	TokenId     string
	Audience    []string
	CurrentTime time.Time
	ExpiredAt   time.Time
}

type AccessTokenRequest struct {
	GrandType   grandTypeEnums.GrandType
	Id          string
	TokenId     string
	Audience    []string
	CurrentTime time.Time
	ExpiredAt   time.Time
	ClientId    string
	Authorities []string
	Username    string
	Roles       []string
}

type accessToken struct {
	Id          string                   `json:"id"`
	Username    string                   `json:"username,omitempty"`
	Roles       []string                 `json:"roles,omitempty"`
	ClientId    string                   `json:"client_id,omitempty"`
	Authorities []string                 `json:"authorities,omitempty"`
	GrandType   grandTypeEnums.GrandType `json:"grand_type"`
	jwt.RegisteredClaims
}

type refreshToken struct {
	Id        string                   `json:"id"`
	GrandType grandTypeEnums.GrandType `json:"grand_type"`
	jwt.RegisteredClaims
}

type token struct {
	Id        string                   `json:"id"`
	GrandType grandTypeEnums.GrandType `json:"grand_type"`
	jwt.RegisteredClaims
}

type PublicKeysData struct {
	Keys []keyData `json:"keys"`
}

func (p PublicKeysData) ToRaw() []byte {
	jwkMarhaled, err := json.Marshal(p)

	if err != nil {
		panic(err)
	}

	return jwkMarhaled
}

type keyData struct {
	Alg string `json:"alg,omitempty"`
	Kty string `json:"kty,omitempty"`
	Kid string `json:"kid,omitempty"`
	Use string `json:"use,omitempty"`
	N   string `json:"n,omitempty"`
	E   string `json:"e,omitempty"`
}

type Config struct {
	privateKey    *rsa.PrivateKey
	publicKey     *rsa.PublicKey
	keyIdentifier string
}

func InitConfig(privateKeyLoc, publicKeyLoc, keyIdentifier string) Config {
	privateKeyReader, err := os.ReadFile(privateKeyLoc)

	if err != nil {
		panic("private key rsa not found")
	}

	publicKeyReader, err := os.ReadFile(publicKeyLoc)

	if err != nil {
		panic("public key rsa not found")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyReader)

	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyReader)

	if err != nil {
		panic(err)
	}

	return Config{
		privateKey:    privateKey,
		publicKey:     publicKey,
		keyIdentifier: keyIdentifier,
	}
}

func (c Config) GenerateAccessToken(req AccessTokenRequest) (string, error) {
	claimStrings := []string{}

	for x := range req.Audience {
		claimStrings = append(claimStrings, req.Audience[x])
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessToken{
		Id:          req.Id,
		Username:    req.Username,
		Roles:       req.Roles,
		ClientId:    req.ClientId,
		Authorities: req.Authorities,
		GrandType:   req.GrandType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "CREDENTIAL",
			Audience: claimStrings,
			ExpiresAt: &jwt.NumericDate{
				Time: req.ExpiredAt,
			},
			IssuedAt: &jwt.NumericDate{
				Time: req.CurrentTime,
			},
			ID: req.TokenId,
		},
	})

	newToken.Header["kid"] = c.keyIdentifier

	return newToken.SignedString(c.privateKey)
}

func (c Config) GenerateRefreshToken(req RefreshTokenRequest) (string, error) {
	claimStrings := []string{}

	for x := range req.Audience {
		claimStrings = append(claimStrings, req.Audience[x])
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshToken{
		Id:        req.Id,
		GrandType: grandTypeEnums.REFRESH_TOKEN,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "CREDENTIAL",
			Audience: claimStrings,
			ExpiresAt: &jwt.NumericDate{
				Time: req.ExpiredAt,
			},
			IssuedAt: &jwt.NumericDate{
				Time: req.CurrentTime,
			},
			ID: req.TokenId,
		},
	})

	newToken.Header["kid"] = c.keyIdentifier

	return newToken.SignedString(c.privateKey)
}

func (c Config) ParseToken(tokenString string) (*token, error) {
	tokenParsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("RS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return c.publicKey, nil
	})

	if tokenParsed == nil || err != nil {
		return nil, errors.UnauthorizedError("Invalid Access Token")
	}

	claim := tokenParsed.Claims.(jwt.MapClaims)

	if claim.VerifyExpiresAt(dateHelper.TimeNow().Unix(), false) == false {
		return nil, errors.UnauthorizedError("Token Expired")
	}

	// Handle Payload Token
	if claim["id"] == nil {
		return nil, errors.UnauthorizedError("Invalid Access Token, payload issue")
	}

	if claim["grand_type"] == nil {
		return nil, errors.UnauthorizedError("Invalid Access Token, payload issue")
	}

	return &token{
		Id:        claim["id"].(string),
		GrandType: grandTypeEnums.FromString(claim["grand_type"].(string)),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: claim["iss"].(string),
			ID:     claim["jti"].(string),
		},
	}, nil
}

func (c Config) ParseRefreshToken(tokenString string) (*refreshToken, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("RS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return c.publicKey, nil
	})

	if token == nil {
		if err != nil {
			return nil, errors.UnauthorizedError(fmt.Sprintf("Invalid Access Token %s", err.Error()))
		}

		return nil, errors.UnauthorizedError("Invalid Access Token")
	}

	claim := token.Claims.(jwt.MapClaims)

	if claim.VerifyExpiresAt(dateHelper.TimeNow().Unix(), false) == false {
		return nil, errors.UnauthorizedError("Token Expired")
	}

	// Handle Payload Token
	if claim["id"] == nil {
		return nil, errors.UnauthorizedError("Invalid Access Token, payload issue")
	}

	if claim["grand_type"] == nil {
		return nil, errors.UnauthorizedError("Invalid Access Token, payload issue")
	}

	return &refreshToken{
		Id:        claim["id"].(string),
		GrandType: grandTypeEnums.FromString(claim["grand_type"].(string)),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: claim["iss"].(string),
			ID:     claim["jti"].(string),
		},
	}, nil
}

func (c Config) ParseAccessToken(tokenString string) (*accessToken, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("RS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return c.publicKey, nil
	})

	if token == nil {
		if err != nil {
			return nil, errors.UnauthorizedError(fmt.Sprintf("Invalid Access Token %s", err.Error()))
		}

		return nil, errors.UnauthorizedError("Invalid Access Token")
	}

	return c.toAccessToken(*token)
}

func (c Config) toAccessToken(jwtToken jwt.Token) (*accessToken, error) {
	claim := jwtToken.Claims.(jwt.MapClaims)

	if claim.VerifyExpiresAt(dateHelper.TimeNow().Unix(), false) == false {
		return nil, errors.UnauthorizedError("Token Expired")
	}

	// Handle Payload Token
	if claim["id"] == nil {
		return nil, errors.UnauthorizedError("Invalid Access Token, payload issue")
	}

	if claim["grand_type"] == nil {
		return nil, errors.UnauthorizedError("Invalid Access Token, payload issue")
	}

	var username string
	var clientId string
	roles := []string{}
	authorities := []string{}
	grandType := grandTypeEnums.FromString(claim["grand_type"].(string))

	if grandType == grandTypeEnums.PASSWORD {
		for _, val := range claim["roles"].([]interface{}) {
			roles = append(roles, val.(string))
		}

		username = claim["username"].(string)
	}

	if grandType == grandTypeEnums.CLIENT_CREDENTIAL {
		for _, val := range claim["authorities"].([]interface{}) {
			authorities = append(authorities, val.(string))
		}

		clientId = claim["client_id"].(string)
	}

	return &accessToken{
		Id:          claim["id"].(string),
		Username:    username,
		Roles:       roles,
		ClientId:    clientId,
		Authorities: authorities,
		GrandType:   grandType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: claim["iss"].(string),
			ID:     claim["jti"].(string),
		},
	}, nil
}

func (c Config) ParseAccessTokenByJWK(tokenString string, publicKeysData PublicKeysData) (*accessToken, error) {
	jwks, err := keyfunc.NewJSON(publicKeysData.ToRaw())

	if err != nil {
		return nil, errors.InternalServerError(fmt.Sprintf("Failed to create JWKS from JSON.\nError: %s", err.Error()))
	}

	// Parse the JWT.
	token, err := jwt.Parse(tokenString, jwks.Keyfunc)

	if token == nil {
		if err != nil {
			return nil, errors.UnauthorizedError(fmt.Sprintf("Invalid Access Token %s", err.Error()))
		}

		return nil, errors.UnauthorizedError("Invalid Access Token")
	}

	return c.toAccessToken(*token)
}

func (c Config) GenerateJWKS() PublicKeysData {
	return PublicKeysData{
		Keys: []keyData{
			{
				Alg: "RS256",
				Kty: "RSA",
				Use: "sig",
				Kid: c.keyIdentifier,
				N:   base64.RawURLEncoding.EncodeToString((*c.publicKey.N).Bytes()),
				E:   base64.StdEncoding.EncodeToString(big.NewInt(int64(c.publicKey.E)).Bytes()),
			},
		},
	}
}
