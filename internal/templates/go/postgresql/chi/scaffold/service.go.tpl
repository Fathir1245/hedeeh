package service

import (
	"context"
	"{{.ModuleName}}/internal/model"
)

// UserService defines the interface for user-related operations.
type UserService interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, userID int64) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, userID int64) error
}

// UserImpl is the implementation of UserService.
type UserImpl struct {
	repo Repository
}

// Create creates a new user.
func (s *UserImpl) Create(ctx context.Context, user *model.User) error {
	return nil
}

// GetByID retrieves a user by its ID.
func (s *UserImpl) GetByID(ctx context.Context, userID int64) (*model.User, error) {
	return nil
}

// Update updates an existing user.
func (s *UserImpl) Update(ctx context.Context, user *model.User) error {
	return nil
}

// Delete deletes a user by its ID.
func (s *UserImpl) Delete(ctx context.Context, userID int64) error {
	return nil
}