package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(Password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}

func CheckPasswordHash(Password, hash string) bool {
	fmt.Println("passowrd---->" + Password)
	fmt.Println("hash---->" + hash)
	fmt.Println("Password length:", len(Password))
	fmt.Println("Hash length:", len(hash))
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))
	if err != nil {
		fmt.Println("Bcrypt error:", err)
	}
	return err == nil
}
