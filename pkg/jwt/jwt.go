package jwt

import (
	"github.com/go-chi/jwtauth"

	"masuk/internal/jwt"
)

type TokenAuth struct {
	tokenAuth *jwtauth.JWTAuth
}

func New(config jwt.Jwt) *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(config.Secret), nil)
}
