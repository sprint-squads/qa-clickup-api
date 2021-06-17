package api

import (
	"github.com/gin-gonic/gin"
	webHandler "github.com/sprint-squads/qa-clickup-api/internal/api/handler/http/web"
	"github.com/sprint-squads/qa-clickup-api/internal/usecase/web/clickup"
	"github.com/sprint-squads/qa-clickup-api/pkg/application"
)

// New - creates new instance of gin.Engine
func NewRouter(app application.Application) (*gin.Engine, error) {
	router := gin.Default()
	v1 := router.Group("/v1")
	NewWebRouter(v1, app)

	return router, nil
}

func NewWebRouter(r *gin.RouterGroup, app application.Application) (web *gin.RouterGroup) {
	//services
	clickupService := clickup.NewService(app)

	// group
	web = r.Group("")

	// routes for get tags, create issues
	webHandler.NewClickupHandler(web, app, clickupService)

	return
}
