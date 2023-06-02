package encryption

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

const key = "13tgh4389u309g03h94r29hf"

func SignedLoginToken(citizenID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": citizenID})

	return token.SignedString([]byte(key))
}

func CheckSignedToken(token string) (bool, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, errors.New("unexpected signing method")
		}
		return []byte(key), nil

	})
	if err != nil {
		return false, err
	}
	return t.Valid, nil
}
