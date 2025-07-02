package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/thefoxkd/eventflow/internal/entity"
	"github.com/thefoxkd/eventflow/internal/repository/filters"
)

type EventRepository interface {
	// Create stores new event in database
	Create(ctx context.Context, event *entity.Event) error

	// GetByID retrieves event by unique identifier
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Event, error)

	// Update modifies existing event data
	Update(ctx context.Context, event *entity.Event) error

	// Delete removes event from database
	Delete(ctx context.Context, id uuid.UUID) error

	// Retrieve all events for a user
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*entity.Event, error)

	// Retrieve all events for a category
	GetByCategory(ctx context.Context, category entity.Category, limit, offset int) ([]*entity.Event, error)

	// Retrieve events by filter and pagination
	List(ctx context.Context, filters filters.EventFilter, limit, offset int) ([]*entity.Event, error)

	// Search by text on title and description
	SearchByText(ctx context.Context, query string, limit, offset int) ([]*entity.Event, error)

	// Total count for pagination
	Count(ctx context.Context, filters filters.EventFilter) (int, error)

	// Count by category
	CountByCategory(ctx context.Context, category entity.Category) (int, error)
}
