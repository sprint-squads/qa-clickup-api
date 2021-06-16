package model

//type CDNUploadResponse struct {
//	Data models.File `json:"data"`
//}

type CDNUploadDirectResponse struct {
	BaseResponse
	Data string `json:"data"`
}

type BaseResponse struct {
	HTTPStatus int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}