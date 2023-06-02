package xjwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type BaseToken struct {
	Config *Config
}

func (j *BaseToken) IDClaims(id int, expire time.Duration, issuer string) *jwt.RegisteredClaims {
	sub := strconv.Itoa(id)
	exp := jwt.NewNumericDate(time.Now().Add(expire))
	jti := RandString(8)
	jwtToken := &jwt.RegisteredClaims{
		Subject:   sub,
		ExpiresAt: exp,
		ID:        jti,
		Issuer:    issuer,
	}
	return jwtToken
}

func (j *BaseToken) Token(claims *jwt.RegisteredClaims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(j.Config.GetSecret())
}

func (j *BaseToken) Parse(tokenString string, issuer string) (*jwt.RegisteredClaims, error) {

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return j.Config.GetSecret(), nil
	})
	if err != nil {
		return claims, err
	}

	if token.Valid == false {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("no valid token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errors.New("token expire")
			} else {
				return nil, errors.New(fmt.Sprintf("Couldn't handle this token: %s", err))
			}
		} else {
			return nil, errors.New(fmt.Sprintf("Couldn't handle this token: %s", err))
		}
	}

	if issuer != "" {
		if claims.VerifyIssuer(issuer, true) == false {
			return nil, errors.New("token issuer error")
		}
	}

	return claims, nil
}

func (j *BaseToken) ParseID(claims *jwt.RegisteredClaims) (int, error) {
	return strconv.Atoi(claims.Subject)
}
