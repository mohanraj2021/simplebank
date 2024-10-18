package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secreteKeyLen = 32

type JWTMaker struct {
	secretkey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < secreteKeyLen {
		return nil, fmt.Errorf("len of error string less than %d", secreteKeyLen)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(maker.secretkey))
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {

	keyfunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalid
		}

		return []byte(maker.secretkey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyfunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)

		if ok && errors.Is(verr.Inner, ErrTokenExpired) {
			return nil, ErrTokenExpired
		}

		return nil, ErrTokenInvalid
	}

	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, ErrTokenInvalid
	}

	return payload, nil
}
