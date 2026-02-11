package public

// HealthUseCase abstracts the health-related business operations used by delivery.
type HealthUseCase interface {
	Health() (status, app, env string)
	HealthDB() (status, app, env string)
}
