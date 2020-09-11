package datastore

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConn struct {
	DB *sqlx.DB
}

type Database struct {
	Driver  string `yaml:"DRIVER"`
	Host    string `yaml:"HOST"`
	Port    string `yaml:"PORT"`
	Name    string `yaml:"NAME"`
	User    string `yaml:"USER"`
	Pass    string `yaml:"PASS"`
	SslMode string `yaml:"SSL_MODE"`
}

func New(cfg *Database) (*DBConn, error) {
	dsn := fmt.Sprintf("%s://%s/%s?sslmode=%s&user=%s&password=%s",
		cfg.Driver,
		cfg.Host,
		cfg.Name,
		cfg.SslMode,
		cfg.User,
		cfg.Pass,
	)

	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer conn.Close()

	return &DBConn{
		DB: conn,
	}, nil
}
