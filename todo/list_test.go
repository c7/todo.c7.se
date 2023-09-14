package todo

import "testing"

func TestList(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		list := NewList()

		item := list.Add("test item")

		if got, want := item.Description, "test item"; got != want {
			t.Fatalf("item.Description = %q, want %q", got, want)
		}
	})
}
