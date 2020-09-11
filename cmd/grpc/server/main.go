package main

import (
	"flag"

	"masuk/internal/api"
	"masuk/internal/configs"
	"masuk/internal/datastore"
	grpcAPI "masuk/internal/grpc"
	"masuk/internal/grpc/impl"
	"masuk/internal/grpc/service"
	jwtPkg "masuk/pkg/jwt"
)

var flagConfig = flag.String("config", "./config/dev.yml", "path to the config file")

func main() {
	cfg, err := configs.New(*flagConfig)
	if err != nil {
		panic(err)
	}

	dataStoreCfg, err := cfg.DataStore()
	if err != nil {
		panic(err)
	}

	db, err := datastore.New(dataStoreCfg)
	if err != nil {
		panic(err)
	}

	tokenAuth := jwtPkg.New(cfg.JWT)

	grpcConfig, err := cfg.GRPC()
	if err != nil {
		panic(err)
	}

	grpcServer := grpcAPI.New(grpcConfig)

	apiServer, err := api.New(grpcServer, db, tokenAuth, cfg)
	if err != nil {
		panic(err)
	}

	repositoryServiceImpl := impl.NewRepositoryServiceGrpcImpl(apiServer)
	service.RegisterUserServiceServer(apiServer.GrpcServer.Server, repositoryServiceImpl)

	err = apiServer.StartGrpc()
	if err != nil {
		panic(err)
	}
}
