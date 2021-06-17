package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sprint-squads/qa-clickup-api/internal/api/handlers"
	"github.com/sprint-squads/qa-clickup-api/pkg/application"
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

		clickup := v1.Group("/clickup")
		{
			clickup.GET("/tags", handler.GetTags)
			clickup.POST("/create-issues", handler.CreateTask)
		}

	}

	return router, nil
}
