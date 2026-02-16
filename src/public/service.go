package public

import (
	"github.com/xeventa/base-service/core/environment"
)

type Service struct {
	config *environment.Config
}

func NewService(config *environment.Config) *Service {
	return &Service{config: config}
}

func (s *Service) HealthCheck() interface{} {

	data := make(map[string]interface{})

	data["status"] = "OK"
	data["config"] = s.config
	return data
}
