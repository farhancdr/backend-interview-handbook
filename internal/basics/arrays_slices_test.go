package basics

import "testing"

func TestArray_FixedSize(t *testing.T) {
	arr := ArrayExample()

	if len(arr) != 3 {
		t.Errorf("expected array length 3, got %d", len(arr))
	}

	if arr[0] != 1 || arr[1] != 2 || arr[2] != 3 {
		t.Errorf("expected [1, 2, 3], got %v", arr)
	}
}

func TestSlice_DynamicSize(t *testing.T) {
	s := SliceExample()

	if len(s) != 3 {
		t.Errorf("expected slice length 3, got %d", len(s))
	}

	// Slices can grow
	s = append(s, 4, 5)
	if len(s) != 5 {
		t.Errorf("expected slice length 5 after append, got %d", len(s))
	}
}

func TestArray_PassedByValue(t *testing.T) {
	arr := [3]int{1, 2, 3}

	// Pass array to function
	ArrayPassedByValue(arr)

	// Original should be unchanged (array was copied)
	if arr[0] != 1 {
		t.Errorf("expected arr[0] to be 1, got %d", arr[0])
	}
}

func TestSlice_PassedByReference(t *testing.T) {
	s := []int{1, 2, 3}

	// Pass slice to function
	SlicePassedByReference(s)

	// Original should be modified (slice shares backing array)
	if s[0] != 999 {
		t.Errorf("expected s[0] to be 999, got %d", s[0])
	}
}

func TestSlice_LengthVsCapacity(t *testing.T) {
	length, capacity := SliceCapacityExample()

	if length != 3 {
		t.Errorf("expected length 3, got %d", length)
	}

	if capacity != 5 {
		t.Errorf("expected capacity 5, got %d", capacity)
	}
}

func TestSlice_AppendBehavior(t *testing.T) {
	original, appended := SliceAppendExample()

	// Original should still be [1, 2, 3]
	if len(original) != 3 {
		t.Errorf("expected original length 3, got %d", len(original))
	}

	// Appended should be [1, 2, 3, 4]
	if len(appended) != 4 {
		t.Errorf("expected appended length 4, got %d", len(appended))
	}

	if appended[3] != 4 {
		t.Errorf("expected appended[3] to be 4, got %d", appended[3])
	}
}

func TestSlice_ResliceSharesBackingArray(t *testing.T) {
	original, subslice := SliceResliceExample()

	// Subslice modification should affect original
	if original[1] != 999 {
		t.Errorf("expected original[1] to be 999, got %d", original[1])
	}

	// Subslice should be [999, 3, 4]
	if subslice[0] != 999 {
		t.Errorf("expected subslice[0] to be 999, got %d", subslice[0])
	}
}

func TestArrayToSlice_Conversion(t *testing.T) {
	arr := [3]int{1, 2, 3}
	s := ArrayToSlice(arr)

	if len(s) != 3 {
		t.Errorf("expected slice length 3, got %d", len(s))
	}

	// Verify values
	for i := 0; i < 3; i++ {
		if s[i] != arr[i] {
			t.Errorf("expected s[%d] to be %d, got %d", i, arr[i], s[i])
		}
	}
}

func TestArray_Comparable(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2, 4}

	// Arrays can be compared with ==
	if !CompareArrays(a, b) {
		t.Error("expected arrays a and b to be equal")
	}

	if CompareArrays(a, c) {
		t.Error("expected arrays a and c to be different")
	}
}

func TestSlice_NotComparable(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}

	// Slices cannot be compared with ==, must use custom function
	if !CompareSlices(a, b) {
		t.Error("expected slices a and b to be equal")
	}

	if CompareSlices(a, c) {
		t.Error("expected slices a and c to be different")
	}
}

func TestNilSlice_VsEmptySlice(t *testing.T) {
	nilSlice, emptySlice, nilIsNil, emptyIsNil := NilSliceVsEmptySlice()

	// Nil slice
	if !nilIsNil {
		t.Error("expected nil slice to be nil")
	}
	if len(nilSlice) != 0 {
		t.Errorf("expected nil slice length 0, got %d", len(nilSlice))
	}

	// Empty slice
	if emptyIsNil {
		t.Error("expected empty slice to not be nil")
	}
	if len(emptySlice) != 0 {
		t.Errorf("expected empty slice length 0, got %d", len(emptySlice))
	}

	// Both have length 0, but only nil slice is nil
	if len(nilSlice) != len(emptySlice) {
		t.Error("expected both slices to have length 0")
	}
}

func TestSlice_AppendToNil(t *testing.T) {
	var s []int // nil slice

	// Appending to nil slice works fine
	s = append(s, 1, 2, 3)

	if len(s) != 3 {
		t.Errorf("expected length 3, got %d", len(s))
	}

	if s[0] != 1 || s[1] != 2 || s[2] != 3 {
		t.Errorf("expected [1, 2, 3], got %v", s)
	}
}
