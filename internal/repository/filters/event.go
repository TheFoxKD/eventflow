package filters

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/thefoxkd/eventflow/internal/entity"
)

const (
	// Sort fields
	SortByCreatedAt = "createdAt"
	SortByTitle     = "title"
	SortByCategory  = "category"
	SortByRelevance = "relevance"

	// Sort orders
	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"

	// Search constraints
	SearchMinLength = 3
	SearchMaxLength = 100

	// Pagination
	DefaultLimit = 20
	MaxLimit     = 100
)

type EventFilter struct {
	UserID   *uuid.UUID       `json:"userId,omitempty" validate:"omitempty,uuid4"`
	Category *entity.Category `json:"category,omitempty"`
	Title    *string          `json:"title,omitempty" validate:"omitempty,max=255"`
	DateFrom *time.Time       `json:"dateFrom,omitempty"`
	DateTo   *time.Time       `json:"dateTo,omitempty"`

	// Search
	SearchText *string `json:"searchText,omitempty" validate:"omitempty,min=3,max=100"`

	// Sorting
	SortBy    *string `json:"sortBy,omitempty" validate:"omitempty,oneof=createdAt title category relevance"`
	SortOrder *string `json:"sortOrder,omitempty" validate:"omitempty,oneof=asc desc"`
}

func (f EventFilter) Validate() error {
	if f.DateFrom != nil && f.DateTo != nil && f.DateFrom.After(*f.DateTo) {
		return fmt.Errorf("dateFrom cannot be after dateTo")
	}

	if f.SortBy != nil && *f.SortBy == SortByRelevance && f.SearchText == nil {
		return fmt.Errorf("relevance sorting requires searchText")
	}

	return nil
}

func (f *EventFilter) ApplyDefaults() {
	if f.SortBy == nil {
		defaultSortBy := SortByCreatedAt
		f.SortBy = &defaultSortBy
	}

	if f.SortOrder == nil {
		defaultSortOrder := SortOrderDesc
		f.SortOrder = &defaultSortOrder
	}
}
