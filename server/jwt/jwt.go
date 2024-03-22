package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/optionutil"
	"github.com/starudream/go-lib/server/v2/ierr"
)

type claims struct {
	jwt.MapClaims `json:"-"`

	Id       string `json:"jti,omitempty"`
	Issuer   string `json:"iss,omitempty"`
	Subject  string `json:"sub,omitempty"`
	Audience string `json:"aud,omitempty"`
	IssuedAt int    `json:"iat,omitempty"`

	Metadata map[string]any `json:"mtd,omitempty"`

	raw string
}

var _ Interface = (*claims)(nil)

func New(issuer, subject, audience string, options ...Option) Interface {
	return optionutil.Build(&claims{
		Id:       uuid.NewString(),
		Issuer:   issuer,
		Subject:  subject,
		Audience: audience,
		IssuedAt: int(time.Now().Unix()),
		Metadata: map[string]any{},
	}, options...)
}

func (c *claims) JTI() string {
	return c.Id
}

func (c *claims) ISS() string {
	return c.Issuer
}

func (c *claims) SUB() string {
	return c.Subject
}

func (c *claims) AUD() string {
	return c.Audience
}

func (c *claims) IAT() time.Time {
	return time.Unix(int64(c.IssuedAt), 0)
}

func (c *claims) MTD() map[string]any {
	return c.Metadata
}

func (c *claims) Sign() (string, error) {
	var (
		method jwt.SigningMethod = jwt.SigningMethodHS256
		key    any               = secretKey
	)
	if privateKey != nil && publicKey != nil {
		method, key = jwt.SigningMethodRS256, privateKey
	}

	token := jwt.NewWithClaims(method, c)

	raw, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	c.raw = raw

	return c.raw, nil
}

func Parse(raw string) (Interface, error) {
	c := &claims{raw: raw}

	_, err := jwt.ParseWithClaims(raw, c, func(token *jwt.Token) (any, error) {
		if privateKey != nil && publicKey != nil {
			return publicKey, nil
		}
		return secretKey, nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, ierr.Unauthorized(9999, "parse token error: %v", err)
	}

	err = c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *claims) Validate() error {
	if c.IssuedAt < 0 || c.Metadata["privilege"] == true {
		return nil
	}
	if time.Since(c.IAT()) > expireTime {
		return ierr.Unauthorized(9999, "token expired")
	}
	return nil
}

func (c *claims) Privilege() string {
	c.IssuedAt = -1
	c.Metadata = map[string]any{"privilege": true}
	if _, err := c.Sign(); err != nil {
		slog.Warn("sign token error: %v", err)
	}
	return c.raw
}
