package basics

import "testing"

func TestValueType_Copy(t *testing.T) {
	original := ValueType{Value: 42}
	copy := original

	// Modify the copy
	copy.Value = 100

	// Original should be unchanged
	if original.Value != 42 {
		t.Errorf("expected original.Value to be 42, got %d", original.Value)
	}

	// Copy should be modified
	if copy.Value != 100 {
		t.Errorf("expected copy.Value to be 100, ?got %d", copy.Value)
	}
}

func TestValueType_ModifyByValue(t *testing.T) {
	v := ValueType{Value: 42}

	// Pass by value - won't modify original
	ModifyValue(v)

	// Original should be unchanged
	if v.Value != 42 {
		t.Errorf("expected v.Value to be 42, got %d", v.Value)
	}
}

func TestValueType_ModifyByPointer(t *testing.T) {
	v := ValueType{Value: 42}

	// Pass by pointer - will modify original
	ModifyValuePointer(&v)

	// Original should be modified
	if v.Value != 100 {
		t.Errorf("expected v.Value to be 100, got %d", v.Value)
	}
}

func TestSlice_ReferenceSemantics(t *testing.T) {
	original := []int{1, 2, 3}
	reference := original

	// Modify via reference
	reference[0] = 999

	// Both should show the modification
	if original[0] != 999 {
		t.Errorf("expected original[0] to be 999, got %d", original[0])
	}
	if reference[0] != 999 {
		t.Errorf("expected reference[0] to be 999, got %d", reference[0])
	}
}

func TestSlice_ModifyInFunction(t *testing.T) {
	s := []int{1, 2, 3}

	// Pass slice to function
	ModifySlice(s)

	// Original slice should be modified
	if s[0] != 999 {
		t.Errorf("expected s[0] to be 999, got %d", s[0])
	}
}

func TestSliceAndMap_ReferenceSemantics(t *testing.T) {
	original := map[string]int{"key": 42}
	reference := original

	// Modify via reference
	reference["key"] = 999

	// Both should show the modification
	if original["key"] != 999 {
		t.Errorf("expected original[key] to be 999, got %d", original["key"])
	}
	if reference["key"] != 999 {
		t.Errorf("expected reference[key] to be 999, got %d", reference["key"])
	}
}

func TestMap_ModifyInFunction(t *testing.T) {
	m := map[string]int{"key": 42}

	// Pass map to function
	ModifyMap(m)

	// Original map should be modified
	if m["key"] != 999 {
		t.Errorf("expected m[key] to be 999, got %d", m["key"])
	}
}

func TestDemonstrateValueCopy(t *testing.T) {
	original, modified := DemonstrateValueCopy()

	// Original should be unchanged
	if original.Value != 42 {
		t.Errorf("expected original.Value to be 42, got %d", original.Value)
	}

	// Modified should be changed
	if modified.Value != 100 {
		t.Errorf("expected modified.Value to be 100, got %d", modified.Value)
	}
}

func TestDemonstrateSliceReference(t *testing.T) {
	original, reference := DemonstrateSliceReference()

	// Both should show the same modification
	if original[0] != 999 {
		t.Errorf("expected original[0] to be 999, got %d", original[0])
	}
	if reference[0] != 999 {
		t.Errorf("expected reference[0] to be 999, got %d", reference[0])
	}
}

func TestDemonstrateMapReference(t *testing.T) {
	original, reference := DemonstrateMapReference()

	// Both should show the same modification
	if original["a"] != 999 {
		t.Errorf("expected original[a] to be 999, got %d", original["a"])
	}
	if reference["a"] != 999 {
		t.Errorf("expected reference[a] to be 999, got %d", reference["a"])
	}
}

func TestArray_ValueSemantics(t *testing.T) {
	// Arrays are value types (unlike slices)
	original := [3]int{1, 2, 3}
	copy := original

	// Modify the copy
	copy[0] = 999

	// Original should be unchanged
	if original[0] != 1 {
		t.Errorf("expected original[0] to be 1, got %d", original[0])
	}

	// Copy should be modified
	if copy[0] != 999 {
		t.Errorf("expected copy[0] to be 999, got %d", copy[0])
	}
}
