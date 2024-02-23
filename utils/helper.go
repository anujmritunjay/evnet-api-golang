package utils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func GenerateToken(userEmail string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  userEmail,
	})

	tokenString, err := token.SignedString([]byte("MYSECRETKEY"))
	return tokenString, err

}

func VerifyToken(token string) error {
	data, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
		return errors.New("Unauthorized")
	}
	if data.Valid {
		claims, err := data.Claims.(jwt.MapClaims)
		if err {
			return errors.New("Unauthorized")
		}
		email := claims["email"].(string)
		userId := claims["userId"].(int64)
		fmt.Println(email, userId)
	}
	return nil
}
