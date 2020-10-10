package main

import (
	"go8ddd/configs"
	"go8ddd/internal/domain/grpc_gen"
	"go8ddd/internal/library/grpc"
	"go8ddd/internal/library/grpc/impl"
	"go8ddd/internal/library/jwt"
	"go8ddd/internal/library/logger"
	"go8ddd/internal/server/server"
)

const Version = "v0.2.0"


func main() {
	log := logger.New(Version)
	cfg := configs.New(log)
	db := configs.NewDatabase(log, cfg)
	tokenAuth := jwt.New(cfg.Jwt)
	g := grpc.New(log, cfg.Grpc)

	s := server.New(log, db, g, cfg, tokenAuth)


	repositoryServiceImpl := impl.NewRepositoryServiceGrpcImpl(s)
	grpc_gen.RegisterUserServiceServer(s.GrpcServer.Server, repositoryServiceImpl)
	//service.RegisterUserServiceServer(s.GrpcServer.Server, nil)

	err := s.Start()
	if err != nil {
		panic(err)
	}
}