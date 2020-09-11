package impl

import (
	"github.com/go-chi/jwtauth"
	"github.com/jmoiron/sqlx"

	"masuk/internal/api"
	"masuk/internal/configs"
)

//RepositoryServiceGrpcImpl is a implementation of RepositoryService Grpc Service.
type RepositoryServiceGrpcImpl struct {
	DbConn *sqlx.DB
	Jwt    *jwtauth.JWTAuth
	Config *configs.Configs
}

//NewRepositoryServiceGrpcImpl returns the pointer to the implementation.
func NewRepositoryServiceGrpcImpl(api *api.API) *RepositoryServiceGrpcImpl {
	return &RepositoryServiceGrpcImpl{
		DbConn: api.DbConn.DB,
		Jwt:    api.JwtAuth,
		Config: api.Cfg,
	}
}
