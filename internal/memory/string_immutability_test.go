package memory

import (
	"strings"
	"testing"
)

func TestStringsAreImmutable(t *testing.T) {
	original, modified, same := StringsAreImmutable()

	if !same {
		t.Error("original string should not be modified")
	}

	if original != "hello" {
		t.Errorf("expected 'hello', got %s", original)
	}

	if modified != "hello world" {
		t.Errorf("expected 'hello world', got %s", modified)
	}
}

func TestStringConcatenationInLoop(t *testing.T) {
	result := StringConcatenationInLoop(10)

	if len(result) != 10 {
		t.Errorf("expected length 10, got %d", len(result))
	}

	if result != "aaaaaaaaaa" {
		t.Errorf("expected 'aaaaaaaaaa', got %s", result)
	}
}

func TestStringBuilderPattern(t *testing.T) {
	result := StringBuilderPattern(10)

	if len(result) != 10 {
		t.Errorf("expected length 10, got %d", len(result))
	}

	if result != "aaaaaaaaaa" {
		t.Errorf("expected 'aaaaaaaaaa', got %s", result)
	}
}

func TestByteSlicePattern(t *testing.T) {
	result := ByteSlicePattern(10)

	if len(result) != 10 {
		t.Errorf("expected length 10, got %d", len(result))
	}

	if result != "aaaaaaaaaa" {
		t.Errorf("expected 'aaaaaaaaaa', got %s", result)
	}
}

func TestStringToByteSliceConversion(t *testing.T) {
	s := "hello"
	b := StringToByteSliceConversion(s)

	if string(b) != s {
		t.Errorf("expected %s, got %s", s, string(b))
	}

	// Modify byte slice
	b[0] = 'H'

	// Original string should be unchanged
	if s != "hello" {
		t.Error("original string should not be modified")
	}
}

func TestByteSliceToStringConversion(t *testing.T) {
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	s := ByteSliceToStringConversion(b)

	if s != "hello" {
		t.Errorf("expected 'hello', got %s", s)
	}

	// Modify byte slice
	b[0] = 'H'

	// String should be unchanged
	if s != "hello" {
		t.Error("string should not be modified")
	}
}

func TestStringSharesMemory(t *testing.T) {
	original, substring, sharesMemory := StringSharesMemory()

	if !sharesMemory {
		t.Error("substring should share memory with original")
	}

	if original != "hello world" {
		t.Errorf("expected 'hello world', got %s", original)
	}

	if substring != "hello" {
		t.Errorf("expected 'hello', got %s", substring)
	}
}

func TestUnsafeStringToBytes(t *testing.T) {
	s := "hello"
	b := UnsafeStringToBytes(s)

	if string(b) != s {
		t.Errorf("expected %s, got %s", s, string(b))
	}

	// WARNING: Modifying b would modify the string (undefined behavior)
	// Don't do: b[0] = 'H'
}

func TestUnsafeBytesToString(t *testing.T) {
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	s := UnsafeBytesToString(b)

	if s != "hello" {
		t.Errorf("expected 'hello', got %s", s)
	}

	// WARNING: Modifying b after conversion is dangerous
}

func TestStringComparisonCost(t *testing.T) {
	s1 := "hello"
	s2 := "hello"
	s3 := "world"

	if !StringComparisonCost(s1, s2) {
		t.Error("s1 and s2 should be equal")
	}

	if StringComparisonCost(s1, s3) {
		t.Error("s1 and s3 should not be equal")
	}
}

func TestStringLengthIsConstant(t *testing.T) {
	s := "hello world"
	length := StringLengthIsConstant(s)

	if length != 11 {
		t.Errorf("expected length 11, got %d", length)
	}
}

func TestRuneVsByteIteration(t *testing.T) {
	s := "hello"
	byteCount, runeCount := RuneVsByteIteration(s)

	if byteCount != 5 {
		t.Errorf("expected byte count 5, got %d", byteCount)
	}

	if runeCount != 5 {
		t.Errorf("expected rune count 5, got %d", runeCount)
	}

	// Test with multi-byte characters
	s = "hello 世界"
	byteCount, runeCount = RuneVsByteIteration(s)

	if byteCount != 12 { // "hello " (6) + "世界" (6 bytes for 2 runes)
		t.Errorf("expected byte count 12, got %d", byteCount)
	}

	if runeCount != 8 { // 6 ASCII + 2 Chinese characters
		t.Errorf("expected rune count 8, got %d", runeCount)
	}
}

func TestStringInterning(t *testing.T) {
	s1, s2, samePointer := StringInterning()

	if s1 != s2 {
		t.Error("s1 and s2 should be equal")
	}

	// Note: samePointer is implementation-dependent
	_ = samePointer
}

func TestEfficientStringJoin(t *testing.T) {
	parts := []string{"a", "b", "c", "d"}
	result := EfficientStringJoin(parts)

	if result != "a,b,c,d" {
		t.Errorf("expected 'a,b,c,d', got %s", result)
	}
}

func TestInefficientStringJoin(t *testing.T) {
	parts := []string{"a", "b", "c", "d"}
	result := InefficientStringJoin(parts)

	if result != "a,b,c,d" {
		t.Errorf("expected 'a,b,c,d', got %s", result)
	}
}

func TestStringBuilderReuse(t *testing.T) {
	first, second := StringBuilderReuse()

	if first != "first" {
		t.Errorf("expected 'first', got %s", first)
	}

	if second != "second" {
		t.Errorf("expected 'second', got %s", second)
	}
}

func TestMultiByteCharacters(t *testing.T) {
	s := "hello"
	byteLen, runeLen := MultiByteCharacters(s)

	if byteLen != 5 || runeLen != 5 {
		t.Errorf("expected both 5, got byte=%d rune=%d", byteLen, runeLen)
	}

	s = "世界"
	byteLen, runeLen = MultiByteCharacters(s)

	if byteLen != 6 {
		t.Errorf("expected byte length 6, got %d", byteLen)
	}

	if runeLen != 2 {
		t.Errorf("expected rune length 2, got %d", runeLen)
	}
}

// Benchmarks

func BenchmarkStringConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringConcatenationInLoop(100)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuilderPattern(100)
	}
}

func BenchmarkByteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByteSlicePattern(100)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	s := strings.Repeat("a", 1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = StringToByteSliceConversion(s)
	}
}

func BenchmarkBytesToString(b *testing.B) {
	bytes := []byte(strings.Repeat("a", 1000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = ByteSliceToStringConversion(bytes)
	}
}

func BenchmarkUnsafeStringToBytes(b *testing.B) {
	s := strings.Repeat("a", 1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UnsafeStringToBytes(s)
	}
}

func BenchmarkEfficientJoin(b *testing.B) {
	parts := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		EfficientStringJoin(parts)
	}
}

func BenchmarkInefficientJoin(b *testing.B) {
	parts := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		InefficientStringJoin(parts)
	}
}
