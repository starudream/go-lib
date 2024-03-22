package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Interface interface {
	JTI() string
	ISS() string
	SUB() string
	AUD() string
	IAT() time.Time
	MTD() map[string]any

	Sign() (string, error)

	WithContext(ctx context.Context) context.Context

	Privilege() string

	jwt.Claims
}
