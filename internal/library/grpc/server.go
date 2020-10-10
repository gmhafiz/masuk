package grpc

import (
	"fmt"
	"net"
	"strconv"

	"github.com/rs/zerolog"

	"go8ddd/configs"
	"google.golang.org/grpc"
)

// Grpc struct holds all the dependencies required for starting Grpc server
type Grpc struct {
	//tokenAuth *jwtauth.JWTAuth
	//db        *datastore.DBConn
	Server   *grpc.Server
	Listener net.Listener
	//cfg       *Config
}

func New(log zerolog.Logger, cfg *configs.Grpc) *Grpc {
	portNumber, _ := strconv.ParseUint(cfg.Port, 10, 64)
	netListener := getNetListener(portNumber, log)
	grpcServer := grpc.NewServer()

	return &Grpc{
		Server:   grpcServer,
		Listener: netListener,
	}
}

//func (g *Grpc) Start() error {
//	if err := api.GrpcServer.Server.Serve(api.GrpcServer.Listener); err != nil {
//		log.Fatalf("failed to serve: %s", err)
//		return err
//	}
//
//	return nil
//}


func getNetListener(port uint64, logger zerolog.Logger) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Fatal().Msgf("failed to listen: %v", err)
	}

	return lis
}
