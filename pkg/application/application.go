package application

import (
	"github.com/minio/minio-go/v7"
	"github.com/sprint-squads/qa-clickup-api/pkg/config"
	minio_client "github.com/sprint-squads/qa-clickup-api/pkg/minio"
)

// Application model
type Application struct {
	Config *config.Config
	MinIOClient *minio.Client
}

// Get - Application initializer
func Get() (*Application, error) {
	config := config.Get()

	client := minio_client.Get(config)
	return &Application{
		Config: config,
		MinIOClient: client,
	}, nil
}
