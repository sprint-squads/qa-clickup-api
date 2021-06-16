package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"os"
)

// Config model
type Config struct {
	Server struct {
		Host string `env:"HOST"`
		Port int    `env:"HTTP_PORT"`
	}
	Clickup struct {
		Url     string `env:"CLICKUP_URL"`
		SpaceId string `env:"CLICKUP_SPACE_ID"`
		ListId  string `env:"CLICKUP_LIST_ID"`
		Token   string `env:"CLICKUP_ACCESS_TOKEN"`
	}
	Minio struct {
		Url        string `env:"MINIO_URL"`
		AccessKey  string `env:"MINIO_ACCESS_KEY"`
		SecretKey  string `env:"MINIO_SECRET_KEY"`
		UseSSL     bool   `env:"MINIO_USE_SSL"`
		BucketName string `env:"MINIO_BUCKET_NAME"`
	}
}

// Get - Config initializer
func Get() *Config {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	var config Config
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}

	return &config

}
