package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint) (string, error) {

	// Token generating logic
	fmt.Println("Generating  token ")
	calims := jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, calims)
	return token.SignedString(RsaPrivateKey)

}
