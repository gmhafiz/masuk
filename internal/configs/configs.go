package configs

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v3"

	"masuk/internal/datastore"
	"masuk/internal/grpc"
	"masuk/internal/jwt"
)

type Configs struct {
	Grpc     grpc.Config        `yaml:"Grpc"`
	Database datastore.Database `yaml:"Database"`
	JWT      jwt.Jwt            `yaml:"Jwt"`
}

func (cfg *Configs) GRPC() (*grpc.Config, error) {
	return &grpc.Config{
		Port:         cfg.Grpc.Port,
		ReadTimeout:  time.Second * 9995,
		WriteTimeout: time.Second * 9995,
		DialTimeout:  time.Second * 9993,
	}, nil
}

func (cfg *Configs) DataStore() (*datastore.Database, error) {
	return &datastore.Database{
		Driver:  cfg.Database.Driver,
		Host:    cfg.Database.Host,
		Port:    cfg.Database.Port,
		Name:    cfg.Database.Name,
		User:    cfg.Database.User,
		Pass:    cfg.Database.Pass,
		SslMode: cfg.Database.SslMode,
	}, nil
}

func (cfg *Configs) Jwt() (jwt.Jwt, error) {
	return jwt.Jwt{
		Secret: cfg.JWT.Secret,
	}, nil
}

func New(file string) (*Configs, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cfg := &Configs{}
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
