package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/thefoxkd/eventflow/internal/entity"
)

type UserRepository interface {
	// Create stores new user in database
	Create(ctx context.Context, user *entity.User) error

	// GetByID retrieves a user by unique identifier
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)

	// GetByEmail finds user by email address for authentication
	GetByEmail(ctx context.Context, email string) (*entity.User, error)

	// Update modifies existing user data
	Update(ctx context.Context, user *entity.User) error

	// Delete removes user from database
	Delete(ctx context.Context, id uuid.UUID) error

	// ExistsByEmail checks if user with given email already exists
	ExistsByEmail(ctx context.Context, email string) (bool, error)

	// List retrieves all users with pagination support
	List(ctx context.Context, limit, offset int) ([]*entity.User, error)
}
