package systemdesign

import (
	"encoding/base64"
)

// Why interviewers ask this:
// APIs dealing with lists must pagination. Interviewers compare Offset vs Cursor strategies.
// Offset is simple but slow for deep pages (database must skip N rows).
// Cursor is faster (uses index) but inflexible (can't jump to page 10).

// Common pitfalls:
// - Off-by-one errors
// - Base64 encoding cursors without validation
// - Not handling empty results

// Key takeaway:
// Offset-based: `LIMIT 10 OFFSET 50`. Good for dashboards, bad for infinite scroll.
// Cursor-based: `WHERE id > last_seen_id LIMIT 10`. O(1) fetch time, great for scale.

// 1. Offset Pagination Logic
func PaginateSlice[T any](items []T, page, pageSize int) ([]T, bool) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	start := (page - 1) * pageSize
	if start >= len(items) {
		return []T{}, false // Out of bounds
	}

	end := start + pageSize
	if end > len(items) {
		end = len(items)
	}

	return items[start:end], end < len(items)
}

// 2. Cursor Pagination Logic (Simplified ID-based)
type Item struct {
	ID   string
	Data string
}

func PaginateByCursor(items []Item, cursor string, limit int) ([]Item, string, bool) {
	// Simple simulation: Find cursor position in sorted list
	startIndex := 0
	if cursor != "" {
		found := false
		for i, item := range items {
			if item.ID == cursor {
				startIndex = i + 1 // Start after the cursor
				found = true
				break
			}
		}
		if !found {
			// Cursor invalid or reset, start from 0 (or error depending on req)
			startIndex = 0
		}
	}

	if startIndex >= len(items) {
		return []Item{}, "", false
	}

	end := startIndex + limit
	if end > len(items) {
		end = len(items)
	}

	result := items[startIndex:end]

	// Generate next cursor
	nextCursor := ""
	hasMore := end < len(items)
	if len(result) > 0 {
		nextCursor = result[len(result)-1].ID
	}

	return result, nextCursor, hasMore
}

// EncodeCursor helper to demonstrate Opaque Cursors
func EncodeCursor(id string) string {
	return base64.StdEncoding.EncodeToString([]byte(id))
}

func DecodeCursor(cursor string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
