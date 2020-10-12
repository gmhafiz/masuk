package grpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"go8ddd/configs"
	"go8ddd/internal/domain/grpc_gen"
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

	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()
	//
	//mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithInsecure()}
	//addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	//err := grpc_gen.RegisterUserServiceHandlerFromEndpoint(ctx, mux, addr, opts)
	//if err != nil {
	//	log.Panic().Msg(err.Error())
	//}

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

func RunGateway(dialAddr string, logger zerolog.Logger) error {
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithInsecure(),
		//grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		//grpc.WithBlock(),
	)
	if err != nil {
		logger.Fatal().Msg(err.Error())
		return err
	}
	
	gwmux := runtime.NewServeMux()
	err = grpc_gen.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		logger.Fatal().Msg(err.Error())
		return err
	}

	gatewayAddr := "0.0.0.0:7090"

	gwServer := &http.Server{
		Addr:    gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				gwmux.ServeHTTP(w, r)
				return
			}
		}),
	}

	logger.Info().Msgf("Serving gRPC-Gateway http://%s", gatewayAddr)
	err = gwServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
