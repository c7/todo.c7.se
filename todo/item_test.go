package todo

import "testing"

func TestItemUpdate(t *testing.T) {
	item := New("test description")

	if got, want := item.Description, "test description"; got != want {
		t.Fatalf("item.Description = %q, want %q", got, want)
	}

	if got, want := item.Completed, false; got != want {
		t.Fatalf("item.Completed = %v, want %v", got, want)
	}

	item.Update(true, "another description")

	if got, want := item.Description, "another description"; got != want {
		t.Fatalf("item.Description = %q, want %q", got, want)
	}

	if got, want := item.Completed, true; got != want {
		t.Fatalf("item.Completed = %v, want %v", got, want)
	}

	item.Update(false, "another description")

	if got, want := item.Completed, false; got != want {
		t.Fatalf("item.Completed = %v, want %v", got, want)
	}
}
