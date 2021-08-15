package domain

type HealthCheckRepository interface {
	GetHealthInfo() (HealthCheck, error)
	StoreHealthInfo(HealthCheck) (HealthCheck, error)
}
