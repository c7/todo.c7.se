package todo

import (
	"testing"

	"github.com/google/uuid"
)

func TestList(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		list := NewList()

		item := list.Add("test item")

		if got, want := item.Description, "test item"; got != want {
			t.Fatalf("item.Description = %q, want %q", got, want)
		}

		list.Add("another item")

		if got, want := len(*list), 2; got != want {
			t.Fatalf("len(*list) = %d, want %d", got, want)
		}
	})

	t.Run("Remove", func(t *testing.T) {
		list := NewList()

		item := list.Add("test item")

		list.Add("another item")

		if got, want := len(*list), 2; got != want {
			t.Fatalf("len(*list) = %d, want %d", got, want)
		}

		list.Remove(item.ID)

		if got, want := len(*list), 1; got != want {
			t.Fatalf("len(*list) = %d, want %d", got, want)
		}
	})

	t.Run("Update", func(t *testing.T) {
		list := NewList()

		item := list.Add("test item")

		list.Add("another item")

		list.Update(item.ID, true, "changed item")

		item = list.Get(item.ID)

		if got, want := list.Get(item.ID).Completed, true; got != want {
			t.Fatalf("list.Get(item.ID).Completed = %v, want %v", got, want)
		}
	})

	t.Run("Search", func(t *testing.T) {
		list := NewList()

		list.Add("test item")
		list.Add("another item")

		found := list.Search("another")

		if got, want := len(found), 1; got != want {
			t.Fatalf("len(found) = %d, want %d", got, want)
		}
	})

	t.Run("All", func(t *testing.T) {
		list := NewList()

		list.Add("test item")
		list.Add("another item")

		if got, want := len(list.All()), 2; got != want {
			t.Fatalf("len(list.All()) = %d, want %d", got, want)
		}
	})

	t.Run("Get", func(t *testing.T) {
		list := NewList()

		list.Add("first item")

		second := list.Add("second item")

		if list.Get(uuid.New()) != nil {
			t.Fatalf("expected to get nil")
		}

		if list.Get(second.ID) == nil {
			t.Fatalf("got nil, expected item")
		}
	})

	t.Run("Reorder", func(t *testing.T) {
		list := NewList()

		first := list.Add("first item")

		second := list.Add("second item")

		list.Reorder([]uuid.UUID{second.ID, first.ID})

		if got, want := list.indexOf(second.ID), 0; got != want {
			t.Fatalf("list.indexOf(second.ID) = %d, want %d", got, want)
		}

		if got, want := list.indexOf(first.ID), 1; got != want {
			t.Fatalf("list.indexOf(first.ID) = %d, want %d", got, want)
		}

		if got, want := list.indexOf(uuid.New()), -1; got != want {
			t.Fatalf("list.indexOf(uuid.New()) = %d, want %d", got, want)
		}
	})
}
