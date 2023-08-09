package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(duration time.Duration, payload interface{}, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = payload
	claim["exp"] = now.Add(duration).Unix()
	claim["iat"] = now.Unix()
	claim["nbf"] = now.Unix()

	tokenStr, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", fmt.Errorf("Generating jwt failed: %v", err)
	}

	return tokenStr, nil
}

func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("Invalid token claim")
	}

	return claims["sub"], nil
}
