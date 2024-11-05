package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP    ConfigHTTP    `yaml:"http"`
		MongoDB ConfigMongoDB `yaml:"mongodb"`
	}

	ConfigHTTP struct {
		Port    int           `yaml:"port" env:"HTTP_PORT" env-required:"true"`
		Timeout time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT"`
	}

	ConfigMongoDB struct {
		User     string `yaml:"user" env:"MONGODB_USER"`
		Password string `yaml:"password" env:"MONGODB_PASSWORD"`
		DB       string `yaml:"database" env:"MONGODB_DB"`
		Hostname string `yaml:"hostname" env:"MONGODB_HOST"`
		Port     int    `yaml:"port" env:"MONGODB_PORT"`
	}
)

func MustLoad() *Config {
	var cfg Config
	path := fetchConfigPath()
	if path != "" {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			panic("config file doesn't exist: " + path)
		}
		if err := cleanenv.ReadConfig(path, &cfg); err != nil {
			panic("failed to read config file: " + err.Error())
		}
		return &cfg
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("failed to read config from env")
	}

	return &cfg
}

func fetchConfigPath() string {
	var result string
	flag.StringVar(&result, "config", "", "path to config file")
	flag.Parse()
	if result == "" {
		result = os.Getenv("CONFIG_PATH")
	}
	return result
}
