package application

import (
	"github.com/minio/minio-go/v7"
	"github.com/sprint-squads/qa-clickup-api/pkg/config"
	minio_client "github.com/sprint-squads/qa-clickup-api/pkg/minio"
	"github.com/sprint-squads/qa-clickup-api/pkg/services"
)

// Application model
type Application struct {
	Config *config.Config
	Services      *services.Services
	MinIOClient *minio.Client
}

// Get - Application initializer
func Get() (*Application, error) {
	config := config.Get()
	services, err := services.Get(config)
	if err != nil {
		return nil, err
	}

	client := minio_client.Get(config)
	return &Application{
		Config: config,
		Services: services,
		MinIOClient: client,
	}, nil
}
