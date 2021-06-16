package application

import (
	"github.com/Sprint-Squads/qa-clickup-api/pkg/config"
)

// Application model
type Application struct {
	Config *config.Config
}

// Get - Application initializer
func Get() (*Application, error) {
	config := config.Get()

	return &Application{
		Config: config,
	}, nil
}
