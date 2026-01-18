package systemdesign

import (
	"testing"
)

func TestPaginateSlice(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Page 1
	res, hasMore := PaginateSlice(data, 1, 3)
	if len(res) != 3 {
		t.Errorf("expected 3 items, got %d", len(res))
	}
	if res[0] != 1 || res[2] != 3 {
		t.Error("incorrect slicing")
	}
	if !hasMore {
		t.Error("expected hasMore to be true")
	}

	// Last Page
	res, hasMore = PaginateSlice(data, 4, 3) // Start index 9: element [10]
	if len(res) != 1 {
		t.Errorf("expected 1 item, got %d", len(res))
	}
	if res[0] != 10 {
		t.Error("expected item 10")
	}
	if hasMore {
		t.Error("expected hasMore to be false at end")
	}
}

func TestPaginateByCursor(t *testing.T) {
	items := []Item{
		{ID: "a", Data: "1"},
		{ID: "b", Data: "2"},
		{ID: "c", Data: "3"},
		{ID: "d", Data: "4"},
	}

	// Fetch first page
	res, nextCursor, hasMore := PaginateByCursor(items, "", 2)
	if len(res) != 2 {
		t.Errorf("expected 2 items, got %d", len(res))
	}
	if res[0].ID != "a" || res[1].ID != "b" {
		t.Error("incorrect items")
	}
	if nextCursor != "b" {
		t.Errorf("expected cursor 'b', got %s", nextCursor)
	}
	if !hasMore {
		t.Error("expected hasMore")
	}

	// Fetch second page using cursor
	res, nextCursor, hasMore = PaginateByCursor(items, nextCursor, 2)
	if len(res) != 2 {
		t.Errorf("expected 2 items, got %d", len(res))
	}
	if res[0].ID != "c" {
		t.Error("expected start from c")
	}
	if nextCursor != "d" {
		t.Errorf("expected cursor 'd', got %s", nextCursor)
	}
	if hasMore {
		t.Error("expected no more items")
	}
}
