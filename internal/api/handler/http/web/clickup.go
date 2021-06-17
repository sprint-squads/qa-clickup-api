package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sprint-squads/qa-clickup-api/internal/api/presenter"
	"github.com/sprint-squads/qa-clickup-api/internal/usecase/web/clickup"
	"github.com/sprint-squads/qa-clickup-api/pkg/application"
	"strings"
)

type ClickupHandler struct {
	App            application.Application
	ClickupService clickup.UseCase
}

func NewClickupHandler(r *gin.RouterGroup, App application.Application, clickupService clickup.UseCase) {
	clickupHandler := ClickupHandler {
		App: App,
		ClickupService: clickupService,
	}

	clickup := r.Group("/clickup")
	{
		clickup.GET("/tags", clickupHandler.GetTags)
		clickup.POST("/issues", clickupHandler.CreateTask)
	}
}

func (h *ClickupHandler) CreateTask(ctx *gin.Context) {

	// Parse request body
	var taskRequest presenter.ClickUpTaskRequest
	err := ctx.Bind(&taskRequest)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	tags := strings.Split(taskRequest.Tags, " ")
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	//taskRequest.Description = fmt.Sprintf("Email: %v\n\n%v\n\n", taskRequest.Email, taskRequest.Description)
	if form == nil || form.File == nil || len(form.File) == 0 {
		// skip uploading file
	} else {
		filesList := form.File
		if len(filesList) > 0 {
			file := form.File["file"][0]

			errResp := h.ClickupService.UploadFile(file)
			if errResp != nil {
				ctx.JSON(400, err)
				return
			}
		}
	}

	clickUpTask := presenter.ClickUpTask{
		Name:                      taskRequest.Title,
		Description:               taskRequest.Description,
		Assignees:                 []int{5723639},
		Priority:				   taskRequest.Priority,
		Tags:                      tags,
	}

	res, err := h.ClickupService.CreateTask(clickUpTask)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(res.HTTPStatus, res)
}

func (h *ClickupHandler) GetTags(ctx *gin.Context) {
	res, err := h.ClickupService.GetTags()
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(res.HTTPStatus, res)
}