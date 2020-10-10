package server

import (
	"database/sql"
	"github.com/go-chi/jwtauth"

	"github.com/rs/zerolog"

	"go8ddd/configs"
	"go8ddd/internal/library/grpc"
)

type Server struct {
	GrpcServer *grpc.Grpc
	//JwtAuth *jwtauth.JWTAuth
	Cfg    *configs.Configs
	DbConn *sql.DB
	Log    zerolog.Logger
	JWT    *jwtauth.JWTAuth
}

func New(log zerolog.Logger, db *sql.DB, g *grpc.Grpc, cfg *configs.Configs, auth *jwtauth.JWTAuth) *Server {
	return &Server{
		Log: log,
		GrpcServer: g,
		DbConn:     db,
		Cfg:        cfg,
		JWT: auth,
	}
}

func (s *Server) Start() error {
	s.Log.Info().Msgf("server started at %s:%s", s.Cfg.Grpc.Host, s.Cfg.Grpc.Port)
	err := s.GrpcServer.Server.Serve(s.GrpcServer.Listener)
	if err != nil {
		s.Log.Fatal().Msg(err.Error())
		return err
	}

	return nil
}