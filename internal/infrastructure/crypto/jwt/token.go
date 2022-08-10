package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("SECRET_JWT"))

type Claims struct {
	Cpf string `json:"cpf"`
	jwt.StandardClaims
}

func GenerateToken(cpf string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Cpf: cpf,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func ValidadeToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}
	return nil
}

// err = jwt.ValidadeToken(token)
// if err != nil {
// 	return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
// }