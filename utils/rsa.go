package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var (
	PrivateKey    *jwt.SigningMethodRSA
	RsaPrivateKey interface{}
	RsaPublicKey  interface{}
)

func LoadRSAKeys() error {
	privateKeyData, err := os.ReadFile("key/private_key.pem")
	if err != nil {
		return err
	}

	publicKeyData, err := os.ReadFile("key/public_key.pem")
	if err != nil {
		return err
	}

	RsaPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return err
	}

	RsaPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return err
	}

	return nil
}
