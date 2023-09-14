package todo

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	// Add an item to the list
	Add(ctx context.Context, description string) (*Item, error)
	// Remove an item from the list
	Remove(ctx context.Context, id uuid.UUID) error
	// Update an item in the list
	Update(ctx context.Context, id uuid.UUID, completed bool, description string) (*Item, error)
	// Search returns a list of items that match the search string
	Search(ctx context.Context, search string) (List, error)
	// Get returns an item by id
	Get(ctx context.Context, id uuid.UUID) (*Item, error)
	// Sort sorts the items by the given ids
	Sort(ctx context.Context, ids []uuid.UUID) error
	// List returns a copy of the list
	List(ctx context.Context) (List, error)
}
