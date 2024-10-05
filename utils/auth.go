package utils

import (
	"errors"
	"time"

	"deepak.com/web_rest/models"
	"github.com/golang-jwt/jwt/v5"
)

const SECRETKEY = "SUPERSECRET"

func GenerateJwtToken(u models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"id":    u.Id,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(SECRETKEY))
}

func Verify(token string) (int64, error) {
	signedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return "", errors.New("unauthorized token")
		}

		return []byte(SECRETKEY), nil
	})

	if err != nil {
		return 0, errors.New("invalid token")
	}

	claims, ok := signedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid claims used")
	}

	//email := claims["email"]
	id := int64(claims["id"].(float64))

	return id, nil

}
