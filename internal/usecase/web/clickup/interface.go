package clickup

import (
	"github.com/sprint-squads/qa-clickup-api/internal/api/presenter"
	"mime/multipart"
)

type UseCase interface {
	GetTags()(response *presenter.TagsList, err error)
	CreateTask(body presenter.ClickUpTask) (response presenter.BaseResponse, err error)
	UploadFile(file *multipart.FileHeader) error
}