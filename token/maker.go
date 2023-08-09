package token

import "time"

type Maker interface {
	CreateToken(userId string, duration time.Duration) (string, error)
	VeryfiToken(token string) (*Payload, error)
}
