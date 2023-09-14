package todo

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID
	Description string
	Completed   bool
	CreatedAt   time.Time
}

func New(description string) *Item {
	return &Item{
		ID:          uuid.New(),
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
}

// Update an item
func (item *Item) Update(completed bool, description string) {
	item.Completed = completed
	item.Description = description
}
