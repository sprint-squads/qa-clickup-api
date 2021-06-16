package application

import (
	"github.com/Sprint-Squads/qa-clickup-api/pkg/config"
	"github.com/Sprint-Squads/qa-clickup-api/pkg/services"
	"github.com/joho/godotenv"
)

// Application model
type Application struct {
	Config *config.Config
	Services      *services.Services
}

// Get - Application initializer
func Get() (*Application, error) {
	config := config.Get()
	services, err := services.Get(config)
	if err != nil {
		return nil, err
	}

	err = godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Application{
		Config: config,
		Services: services,
	}, nil
}
