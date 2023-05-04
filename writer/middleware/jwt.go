package middleware

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/shorti.fy/writer/config"
	"net/http"
	"strings"
)

type TokenPayload struct {
	jwt.StandardClaims
	Iss   string   `json:"iss"`
	Sub   string   `json:"sub"`
	Aud   []string `json:"aud"`
	Iat   int      `json:"iat"`
	Exp   int      `json:"exp"`
	Azp   string   `json:"azp"`
	Scope string   `json:"scope"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Alg string   `json:"alg"`
	Kty string   `json:"kty"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	Kid string   `json:"kid"`
	X5T string   `json:"x5t"`
	X5C []string `json:"x5c"`
}

func ValidateOAuthToken(ctx iris.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.StopWithStatus(iris.StatusUnauthorized)
		return
	}

	rawToken := strings.TrimPrefix(authHeader, "Bearer ")

	payload := TokenPayload{}

	jwtToken, err := jwt.ParseWithClaims(rawToken, &payload, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		cert, err := getPemCertForKeyId(token.Header["kid"].(string))
		if err != nil {
			ctx.StopWithPlainError(iris.StatusUnauthorized, err)
		}
		return cert, nil
	})
	if err != nil {
		ctx.StopWithPlainError(iris.StatusUnauthorized, err)
	}

	err = jwtToken.Claims.Valid()
	if err != nil {
		ctx.StopWithPlainError(iris.StatusUnauthorized, err)
	}

	err = payload.Validate()
	if err != nil {
		ctx.StopWithPlainError(iris.StatusUnauthorized, err)
	}

	ctx.Next()
}

func getPemCertForKeyId(keyId string) (*rsa.PublicKey, error) {
	cert := ""
	resp, err := http.Get(config.EnvVariables.IDPDomain + ".well-known/jwks.json")
	if err != nil {
		return nil, err
	}

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return nil, err
	}

	for k := range jwks.Keys {
		if keyId == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5C[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return nil, errors.New("unable to find appropriate key")
	}

	c, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (t TokenPayload) Validate() error {
	if !checkAudExist(config.EnvVariables.IDPAudience, t.Aud) {
		return errors.New("invalid audience")
	}
	if !strings.Contains(t.Scope, config.EnvVariables.IDPScope) {
		return errors.New("invalid scope")
	}
	if t.Iss != config.EnvVariables.IDPDomain {
		return errors.New("invalid issuer")
	}
	return nil
}

func checkAudExist(aud string, audList []string) bool {
	if len(audList) == 0 {
		return false
	}

	for _, eachAud := range audList {
		if eachAud == aud {
			return true
		}
	}

	return false
}
