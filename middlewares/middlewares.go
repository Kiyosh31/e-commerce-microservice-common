package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Kiyosh31/e-commerce-microservice-common/env"
	"github.com/Kiyosh31/e-commerce-microservice-common/token"
	"github.com/Kiyosh31/e-commerce-microservice-common/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func errorResponse(err error) *gin.H {
	return &gin.H{"Errors": err.Error()}
}

func AuthHttpMiddleware(bearerToken string, tokenSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(env.AuthHeaderKey)
		fields := strings.Fields(authHeader)

		// validate token is provided
		if len(fields) == 0 {
			err := errors.New("Authorization is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// validate auth has Bearer token
		if len(fields) < 2 {
			err := errors.New("Invalid authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// validate auth is Bearer token type
		authType := strings.ToLower(fields[0])
		if authType != env.AuthTypeBearer {
			err := fmt.Errorf("Unsuported auth type: %s", authType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// extracting the token code
		// checking token is valid
		tokenCode := fields[1]
		tokenUserId, err := token.ValidateToken(tokenCode, tokenSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		c.Set(env.AuthPayloadKey, tokenUserId)
		c.Next()
	}
}

func AuthGrpcMiddleware(ctx context.Context, tokenSecret string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("Missing metadata")
	}

	values := md.Get(env.AuthHeaderKey)
	if len(values) == 0 {
		return "", fmt.Errorf("Missing authorization header")
	}

	bearer := values[0]
	fields := strings.Fields(bearer)

	if len(fields) < 2 {
		return "", fmt.Errorf("Invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != env.AuthTypeBearer {
		return "", fmt.Errorf("Unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]

	payloadToken, err := token.ValidateToken(accessToken, tokenSecret)

	if err != nil {
		return "", fmt.Errorf("Invalid access token: %s", err)
	}

	userId, err := utils.ParseInterfaceToString(payloadToken)
	if err != nil {
		return "", fmt.Errorf("Could not parse token: %s", err)
	}

	return userId, nil
}
