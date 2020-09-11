package api

import (
	"github.com/go-chi/jwtauth"
	"log"
	"masuk/internal/configs"
	"masuk/internal/datastore"
	"masuk/internal/grpc"
)

type API struct {
	GrpcServer *grpc.Grpc
	DbConn     *datastore.DBConn
	JwtAuth    *jwtauth.JWTAuth
	Cfg        *configs.Configs
}

func New(grpcServer *grpc.Grpc, db *datastore.DBConn, jwtAuth *jwtauth.JWTAuth, cfg *configs.Configs) (*API, error) {
	return &API{
		GrpcServer: grpcServer,
		DbConn:     db,
		JwtAuth:    jwtAuth,
		Cfg:        cfg,
	}, nil
}

func (api *API) StartGrpc() error {
	if err := api.GrpcServer.Server.Serve(api.GrpcServer.Listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err
	}

	return nil
}
