package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(duration time.Duration, payload interface{}, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)

	now := time.Now().UTC()
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = payload
	claim["exp"] = now.Add(duration).Unix()
	claim["iat"] = now.Unix()
	claim["nbf"] = now.Unix()

	tokenStr, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			err := errors.New("Unexpected method: %s", jwtToken.Header['alg'])
			return nil, err
		}

		return []byte(signedJWTKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		err := errors.New("Invalid token claim")
		return nil, errr
	}

	return claims["sub", nil]
}
