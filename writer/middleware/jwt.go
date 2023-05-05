package middleware

import (
	"context"
	"errors"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/shorti.fy/writer/config"
	"log"
	"net/http"
	"net/url"
	"time"
)

type CustomClaims struct {
	Aud         []string `json:"aud"`
	Permissions []string `json:"permissions"`
	Iss         string   `json:"iss"`
}

var customClaims = func() validator.CustomClaims {
	return &CustomClaims{}
}

func CheckJWT() iris.Handler {
	issuerURL, err := url.Parse(config.EnvVariables.IDPDomain)
	if err != nil {
		log.Fatalf("failed to parse the issuer url: %v", err)
	}
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	// Set up the validator.
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		config.EnvVariables.IDPDomain,
		[]string{config.EnvVariables.IDPAudience},
		validator.WithCustomClaims(customClaims),
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(ctx iris.Context) {
		encounteredError := true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			ctx.ResetRequest(r)
			ctx.Next()
		}

		middleware.CheckJWT(handler).ServeHTTP(ctx.ResponseWriter(), ctx.Request())

		if encounteredError {
			ctx.StopWithJSON(
				iris.StatusUnauthorized,
				map[string]string{"message": "JWT is invalid."},
			)
		}
	}
}

func (c *CustomClaims) Validate(ctx context.Context) error {
	if !checkItemExists(config.EnvVariables.IDPAudience, c.Aud) {
		return errors.New("invalid audience")
	}
	if !checkItemExists(config.EnvVariables.IDPScope, c.Permissions) {
		return errors.New("invalid permission")
	}
	if c.Iss != config.EnvVariables.IDPDomain {
		return errors.New("invalid issuer")
	}
	return nil
}

func checkItemExists(item string, arr []string) bool {
	if len(arr) == 0 {
		return false
	}

	for _, eachAud := range arr {
		if eachAud == item {
			return true
		}
	}

	return false
}
