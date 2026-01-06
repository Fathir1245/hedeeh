package repository

import (
	"context"
	"database/sql"
	"{{.ModuleName}}/internal/model"
)


// UserRepository defines the interface for user operations.
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int) error
}

// UserRepositoryImpl provides a concrete implementation of UserRepository using SQL.
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepositoryImpl.
func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// Create adds a new user to the database.
func (r *UserRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	// Implementation goes here
	return nil
}

// GetByID retrieves a user by their ID from the database.
func (r *UserRepositoryImpl) GetByID(ctx context.Context, id int) (*model.User, error) {
	// Implementation goes here
	return nil, nil
}

// Update updates an existing user in the database.
func (r *UserRepositoryImpl) Update(ctx context.Context, user *model.User) error {
	// Implementation goes here
	return nil
}

// Delete removes a user from the database by their ID.
func (r *UserRepositoryImpl) Delete(ctx context.Context, id int) error {
	// Implementation goes here
	return nil
}