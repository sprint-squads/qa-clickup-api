package manager

import (
	"github.com/Sprint-Squads/qa-clickup-api/pkg/model"
)

func (m *Manager) GetTags() (*model.TagsList, error) {
	return m.App.Services.GetTags()
}
