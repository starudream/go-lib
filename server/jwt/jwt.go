package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/starudream/go-lib/core/v2/utils/optionutil"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/server/v2/ierr"
)

type Claims struct {
	jwt.MapClaims `json:"-"`

	Id       string `json:"jti,omitempty"`
	Issuer   string `json:"iss,omitempty"`
	Subject  string `json:"sub,omitempty"`
	Audience string `json:"aud,omitempty"`
	IssuedAt int    `json:"iat,omitempty"`

	Metadata map[string]string `json:"mtd,omitempty"`

	raw string
}

var _ jwt.Claims = (*Claims)(nil)

func Sign(issuer, subject, audience string, options ...Option) (string, error) {
	claims := optionutil.Build(&Claims{
		Id:       uuid.NewString(),
		Issuer:   issuer,
		Subject:  subject,
		Audience: audience,
		IssuedAt: int(time.Now().Unix()),
		Metadata: map[string]string{},
	}, options...)

	var (
		method jwt.SigningMethod = jwt.SigningMethodHS256
		key    any               = secretKey
	)
	if privateKey != nil && publicKey != nil {
		method, key = jwt.SigningMethodRS256, privateKey
	}

	token := jwt.NewWithClaims(method, claims)

	raw, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	claims.raw = raw

	return claims.raw, nil
}

func (c *Claims) Validate() error {
	return nil
}

func Parse(raw string) (*Claims, error) {
	claims := &Claims{raw: raw}

	_, err := jwt.ParseWithClaims(raw, claims, func(token *jwt.Token) (any, error) {
		if privateKey != nil && publicKey != nil {
			return publicKey, nil
		}
		return secretKey, nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, err
	}

	err = claims.Validate()
	if err != nil {
		return nil, err
	}

	return claims, nil
}

const jwtCtxkey = "jwt-ctxkey"

func (c *Claims) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, jwtCtxkey, c)
}

func FromContext(ctx context.Context) (*Claims, error) {
	if c, ok := ctx.Value(jwtCtxkey).(*Claims); ok {
		return c, nil
	}
	return nil, ierr.Unauthorized(9999, "does not contain jwt claims")
}

func MustFromContext(ctx context.Context) *Claims {
	return osutil.Must1(FromContext(ctx))
}
