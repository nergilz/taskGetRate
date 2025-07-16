package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env           string     `yaml:"env" env-default:"local"`
	GrpcCfg       GRPCConfig `yaml:"grpc"`
	GrinexBaseUrl string     `yaml:"grinex_api_url" env-required:"true"`
}

type GRPCConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type HttpServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timout"`
}

func Load() *Config {
	var cfg Config

	if _, err := os.Stat("./config/config.yaml"); os.IsNotExist(err) {
		panic("config file does not exist")
	}

	err := cleanenv.ReadConfig("./config/config.yaml", &cfg)
	if err != nil {
		panic("config is empty: " + err.Error())
	}

	return &cfg
}
