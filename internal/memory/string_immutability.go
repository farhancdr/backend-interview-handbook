package memory

// Why interviewers ask this:
// String immutability is a fundamental Go concept that affects performance and memory usage.
// Interviewers want to see if you understand why strings are immutable, the cost of string
// concatenation, and when to use strings vs byte slices for better performance.

// Common pitfalls:
// - Concatenating strings in loops (creates many intermediate strings)
// - Not knowing that string-to-[]byte conversion copies data
// - Assuming string modification is possible (it's not)
// - Not using strings.Builder for efficient string building
// - Converting between string and []byte unnecessarily

// Key takeaway:
// Strings are immutable in Go. Any "modification" creates a new string. String concatenation
// in loops is O(n²) due to copying. Use strings.Builder or []byte for efficient string
// building. Converting between string and []byte copies data. Use unsafe for zero-copy
// conversion only when absolutely necessary and safe.

import (
	"strings"
	"unsafe"
)

// StringsAreImmutable demonstrates string immutability
func StringsAreImmutable() (original, modified string, same bool) {
	original = "hello"

	// This doesn't modify original, it creates a new string
	modified = original + " world"

	same = (original == "hello")

	return original, modified, same
}

// StringConcatenationInLoop shows inefficient pattern
// Each concatenation creates a new string (O(n²) complexity)
func StringConcatenationInLoop(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += "a"
	}
	return result
}

// StringBuilderPattern shows efficient string building
// strings.Builder reuses buffer (O(n) complexity)
func StringBuilderPattern(n int) string {
	var builder strings.Builder
	builder.Grow(n) // Pre-allocate capacity

	for i := 0; i < n; i++ {
		builder.WriteString("a")
	}

	return builder.String()
}

// ByteSlicePattern shows using []byte for string building
func ByteSlicePattern(n int) string {
	bytes := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		bytes = append(bytes, 'a')
	}

	return string(bytes)
}

// StringToByteSliceConversion shows conversion cost
func StringToByteSliceConversion(s string) []byte {
	// This creates a copy of the string data
	return []byte(s)
}

// ByteSliceToStringConversion shows conversion cost
func ByteSliceToStringConversion(b []byte) string {
	// This creates a copy of the byte slice data
	return string(b)
}

// StringSharesMemory demonstrates that strings can share underlying data
func StringSharesMemory() (original, substring string, sharesMemory bool) {
	original = "hello world"

	// Substring shares the underlying array (no copy)
	substring = original[0:5]

	// They share memory (implementation detail)
	sharesMemory = true

	return original, substring, sharesMemory
}

// UnsafeStringToBytes shows zero-copy conversion (DANGEROUS)
// Only use when you're absolutely sure the []byte won't be modified
func UnsafeStringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// UnsafeBytesToString shows zero-copy conversion (DANGEROUS)
// Only use when you're sure the string won't outlive the []byte
func UnsafeBytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// StringComparisonCost shows that string comparison is O(n)
func StringComparisonCost(s1, s2 string) bool {
	// String comparison checks length first, then bytes
	return s1 == s2
}

// StringLengthIsConstant shows len() is O(1)
func StringLengthIsConstant(s string) int {
	// Length is stored, not calculated
	return len(s)
}

// RuneVsByteIteration shows different iteration methods
func RuneVsByteIteration(s string) (byteCount, runeCount int) {
	// Byte iteration
	for i := 0; i < len(s); i++ {
		byteCount++
	}

	// Rune iteration (handles multi-byte UTF-8)
	for range s {
		runeCount++
	}

	return byteCount, runeCount
}

// StringInterning demonstrates string literal behavior
func StringInterning() (s1, s2 string, samePointer bool) {
	s1 = "hello"
	s2 = "hello"

	// String literals may be interned (share same memory)
	// This is an implementation detail
	samePointer = (*(*uintptr)(unsafe.Pointer(&s1)) == *(*uintptr)(unsafe.Pointer(&s2)))

	return s1, s2, samePointer
}

// EfficientStringJoin shows using strings.Join
func EfficientStringJoin(parts []string) string {
	return strings.Join(parts, ",")
}

// InefficientStringJoin shows manual concatenation
func InefficientStringJoin(parts []string) string {
	result := ""
	for i, part := range parts {
		if i > 0 {
			result += ","
		}
		result += part
	}
	return result
}

// StringBuilderReuse shows reusing a builder
func StringBuilderReuse() (first, second string) {
	var builder strings.Builder

	builder.WriteString("first")
	first = builder.String()

	builder.Reset()

	builder.WriteString("second")
	second = builder.String()

	return first, second
}

// MultiByteCharacters demonstrates UTF-8 handling
func MultiByteCharacters(s string) (byteLen, runeLen int) {
	byteLen = len(s)         // Byte length
	runeLen = len([]rune(s)) // Rune (character) length

	return byteLen, runeLen
}
