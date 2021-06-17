package clickup

import (
	"encoding/json"
	"fmt"
	"github.com/sprint-squads/qa-clickup-api/internal/api/presenter"
	"github.com/sprint-squads/qa-clickup-api/pkg/application"
	"github.com/sprint-squads/qa-clickup-api/pkg/helpers"
	http_v2 "github.com/sprint-squads/qa-clickup-api/pkg/http"
	"github.com/sprint-squads/qa-clickup-api/pkg/minio"
	"mime/multipart"
	"net/http"
	"time"
)

type Service struct {
	App  application.Application
}

func NewService(app application.Application) *Service {
	return &Service{
		App:  app,
	}
}

func (s *Service) GetTags() (response *presenter.TagsList, err error) {
	url := fmt.Sprintf("%v/space/%v/tag", s.App.Config.Clickup.Url, s.App.Config.Clickup.SpaceId )
	headers := make(map[string]string)
	headers["Authorization"] = s.App.Config.Clickup.Token

	httpStatus, _, err := http_v2.RequestJSON(http.MethodGet, url, nil, headers, &response)
	if err != nil {
		return
	}
	response.HTTPStatus = httpStatus
	return
}

func (s *Service) CreateTask(body presenter.ClickUpTask) (response presenter.BaseResponse, err error) {
	url := fmt.Sprintf("%v/list/%v/task", s.App.Config.Clickup.Url, s.App.Config.Clickup.ListId)
	headers := make(map[string]string)
	headers["Authorization"] = s.App.Config.Clickup.Token

	bodyByte, err := json.Marshal(body)
	if err != nil {
		return
	}

	httpStatus, _, err := http_v2.RequestJSON(http.MethodPost, url, bodyByte, headers, nil)
	if err != nil {
		return
	}

	response.HTTPStatus = httpStatus
	response.Message = "task created"
	return
}

func (m *Service) UploadFile(file *multipart.FileHeader) error {
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