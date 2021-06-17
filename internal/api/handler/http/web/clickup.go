package http

import (
	"fmt"
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
		clickup.POST("/issues", clickupHandler.CreateIssue)
	}
}

func (h *ClickupHandler) CreateIssue(ctx *gin.Context) {

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

	taskRequest.Description = fmt.Sprintf("%v\n\nAttachments:\n", taskRequest.Description)
	if form == nil || form.File == nil || len(form.File) == 0 {
		// skip uploading file
	} else {
		filesList := form.File
		if len(filesList) > 0 {
			for _, file := range filesList["file"] {
				response, err := h.ClickupService.UploadFile(file)
				if err != nil {
					ctx.JSON(400, err)
					return
				}
				taskRequest.Description += fmt.Sprintf("%v\n", response)
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

	res, err := h.ClickupService.CreateIssue(clickUpTask)
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