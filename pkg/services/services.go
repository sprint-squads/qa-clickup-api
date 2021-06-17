package services

import (
	"github.com/sprint-squads/qa-clickup-api/pkg/config"
)

type Services struct {
	Config *config.Config
}

func Get(config *config.Config) (*Services, error) {

	return &Services{
		Config: config,
	}, nil
}
