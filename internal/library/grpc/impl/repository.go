package impl

import (
	"database/sql"
	"github.com/go-chi/jwtauth"
	"github.com/rs/zerolog"
	"go8ddd/configs"
	"go8ddd/internal/library/grpc"
	"go8ddd/internal/server/server"
)

type RepositoryServiceGrpcImpl struct {
	DB     *sql.DB
	JWT    *jwtauth.JWTAuth
	Config *configs.Configs
	Log    zerolog.Logger
	Grpc   *grpc.Grpc
}



func NewRepositoryServiceGrpcImpl(s *server.Server) *RepositoryServiceGrpcImpl {
	return &RepositoryServiceGrpcImpl{
		DB:     s.DbConn,
		JWT:    s.JWT,
		Config: s.Cfg,
		Log: s.Log,
		Grpc: s.GrpcServer,
	}
}