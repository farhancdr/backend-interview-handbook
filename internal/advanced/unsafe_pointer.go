package advanced

import (
	"unsafe"
)

// ZeroCopyBytesToString converts a byte slice to a string without allocation.
// ⚠️ WARNING: The byte slice must not be modified while the string is in use.
// If the byte slice is modified, the string will change (which violates string immutability semantics).
func ZeroCopyBytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ZeroCopyStringToBytes converts a string to a byte slice without allocation.
// ⚠️ WARNING: The byte slice must NOT be modified. Strings are immutable in Go.
// Modifying this slice can cause panic or undefined behavior.
func ZeroCopyStringToBytes(s string) []byte {
	// In Go 1.20+, unsafe.Slice(unsafe.StringData(s), len(s)) is preferred.
	// This is the classic way for older Go versions or deep understanding.

	// StringHeader and SliceHeader are deprecated in recent Go but useful for understanding the layout.
	// We will use the modern unsafe ptr conversion way which is safer and standard now.

	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// Note: For older Go versions (pre 1.20), one might use reflect.StringHeader and reflect.SliceHeader
// but that is fraught with danger regarding GC. The above method using unsafe.StringData is safe-ish
// as long as you respect the "do not mutate" rule.
