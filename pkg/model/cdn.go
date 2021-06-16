package model

//type CDNUploadResponse struct {
//	Data models.File `json:"data"`
//}

type CDNUploadDirectResponse struct {
	BaseResponse
	Data string `json:"data"`
}

//type FilesResponse struct {
//	BaseResponse
//	Data  []models.File `json:"data"`
//	Page  int           `json:"page"`
//	Size  int           `json:"size"`
//	Total int           `json:"total"`
//}

type BaseResponse struct {
	HTTPStatus int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}