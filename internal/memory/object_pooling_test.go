package memory

import (
	"strings"
	"testing"
)

func TestWithoutPool(t *testing.T) {
	n := 10
	results := WithoutPool(n)

	if len(results) != n {
		t.Errorf("expected %d results, got %d", n, len(results))
	}

	for i, buf := range results {
		if buf[0] != byte(i) {
			t.Errorf("expected buf[0] = %d, got %d", i, buf[0])
		}
	}
}

func TestWithPool(t *testing.T) {
	n := 10
	results := WithPool(n)

	if len(results) != n {
		t.Errorf("expected %d results, got %d", n, len(results))
	}
}

func TestBufferPool(t *testing.T) {
	buf := GetBuffer()

	if buf == nil {
		t.Error("buffer should not be nil")
	}

	buf.WriteString("test")

	if buf.String() != "test" {
		t.Errorf("expected 'test', got %s", buf.String())
	}

	PutBuffer(buf)

	// Get another buffer (might be the same one, reset)
	buf2 := GetBuffer()

	if buf2.Len() != 0 {
		t.Error("buffer should be reset")
	}

	PutBuffer(buf2)
}

func TestUseBufferPool(t *testing.T) {
	result := UseBufferPool("data")

	if result != "data processed" {
		t.Errorf("expected 'data processed', got %s", result)
	}
}

func TestStructPool(t *testing.T) {
	s := GetStruct()

	if s == nil {
		t.Error("struct should not be nil")
	}

	s.Count = 42
	s.Name = "test"

	PutStruct(s)

	// Get another struct (might be the same one, reset)
	s2 := GetStruct()

	if s2.Count != 0 || s2.Name != "" {
		t.Error("struct should be reset")
	}

	PutStruct(s2)
}

func TestProcessWithStructPool(t *testing.T) {
	result := ProcessWithStructPool(10, "test")

	if result != 20 {
		t.Errorf("expected 20, got %d", result)
	}
}

func TestCustomObjectPool(t *testing.T) {
	obj := GetCustomObject()

	if obj == nil {
		t.Error("object should not be nil")
	}

	obj.items = append(obj.items, 1, 2, 3)
	obj.name = "test"

	if len(obj.items) != 3 {
		t.Errorf("expected 3 items, got %d", len(obj.items))
	}

	PutCustomObject(obj)

	// Get another object (might be the same one, reset)
	obj2 := GetCustomObject()

	if len(obj2.items) != 0 || obj2.name != "" {
		t.Error("object should be reset")
	}

	PutCustomObject(obj2)
}

func TestWhenToUsePool(t *testing.T) {
	// Large, frequent allocations - should use pool
	if !WhenToUsePool(2048, 10000) {
		t.Error("should recommend pool for large, frequent allocations")
	}

	// Small, infrequent allocations - should not use pool
	if WhenToUsePool(100, 10) {
		t.Error("should not recommend pool for small, infrequent allocations")
	}
}

func TestPoolLifecycle(t *testing.T) {
	// Just verify it doesn't panic
	PoolLifecycle()
}

func TestConcurrentPoolUsage(t *testing.T) {
	// Verify no race conditions
	ConcurrentPoolUsage(100)
}

func TestBufferPoolReset(t *testing.T) {
	buf := GetBuffer()
	buf.WriteString("first use")

	firstContent := buf.String()
	if firstContent != "first use" {
		t.Errorf("expected 'first use', got %s", firstContent)
	}

	PutBuffer(buf)

	// Get buffer again
	buf2 := GetBuffer()

	// Should be empty (reset)
	if buf2.Len() != 0 {
		t.Errorf("expected empty buffer, got length %d", buf2.Len())
	}

	buf2.WriteString("second use")
	PutBuffer(buf2)
}

// Benchmarks to demonstrate pool benefits

func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithoutPool(100)
	}
}

func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithPool(100)
	}
}

func BenchmarkBufferWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		buf.WriteString("test data")
		buf.WriteString(" more data")
		_ = buf.String()
	}
}

func BenchmarkBufferWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := GetBuffer()
		buf.WriteString("test data")
		buf.WriteString(" more data")
		_ = buf.String()
		PutBuffer(buf)
	}
}

func BenchmarkStructWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := &LargeStruct{}
		s.Count = i
		s.Name = "test"
		_ = s.Count * 2
	}
}

func BenchmarkStructWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := GetStruct()
		s.Count = i
		s.Name = "test"
		_ = s.Count * 2
		PutStruct(s)
	}
}

func BenchmarkConcurrentPoolAccess(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := GetBuffer()
			buf.WriteString("concurrent")
			PutBuffer(buf)
		}
	})
}
