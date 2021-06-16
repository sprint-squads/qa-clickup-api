package services

import (
	"encoding/json"
	"fmt"
	http_v2 "github.com/Sprint-Squads/qa-clickup-api/pkg/http"
	"github.com/Sprint-Squads/qa-clickup-api/pkg/model"
	"net/http"
	"os"
)

// CreateClickUpTask - creates task using clickup api
func (s *Services) CreateTask(body model.ClickUpTask) (response *model.BaseResponse, err error) {
	clickupUrl := os.Getenv("CLICKUP_URL")
	listId := os.Getenv("LIST_ID")
	url := fmt.Sprintf("%v/list/%v/task", clickupUrl, listId)
	headers := make(map[string]string)
	headers["Authorization"] = "3851228_2087f0167d551169256f2f86e8b21fec4bc90075"

	bodyByte, err := json.Marshal(body)
	if err != nil {
		return
	}

	_, responseBody, err := http_v2.RequestJSON(http.MethodPost, url, bodyByte, headers, &response)
	if err != nil {
		return
	}
	fmt.Println("responseBody", string(responseBody))
	return
}

func (s *Services) GetTags() (response *model.TagsList, err error) {
	clickupUrl := os.Getenv("CLICKUP_URL")
	spaceId := os.Getenv("SPACE_ID")

	url := fmt.Sprintf("%v/space/%v/tag", clickupUrl, spaceId )
	headers := make(map[string]string)
	headers["Authorization"] = "3851228_2087f0167d551169256f2f86e8b21fec4bc90075"

	_, _, err = http_v2.RequestJSON(http.MethodGet, url, nil, headers, &response)
	if err != nil {
		return
	}
	response.HTTPStatus = 200
	return
}
