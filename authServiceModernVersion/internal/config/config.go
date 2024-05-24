package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string            `yaml:"env" env-default:"local"`
	Gin         GinConfig         `yaml:"gin"`
	Storage     StorageConfig     `yaml:"storage"`
	PostService PostServiceConfig `yaml:"post_service"`
}

type PostServiceConfig struct {
	Host string `yaml:"host" env-default:"0.0.0.0"`
	Port string `yaml:"port" env-default:"50052"`
}

type GinConfig struct {
	Host string `yaml:"host" env-default:"0.0.0.0"`
	Port string `yaml:"port" env-default:"8080"`
}

type StorageConfig struct {
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	Name     string `yaml:"name" env-default:"auth_db"`
	Host     string `yaml:"host" env-default:"0.0.0.0"`
	Port     string `yaml:"port" env-default:"5432"`
	Schema   string `yaml:"schema" env-default:"public"`
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
		err := godotenv.Load("./config/.env")
		if err != nil {
			log.Fatal(err)
		}

		res = os.Getenv("CONFIG_PATH")
	}

	return res
	//--> we can start our programm using config path
	// first ability: CONFIG_PATH=./path/to/config/file.yaml authservice
	// second ability: authservice --config=./path...
}
