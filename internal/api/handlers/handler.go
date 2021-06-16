package handlers

import (
	"github.com/Sprint-Squads/qa-clickup-api/pkg/application"
)

// Handler model
type Handler struct {
	App application.Application
}

// Get - Handler initializer
func Get(app application.Application) *Handler {
	var handler Handler
	handler.App = app

	return &handler
}
