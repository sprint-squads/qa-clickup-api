package api

import (
	"github.com/Sprint-Squads/qa-clickup-api/internal/api/handlers"
	"github.com/Sprint-Squads/qa-clickup-api/pkg/application"
	"github.com/gin-gonic/gin"
)

// New - creates new instance of gin.Engine
func New(app application.Application) (*gin.Engine, error) {
	router := gin.Default()
	handler := handlers.Get(app)

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, "pong")
		})

		v1.GET("/welcome", handler.Welcome)

	}

	return router, nil
}
