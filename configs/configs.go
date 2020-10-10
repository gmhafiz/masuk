package configs

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

type Configs struct {
	Api      *Api
	Database *Database
	Cache    *Cache
	Jwt      *Jwt
	Grpc     *Grpc
}

func New(logger zerolog.Logger) *Configs {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	api := API()
	dataStore := DataStore()
	cache := NewCache()
	jwt := NewJWT()
	grpc := NewGrpc()

	return &Configs{
		Api:      api,
		Database: dataStore,
		Cache:    cache,
		Jwt: jwt,
		Grpc: grpc,
	}
}


