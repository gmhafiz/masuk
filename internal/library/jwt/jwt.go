package jwt

import (
	"github.com/go-chi/jwtauth"
	"go8ddd/configs"
)

type TokenAuth struct {
	tokenAuth *jwtauth.JWTAuth
}

func New(cfg *configs.Jwt) *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(cfg.Secret), nil)
}
