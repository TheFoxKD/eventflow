package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/thefoxkd/eventflow/internal/entity"
)

type SubscriptionRepository interface {
	// Create a new subscription
	Create(ctx context.Context, subscription *entity.Subscription) error

	// Get a subscription by ID
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Subscription, error)

	// Update a subscription
	Update(ctx context.Context, subscription *entity.Subscription) error

	// Delete a subscription
	Delete(ctx context.Context, id uuid.UUID) error

	// Get by user ID
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.Subscription, error)

	// Get by category
	GetByCategory(ctx context.Context, category entity.Category) ([]*entity.Subscription, error)

	// Get by user and category
	GetByUserAndCategory(ctx context.Context, userID uuid.UUID, category entity.Category) ([]*entity.Subscription, error)

	// Exists by user and category
	ExistsByUserAndCategory(ctx context.Context, userID uuid.UUID, category entity.Category) (bool, error)

	// Delete by user and category
	DeleteByUserAndCategory(ctx context.Context, userID uuid.UUID, category entity.Category) error

	// List by user ID
	ListByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.Subscription, error)
}
