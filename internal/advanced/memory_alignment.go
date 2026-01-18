package advanced

import "unsafe"

// BadStruct has fields ordered inefficiently, causing padding.
// bool (1 byte) + padding (7 bytes) + int64 (8 bytes) + bool (1 byte) + padding (7 bytes)
// Total size = 24 bytes (on 64-bit systems)
type BadStruct struct {
	Flag1 bool
	Value int64
	Flag2 bool
}

// GoodStruct has fields ordered efficiently.
// int64 (8 bytes) + bool (1 byte) + bool (1 byte) + padding (6 bytes)
// Total size = 16 bytes (on 64-bit systems)
type GoodStruct struct {
	Value int64
	Flag1 bool
	Flag2 bool
}

// GetStructSizes returns the size in bytes of bad and good structs.
func GetStructSizes() (bad, good uintptr) {
	return unsafe.Sizeof(BadStruct{}), unsafe.Sizeof(GoodStruct{})
}
