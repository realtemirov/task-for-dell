package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Logger   Logger
	Postgres PostgresConfig
}

type ServerConfig struct {
	AppVersion     string
	Mode           string
	Port           string
	Debug          bool
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	CtxDefaultTime time.Duration
}

type Logger struct {
	Development bool
	Level       string
	Encoding    string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	PgDriver string
}

func LoadConfig(filename string) (*Config, error) {

	var cfg Config

	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print(err.Error())
			return nil, errors.New("failed to read config file")
		}
		return nil, err
	}

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Print("unable to decode into struct")
		return nil, err
	}

	return &cfg, nil
}
