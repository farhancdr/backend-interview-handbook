package memory

import (
	"reflect"
	"testing"
)

func TestSliceGrowthPattern(t *testing.T) {
	capacities := SliceGrowthPattern(100)

	// Verify that capacities are growing
	if len(capacities) == 0 {
		t.Error("expected some capacity changes")
	}

	// Verify growth pattern (should generally double)
	for i := 1; i < len(capacities); i++ {
		if capacities[i] <= capacities[i-1] {
			t.Errorf("capacity should increase: %d -> %d", capacities[i-1], capacities[i])
		}
	}
}

func TestAppendWithoutPreallocation(t *testing.T) {
	n := 10
	result := AppendWithoutPreallocation(n)

	if len(result) != n {
		t.Errorf("expected length %d, got %d", n, len(result))
	}

	for i := 0; i < n; i++ {
		if result[i] != i {
			t.Errorf("expected result[%d] = %d, got %d", i, i, result[i])
		}
	}
}

func TestAppendWithPreallocation(t *testing.T) {
	n := 10
	result := AppendWithPreallocation(n)

	if len(result) != n {
		t.Errorf("expected length %d, got %d", n, len(result))
	}

	if cap(result) < n {
		t.Errorf("expected capacity >= %d, got %d", n, cap(result))
	}

	for i := 0; i < n; i++ {
		if result[i] != i {
			t.Errorf("expected result[%d] = %d, got %d", i, i, result[i])
		}
	}
}

func TestAppendWithLengthPreallocation(t *testing.T) {
	n := 10
	result := AppendWithLengthPreallocation(n)

	if len(result) != n {
		t.Errorf("expected length %d, got %d", n, len(result))
	}

	for i := 0; i < n; i++ {
		if result[i] != i {
			t.Errorf("expected result[%d] = %d, got %d", i, i, result[i])
		}
	}
}

func TestSliceCapacityInfo(t *testing.T) {
	s := make([]int, 5, 10)

	length, capacity := SliceCapacityInfo(s)

	if length != 5 {
		t.Errorf("expected length 5, got %d", length)
	}

	if capacity != 10 {
		t.Errorf("expected capacity 10, got %d", capacity)
	}
}

func TestDemonstrateSliceSharing(t *testing.T) {
	original, slice1, slice2, modified := DemonstrateSliceSharing()

	if !modified {
		t.Error("expected slices to share underlying array")
	}

	if original[2] != 99 {
		t.Errorf("expected original[2] = 99, got %d", original[2])
	}

	if slice1[2] != 99 {
		t.Errorf("expected slice1[2] = 99, got %d", slice1[2])
	}

	if slice2[0] != 99 {
		t.Errorf("expected slice2[0] = 99, got %d", slice2[0])
	}
}

func TestAppendCausingReallocation(t *testing.T) {
	original, appended, sameArray := AppendCausingReallocation()

	if sameArray {
		t.Error("expected different arrays after reallocation")
	}

	if original[0] != 1 {
		t.Errorf("expected original[0] = 1, got %d", original[0])
	}

	if appended[0] != 99 {
		t.Errorf("expected appended[0] = 99, got %d", appended[0])
	}

	if len(appended) != 4 {
		t.Errorf("expected appended length 4, got %d", len(appended))
	}
}

func TestCopyVsAppend(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}

	copied, appended := CopyVsAppend(src)

	if !reflect.DeepEqual(copied, src) {
		t.Errorf("copied should equal src: %v vs %v", copied, src)
	}

	if !reflect.DeepEqual(appended, src) {
		t.Errorf("appended should equal src: %v vs %v", appended, src)
	}

	// Modify copied and verify independence
	copied[0] = 99
	if src[0] == 99 {
		t.Error("copied should be independent of src")
	}
}

func TestSliceCapacityAfterSlicing(t *testing.T) {
	results := SliceCapacityAfterSlicing()

	if results["original_len"] != 5 {
		t.Errorf("expected original_len 5, got %d", results["original_len"])
	}

	if results["original_cap"] != 10 {
		t.Errorf("expected original_cap 10, got %d", results["original_cap"])
	}

	if results["sliced_len"] != 2 {
		t.Errorf("expected sliced_len 2, got %d", results["sliced_len"])
	}

	// Capacity should be from sliced position to end (10 - 1 = 9)
	if results["sliced_cap"] != 9 {
		t.Errorf("expected sliced_cap 9, got %d", results["sliced_cap"])
	}
}

func TestFullSliceExpression(t *testing.T) {
	normal, limited, normalCap, limitedCap := FullSliceExpression()

	if len(normal) != 2 {
		t.Errorf("expected normal length 2, got %d", len(normal))
	}

	if len(limited) != 2 {
		t.Errorf("expected limited length 2, got %d", len(limited))
	}

	if normalCap != 9 {
		t.Errorf("expected normal capacity 9, got %d", normalCap)
	}

	if limitedCap != 2 {
		t.Errorf("expected limited capacity 2, got %d", limitedCap)
	}

	// Verify values
	expected := []int{1, 2}
	if !reflect.DeepEqual(normal, expected) {
		t.Errorf("expected normal %v, got %v", expected, normal)
	}

	if !reflect.DeepEqual(limited, expected) {
		t.Errorf("expected limited %v, got %v", expected, limited)
	}
}

func TestNilVsEmptySlice(t *testing.T) {
	nilSlice, emptySlice, nilIsNil, emptyIsNil := NilVsEmptySlice()

	if !nilIsNil {
		t.Error("nil slice should be nil")
	}

	if emptyIsNil {
		t.Error("empty slice should not be nil")
	}

	// Both should have length 0
	if len(nilSlice) != 0 {
		t.Errorf("expected nil slice length 0, got %d", len(nilSlice))
	}

	if len(emptySlice) != 0 {
		t.Errorf("expected empty slice length 0, got %d", len(emptySlice))
	}

	// Both can be appended to
	nilSlice = append(nilSlice, 1)
	emptySlice = append(emptySlice, 1)

	if len(nilSlice) != 1 || len(emptySlice) != 1 {
		t.Error("both slices should be appendable")
	}
}

func TestPrintSliceInfo(t *testing.T) {
	s := []int{1, 2, 3}
	info := PrintSliceInfo("test", s)

	if info == "" {
		t.Error("expected non-empty info string")
	}

	// Should contain key information
	if len(info) < 10 {
		t.Error("info string seems too short")
	}
}

// Benchmarks to demonstrate performance differences

func BenchmarkAppendWithoutPreallocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendWithoutPreallocation(1000)
	}
}

func BenchmarkAppendWithPreallocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendWithPreallocation(1000)
	}
}

func BenchmarkAppendWithLengthPreallocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendWithLengthPreallocation(1000)
	}
}

func BenchmarkSliceGrowth_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int
		for j := 0; j < 100; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkSliceGrowth_Preallocated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 100)
		for j := 0; j < 100; j++ {
			s = append(s, j)
		}
	}
}
