package grpc

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

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


// Config holds all the configuration required to start the Grpc server
type Config struct {
	Host         string        `yaml:"HOST"`
	Port         string        `yaml:"PORT"`
	ReadTimeout  time.Duration `yaml:"READ_TIMEOUT"`
	WriteTimeout time.Duration `yaml:"WRITE_TIMEOUT"`
	DialTimeout  time.Duration `yaml:"DIAL_TIMEOUT"`
}

func New(cfg *Config) *Grpc {

	portNumber, _ := strconv.ParseUint(cfg.Port, 10, 64)
	netListener := getNetListener(portNumber)
	grpcServer := grpc.NewServer()

	return &Grpc{
		Server:   grpcServer,
		Listener: netListener,
	}
}

func getNetListener(port uint64) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
