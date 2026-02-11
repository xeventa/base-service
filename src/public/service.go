package public

import "github.com/rs/zerolog"

// HealthService provides health-related business logic.
type HealthService struct {
	Logger zerolog.Logger
	DB     DBPinger // optional
	App    string
	Env    string
}

func NewHealthService(logger zerolog.Logger, db DBPinger, app, env string) *HealthService {
	return &HealthService{Logger: logger, DB: db, App: app, Env: env}
}

func (s *HealthService) Health() (status, app, env string) {
	return "ok", s.App, s.Env
}

func (s *HealthService) HealthDB() (status, app, env string) {
	if s.DB == nil {
		return "skipped", s.App, s.Env
	}
	if err := s.DB.Ping(); err != nil {
		s.Logger.Error().Err(err).Msg("db ping failed")
		return "error", s.App, s.Env
	}
	return "ok", s.App, s.Env
}
