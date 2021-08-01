package util

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

const SecretKey = "alisafdarirepo"

func GenerateJwt(Issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    Issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	return claims.SignedString([]byte(SecretKey))

}
