package helper

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(plainText string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
}
func ComparePassword(FromDb, FromBody string) bool {
	fmt.Println("---------")

	fmt.Println(FromDb)
	fmt.Println("-----------")
	fmt.Println(FromBody)

	err := bcrypt.CompareHashAndPassword([]byte(FromDb), []byte(FromBody))
	if err != nil {
		return false
	} else {
		return true
	}
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
