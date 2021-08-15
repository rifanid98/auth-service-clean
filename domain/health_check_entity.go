package domain

type (
	HealthCheck struct {
		Status string
	}
)

func NewHealthCheck() HealthCheck {
	return HealthCheck{"Ok"}
}
