package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" json:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" json:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" json:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc" json:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" json:"port" env-default:"50051"`
	Timeout time.Duration `yaml:"timeout" json:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file doesn't exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	// --config="path/to/config.yaml"
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
	//--> we can start our programm using config path
	// first ability: CONFIG_PATH=./path/to/config/file.yaml postservice
	// second ability: postservice --config=./path...
}
