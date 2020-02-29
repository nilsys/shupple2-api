package repository

type (
	HealthCheckRepository interface {
		CheckDBAlive() error
	}
)
