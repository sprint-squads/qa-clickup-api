package manager

import (
	"github.com/Sprint-Squads/qa-clickup-api/pkg/model"
	"mime/multipart"
)

func (m *Manager) UploadFileDirect(folderName string, file *multipart.FileHeader) (*model.CDNUploadDirectResponse, error) {
	return m.App.Services.UploadFileDirect(folderName, file)
}

