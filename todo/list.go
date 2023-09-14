package todo

import (
	"strings"

	"github.com/google/uuid"
)

// List of TODO items
type List []*Item

func NewList() *List {
	return &List{}
}

func (l *List) Add(description string) *Item {
	item := New(description)

	*l = append(*l, item)

	return item
}

// Remove removes an item from the list
func (l *List) Remove(id uuid.UUID) {
	index := l.indexOf(id)
	if index == -1 {
		return
	}

	*l = append((*l)[:index], (*l)[index+1:]...)
}

// Update updates an item in the list
func (l *List) Update(id uuid.UUID, completed bool, description string) *Item {
	index := l.indexOf(id)
	if index == -1 {
		return nil
	}
	item := (*l)[index]

	item.Update(completed, description)

	return item
}

// Search returns a list of items that match the search string
func (l *List) Search(search string) List {
	list := make(List, 0)

	for _, item := range *l {
		if strings.Contains(item.Description, search) {
			list = append(list, item)
		}
	}

	return list
}

// All returns a copy of the list of items
func (l *List) All() List {
	list := make(List, len(*l))

	copy(list, *l)

	return list
}

// Get returns an item by id
func (l *List) Get(id uuid.UUID) *Item {
	index := l.indexOf(id)
	if index == -1 {
		return nil
	}

	return (*l)[index]
}

// Reorder reorders the list of items
func (l *List) Reorder(ids []uuid.UUID) List {
	newList := make(List, len(ids))

	for i, id := range ids {
		newList[i] = (*l)[l.indexOf(id)]
	}

	copy(*l, newList)

	return newList
}

// indexOf returns the index of the item with the given id or -1 if not found
func (l List) indexOf(id uuid.UUID) int {
	for i, item := range l {
		if item.ID == id {
			return i
		}
	}

	return -1
}
