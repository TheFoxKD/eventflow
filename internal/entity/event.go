package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID  `json:"id" db:"id" validate:"required,uuid4"`
	UserID      uuid.UUID  `json:"userId" db:"user_id" validate:"required,uuid4"`
	Category    Category   `json:"category" db:"category" validate:"required"`
	Title       string     `json:"title" db:"title" validate:"required,max=255"`
	Description *string    `json:"description,omitempty" db:"description" validate:"omitempty,max=1000"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" db:"updated_at"`
}

func (Event) TableName() string {
	return "events"
}

func (e Event) String() string {
	desc := NilValue
	if e.Description != nil {
		desc = *e.Description
	}
	updatedAt := NilValue
	if e.UpdatedAt != nil {
		updatedAt = e.UpdatedAt.Format(time.DateTime)
	}

	createdAt := e.CreatedAt.Format(time.DateTime)
	return fmt.Sprintf("Event(ID: %s, UserID: %s, Title: %s, Description: %s, Category: %s, CreatedAt: %s, UpdatedAt: %s)",
		e.ID, e.UserID, e.Title, desc, e.Category, createdAt, updatedAt)
}
