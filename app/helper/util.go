package helper

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(plainText []byte) ([]byte, error) {

	return bcrypt.GenerateFromPassword(plainText, bcrypt.DefaultCost)
}

func GenerateToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, tokenErr := token.SignedString([]byte("secret"))
	if tokenErr != nil {
		fmt.Println(tokenErr)
		return t
	} else {
		return t
	}
}
