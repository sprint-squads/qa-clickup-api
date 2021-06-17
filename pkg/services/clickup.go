package services

import (
	"encoding/json"
	"fmt"
	http_v2 "github.com/sprint-squads/qa-clickup-api/pkg/http"
	"github.com/sprint-squads/qa-clickup-api/pkg/model"
	"net/http"
)

// CreateClickUpTask - creates task using clickup api
func (s *Services) CreateTask(body model.ClickUpTask) (response model.BaseResponse, err error) {
	url := fmt.Sprintf("%v/list/%v/task", s.Config.Clickup.Url, s.Config.Clickup.ListId)
	headers := make(map[string]string)
	headers["Authorization"] = s.Config.Clickup.Token

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

func (s *Services) GetTags() (response *model.TagsList, err error) {
	url := fmt.Sprintf("%v/space/%v/tag", s.Config.Clickup.Url, s.Config.Clickup.SpaceId )
	headers := make(map[string]string)
	headers["Authorization"] = s.Config.Clickup.Token

	httpStatus, _, err := http_v2.RequestJSON(http.MethodGet, url, nil, headers, &response)
	if err != nil {
		return
	}
	response.HTTPStatus = httpStatus
	return
}
