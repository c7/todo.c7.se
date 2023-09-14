package todo

import "github.com/google/uuid"

type Repository interface {
	Add(description string) *Item
	Remove(id uuid.UUID)
	Update(id uuid.UUID, completed bool, description string) *Item
	Search(search string) List
	All() List
	Get(id uuid.UUID) *Item
	Reorder(ids []uuid.UUID) List
}
