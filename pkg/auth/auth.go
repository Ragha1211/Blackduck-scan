package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

var customClaims = func() validator.CustomClaims {
	return &CustomClaimsExample{}
}

type authConfig struct {
	AUTH_ISSUER_URL    string
	SECRET_SIGNING_KEY string
	AUTH_AUD           string
}

func loadAuthConfig() *authConfig {
	return &authConfig{
		AUTH_ISSUER_URL:    os.Getenv("AUTH_ISSUER_URL"),
		SECRET_SIGNING_KEY: os.Getenv("SECRET_SIGNING_KEY"),
		AUTH_AUD:           os.Getenv("AUTH_AUD"),
	}
}

type CustomClaimsExample struct {
	Issuer       string `json:"iss"`
	Sub          string `json:"sub"`
	Azp          string `json:"azp"`
	ShouldReject bool   `json:"shouldreject,omitempty"`
}

// Validate errors out if `ShouldReject` is true.
func (c *CustomClaimsExample) Validate(ctx context.Context) error {
	if c.ShouldReject {
		return errors.New("should reject was set to true")
	}
	return nil
}

// validateJWTToken is a gin.HandlerFunc middleware
// that will check the validity of our JWT.
func ValidateJWTToken() gin.HandlerFunc {

	jwtConfig := loadAuthConfig()

	// Set up the validator.
	jwtValidator, err := validator.New(
		func(ctx context.Context) (interface{}, error) {
			return []byte(jwtConfig.SECRET_SIGNING_KEY), nil
		},
		validator.HS256,
		jwtConfig.AUTH_ISSUER_URL,
		[]string{jwtConfig.AUTH_AUD},
		validator.WithCustomClaims(customClaims),
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		log.Fatalf("failed to set up the validator:  %v", err)
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT:  %v", err)
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(ctx *gin.Context) {
		encounteredError := true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			ctx.Request = r
			ctx.Next()
		}

		middleware.CheckJWT(handler).ServeHTTP(ctx.Writer, ctx.Request)

		if encounteredError {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				map[string]string{"message": "JWT is invalid."},
			)
		}
	}
}
