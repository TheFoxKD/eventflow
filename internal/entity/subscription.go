package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID        uuid.UUID  `json:"id" db:"id" validate:"required,uuid4"`
	UserID    uuid.UUID  `json:"userId" db:"user_id" validate:"required,uuid4"`
	Category  Category   `json:"category" db:"category" validate:"required"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}

func (s Subscription) String() string {
	updatedAt := NilValue
	if s.UpdatedAt != nil {
		updatedAt = s.UpdatedAt.Format(time.DateTime)
	}

	createdAt := s.CreatedAt.Format(time.DateTime)
	return fmt.Sprintf("Subscription(ID: %s, UserID: %s, Category: %s, CreatedAt: %s, UpdatedAt: %s)", s.ID, s.UserID, s.Category, createdAt, updatedAt)
}
