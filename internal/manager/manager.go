package manager

import (
	"github.com/sprint-squads/qa-clickup-api/pkg/application"
)

type Manager struct {
	App *application.Application
}

// Get - creates new Manager instance
func Get(app *application.Application) (*Manager, error) {
	return &Manager{App: app}, nil
}
