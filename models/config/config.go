package config

import (
	"github.com/rs/zerolog"
)

type Global struct {
	Api      Api
	Database Database
	Server   Server
	Log      Log
}

type Api struct {
	Key string `env:"API_KEY,required"`
}

type Server struct {
	Host string `env:"SERVER_HOST,required"`
	Port string `env:"SERVER_PORT,required"`
}

type Database struct {
	User     string `env:"MYSQL_USER,required"`
	Password string `env:"MYSQL_PASSWORD,required"`
	Host     string `env:"MYSQL_HOST,required"`
	Port     string `env:"MYSQL_PORT,required"`
	Name     string `env:"MYSQL_DBNAME,required"`
}

type Log struct {
	Level zerolog.Level `env:"LOG_LEVEL,required"`
}
