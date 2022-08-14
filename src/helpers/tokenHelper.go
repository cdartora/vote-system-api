package helper

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Name     string
	Password string
	jwt.StandardClaims
}

var (
	Secret = "secret"
)

func CreateToken(name string, password string) (signedToken string, err error) {
	claims := &SignedDetails{
		Name:     name,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(), // 24 horas
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(Secret))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, err
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(Secret), nil
	})
}
