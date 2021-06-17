package manager

import (
	"github.com/sprint-squads/qa-clickup-api/pkg/helpers"
	"github.com/sprint-squads/qa-clickup-api/pkg/minio"
	"github.com/sprint-squads/qa-clickup-api/pkg/model"
	"mime/multipart"
	"time"
)

func (m *Manager) GetTags() (*model.TagsList, error) {
	return m.App.Services.GetTags()
}

func (m *Manager) CreateTask(body model.ClickUpTask) (model.BaseResponse, error) {
	return m.App.Services.CreateTask(body)
}

func (m *Manager) UploadFile(file *multipart.FileHeader) error {
	fileExt := helpers.GetFileExt(file.Filename)
	randomFileName := helpers.RandStringRunes(16)
	objectName := time.Now().Format("20060102") + "_" + randomFileName + "." + fileExt
	contentType := "image/" + fileExt
	_, err := minio.UploadToMinio(m.App.MinIOClient, m.App.Config.Minio.BucketName, objectName, file, contentType)
	if err != nil {
		return err
	}
	return nil
}
