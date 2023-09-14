package todo

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	// Add adds a todo to the list
	Add(ctx context.Context, description string) (*Item, error)
	// Remove removes a todo from the list
	Remove(ctx context.Context, id uuid.UUID) error
	// Update updates a todo in the list
	Update(ctx context.Context, id uuid.UUID, completed bool, description string) (*Item, error)
	// Search returns a list of todos that match the search string
	Search(ctx context.Context, search string) (List, error)
	// Get returns a todo by id
	Get(ctx context.Context, id uuid.UUID) (*Item, error)
	// Sort sorts the todos by the given ids
	Sort(ctx context.Context, ids []uuid.UUID) error
	// List returns a copy of the todos list
	List(ctx context.Context) (List, error)
}
