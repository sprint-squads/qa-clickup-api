package handlers

import (
	"github.com/Sprint-Squads/qa-clickup-api/pkg/model"
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *Handler) CreateTask(ctx *gin.Context) {

	// Parse request body
	var taskRequest model.ClickUpTaskRequest
	err := ctx.Bind(&taskRequest)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	tags := strings.Split(taskRequest.Tags, " ")
	//form, err := ctx.MultipartForm()
	//if err != nil {
	//	ctx.JSON(400, err)
	//	return
	//}

	//taskRequest.Description = fmt.Sprintf("Email: %v\n\n%v\n\n", taskRequest.Email, taskRequest.Description)
	//if form == nil || form.File == nil || len(form.File) == 0 {
	//	// skip uploading file
	//} else {
	//	filesList := form.File
	//	if len(filesList) > 0 {
	//		file := form.File["file"][0]
	//
	//		response, errResp := h.Manager.UploadFileDirect(taskRequest.FolderName, file)
	//		if errResp != nil {
	//			ctx.JSON(400, err)
	//			return
	//		}
	//
	//		taskRequest.Description += fmt.Sprintf("Attachment:%v", response.Data)
	//	}
	//}

	//now := time.Now()

	clickUpTask := model.ClickUpTask{
		Name:                      taskRequest.Title,
		Description:               taskRequest.Description,
		Assignees:                 []int{5723639},
		Priority:				   taskRequest.Priority,
		Tags:                      tags,
	}

	res, err := h.Manager.CreateTask(clickUpTask)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(res.HTTPStatus, res)
}

func (h *Handler) GetTags(ctx *gin.Context) {
	res, err := h.Manager.GetTags()
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(res.HTTPStatus, res)
}