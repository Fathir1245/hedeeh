package repository

// HealthRepository is a placeholder repository
// used to demonstrate layered architecture.
type HealthRepository struct{}

func NewHealthRepository() *HealthRepository {
	return &HealthRepository{}
}
