package configs

import (
	"os"
)

type Jwt struct {
	Secret       string
	//Host         string
	//Port         string
	//ReadTimeout  time.Duration
	//WriteTimeout time.Duration
	//DialTimeout  time.Duration
}

func NewJWT() *Jwt {
	//readTimeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	//if err != nil {
	//	readTimeout = 600
	//}
	//
	//writeTimeout, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	//if err != nil {
	//	writeTimeout = 600
	//}
	//
	//dialTimeout, err := strconv.Atoi(os.Getenv("DIAL_TIMEOUT"))
	//if err != nil {
	//	writeTimeout = 600
	//}

	return &Jwt{
		Secret:       os.Getenv("JWT_SECRET"),
	//	Host:         os.Getenv("JWT_HOST"),
	//	Port:         os.Getenv("JWT_PORT"),
	//	ReadTimeout:  time.Duration(readTimeout),
	//	WriteTimeout: time.Duration(writeTimeout),
	//	DialTimeout:  time.Duration(dialTimeout),
	}
}
