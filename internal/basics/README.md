# 🎯 Go Basics

> **Master the fundamentals before moving to advanced topics**

This package covers the essential Go concepts that every backend engineer must know. These are the building blocks for everything else.

---

## 📖 Topics

| Topic | File | Key Concepts |
|:------|:-----|:-------------|
| **Arrays & Slices** | [arrays_slices.go](arrays_slices.go) | Fixed vs dynamic arrays, slice internals, capacity, append behavior |
| **Maps** | [maps.go](maps.go) | Hash maps, zero values, iteration order, concurrent access |
| **Structs** | [structs.go](structs.go) | Struct composition, embedding, tags, memory layout |
| **Interfaces** | [interfaces.go](interfaces.go) | Interface satisfaction, empty interface, type assertions, type switches |
| **Value vs Reference** | [value_vs_reference.go](value_vs_reference.go) | Pass by value, pointers, when to use pointers, method receivers |
| **Errors** | [errors.go](errors.go) | Error handling, custom errors, error wrapping, sentinel errors |

---

## 🚀 Quick Start

```bash
# Run all tests
go test -v ./basics/

# Run specific test
go test -v ./basics/ -run TestSliceAppend

# See test coverage
go test -cover ./basics/
```

---

## 🎓 Learning Guide

### Start Here
1. **Value vs Reference** - Understand Go's memory model first
2. **Arrays & Slices** - Most commonly used data structure
3. **Maps** - Hash table fundamentals
4. **Structs** - Building custom types
5. **Interfaces** - Go's polymorphism mechanism
6. **Errors** - Idiomatic error handling

### Common Interview Questions

**Slices:**
- "What's the difference between `len()` and `cap()`?"
- "What happens when you append to a slice beyond its capacity?"
- "How do you create a slice with pre-allocated capacity?"

**Interfaces:**
- "What is the empty interface `interface{}`?"
- "How do you check if a type implements an interface?"
- "Explain type assertions vs type switches"

**Pointers:**
- "When should you use pointer receivers vs value receivers?"
- "What's the difference between `new()` and `make()`?"

---

## 💡 Key Takeaways

### Slices
- Slices are **references** to underlying arrays
- Capacity doubles when append exceeds current capacity
- Pre-allocate with `make([]T, 0, capacity)` for performance

### Maps
- Maps are **not safe** for concurrent access
- Iteration order is **random**
- Zero value is `nil`, must use `make()` to initialize

### Interfaces
- Interfaces are satisfied **implicitly**
- Empty interface `interface{}` can hold any type
- Use type assertions to extract concrete types

### Pointers
- Go is **pass-by-value** (even for pointers!)
- Use pointer receivers for: large structs, mutation, consistency
- Use value receivers for: small structs, immutability

---

## 🔗 Next Steps

After mastering basics, move to:
- **[Intermediate](../intermediate/)** - Defer, panic/recover, error wrapping
- **[Advanced](../advanced/)** - Context, generics, reflection
- **[Concurrency](../concurrency/)** - Goroutines and channels

---

[← Back to Internal](../) | [↑ Back to Main](../../README.md)
