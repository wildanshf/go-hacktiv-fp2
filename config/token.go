package config

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type MyClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

var mySigningKey = []byte("MySecrets")

func CreateToken(id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaim{
		ID:             id,
		StandardClaims: jwt.StandardClaims{},
	})

	signedStr, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}

	return signedStr
}

func VerifyToken(token string) bool {
	t, err := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return mySigningKey, nil
	})
	if err != nil {
		panic(err)
	}

	if _, ok := t.Claims.(*MyClaim); ok && t.Valid {
		return true
	}

	return false
}

func GetClaim(token string) *MyClaim {
	t, err := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return mySigningKey, nil
	})
	if err != nil {
		panic(err)
	}

	if v, ok := t.Claims.(*MyClaim); ok && t.Valid {
		return v
	}

	return nil
}
