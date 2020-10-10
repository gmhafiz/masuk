package configs

import (
	"os"
	"strconv"
	"time"
)

type Grpc struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DialTimeout  time.Duration
}

func NewGrpc() *Grpc {
	readTimeout, err := strconv.Atoi(os.Getenv("GRPC_READ_TIMEOUT"))
	if err != nil {
		readTimeout = 600
	}

	writeTimeout, err := strconv.Atoi(os.Getenv("GRPC_WRITE_TIMEOUT"))
	if err != nil {
		writeTimeout = 600
	}

	dialTimeout, err := strconv.Atoi(os.Getenv("GRPC_DIAL_TIMEOUT"))
	if err != nil {
		writeTimeout = 600
	}

	return &Grpc{
		Host:         os.Getenv("GRPC_HOST"),
		Port:         os.Getenv("GRPC_PORT"),
		ReadTimeout:  time.Duration(readTimeout),
		WriteTimeout: time.Duration(writeTimeout),
		DialTimeout:  time.Duration(dialTimeout),
	}
}
