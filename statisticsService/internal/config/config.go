package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string            `yaml:"env" env-default:"local"`
	Gin         GinConfig         `yaml:"gin"`
	GRPC        GRPCConfig        `yaml:"grpc"`
	Storage     StorageConfig     `yaml:"storage"`
	Kafka       KafkaConfig       `yaml:"kafka"`
	PostService PostServiceConfig `yaml:"post_service"`
	MainService MainServiceConfig `yaml:"main_service"`
}

type MainServiceConfig struct {
	Host string `yaml:"host" env-default:"auth_service"`
	Port string `yaml:"port" env-default:"8080"`
}

type PostServiceConfig struct {
	Host string `yaml:"host" env-default:"post_service"`
	Port string `yaml:"port" env-default:"50052"`
}

type GRPCConfig struct {
	Host    string        `yaml:"host" env-default:"0.0.0.0"`
	Port    string        `yaml:"port" env-default:"50053"`
	Timeout time.Duration `yaml:"timeout"`
}

type KafkaConfig struct {
	Address      string `yaml:"address"`
	ConsumerPort string `yaml:"consumer_port"`
	Topic        string `yaml:"topic"`
}

type GinConfig struct {
	Host string `yaml:"host" env-default:"0.0.0.0"`
	Port string `yaml:"port" env-default:"8082"`
}

type StorageConfig struct {
	User     string `yaml:"user" env-default:"default"`
	Password string `yaml:"password" env-default:""`
	Name     string `yaml:"name" env-default:"default"`
	Host     string `yaml:"host" env-default:"stats_database"`
	Port     string `yaml:"port" env-default:"9000"`
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
