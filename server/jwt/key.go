package jwt

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

var (
	secretKey = []byte("GI7WgEK8zlasWo9TFoXNddgpAqikexkV")

	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	if s := config.Get("jwt.secret_key").String(); s != "" {
		secretKey = []byte(s)
	}

	var err error

	if pri, pub := config.Get("jwt.private_key").String(), config.Get("jwt.public_key").String(); pri != "" && pub != "" {
		privateKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(pri))
		if err != nil {
			osutil.PanicErr(err)
		}
		publicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(pub))
		if err != nil {
			osutil.PanicErr(err)
		}
	}
}
