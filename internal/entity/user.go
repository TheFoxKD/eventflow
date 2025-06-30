package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" db:"id" validate:"required,uuid4"`
	Email     string     `json:"email" db:"email" validate:"required,email,max=255"`
	Password  string     `json:"-" db:"password_hash" validate:"required,min=8"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" db:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u User) String() string {
	updatedAt := NilValue
	if u.UpdatedAt != nil {
		updatedAt = u.UpdatedAt.Format(time.DateTime)
	}

	createdAt := u.CreatedAt.Format(time.DateTime)
	return fmt.Sprintf("User(ID: %s, Email: %s, CreatedAt: %s, UpdatedAt: %s)", u.ID, u.Email, createdAt, updatedAt)
}
