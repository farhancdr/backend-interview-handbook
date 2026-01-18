# 🚀 Advanced Go

> **Power user features for experienced Go developers**

This package covers advanced Go features that demonstrate deep language knowledge. These topics frequently appear in senior-level interviews.

---

## 📖 Topics

| Topic | File | Key Concepts |
|:------|:-----|:-------------|
| **Context** | [context.go](context.go) | Cancellation, timeouts, deadlines, context values, propagation |
| **Generics** | [generics.go](generics.go) | Type parameters, constraints, generic functions, generic types |
| **Reflection** | [reflection.go](reflection.go) | Type introspection, dynamic method calls, struct tags, reflect package |
| **Unsafe Pointers** | [unsafe_pointer.go](unsafe_pointer.go) | Unsafe operations, pointer arithmetic, memory manipulation |
| **Memory Alignment** | [memory_alignment.go](memory_alignment.go) | Struct padding, alignment rules, memory optimization |
| **Functional Options** | [functional_options.go](functional_options.go) | Builder pattern, optional parameters, API design |

---

## 🚀 Quick Start

```bash
# Run all tests
go test -v ./advanced/

# Run specific test
go test -v ./advanced/ -run TestContext

# Check for race conditions
go test -race ./advanced/
```

---

## 🎓 Learning Guide

### Recommended Order
1. **Context** - Essential for production Go code
2. **Functional Options** - Clean API design pattern
3. **Generics** - Modern Go (1.18+) feature
4. **Memory Alignment** - Performance optimization
5. **Reflection** - Powerful but use sparingly
6. **Unsafe Pointers** - Understand but avoid in production

### Common Interview Questions

**Context:**
- "When should you use `context.WithTimeout` vs `context.WithDeadline`?"
- "How do you pass request-scoped values through context?"
- "What happens if you don't cancel a context?"

**Generics:**
- "How do you constrain generic types?"
- "What's the difference between `any` and `interface{}`?"
- "When should you use generics vs interfaces?"

**Reflection:**
- "What are the performance implications of reflection?"
- "How do you read struct tags?"
- "When is reflection appropriate?"

---

## 💡 Key Takeaways

### Context
- **Always** propagate context through your call chain
- Use `context.WithTimeout` for operations with time limits
- **Never** store context in structs
- Cancel contexts to free resources

### Generics (Go 1.18+)
- Use generics for **type-safe** data structures and algorithms
- Constraints define what operations are allowed on type parameters
- `any` is an alias for `interface{}`
- Generics have **zero runtime overhead** (compile-time only)

### Reflection
- **Powerful** but **slow** - use only when necessary
- Common use cases: JSON marshaling, ORM, dependency injection
- Can bypass type safety - use with caution
- `reflect.ValueOf()` and `reflect.TypeOf()` are your entry points

### Unsafe Pointers
- Breaks Go's type safety and memory safety guarantees
- Used in low-level libraries (syscalls, cgo, performance-critical code)
- **Avoid in application code**
- Can cause crashes and security vulnerabilities

### Memory Alignment
- Go automatically aligns struct fields
- Reorder fields to minimize padding
- Use `unsafe.Sizeof()` to check struct size
- Can save significant memory in large-scale systems

### Functional Options
- Elegant way to handle optional parameters
- Makes APIs **backward compatible**
- Self-documenting code
- Used in popular libraries (gRPC, Uber's Zap)

---

## ⚠️ Common Pitfalls

1. **Context in Structs** - Don't store context in struct fields
2. **Reflection Performance** - Don't use reflection in hot paths
3. **Unsafe Misuse** - Unsafe code can corrupt memory
4. **Generic Overuse** - Don't use generics when interfaces suffice
5. **Context Values** - Don't abuse context for passing dependencies

---

## 🔗 Related Topics

- **[Basics](../basics/)** - Foundation concepts
- **[Intermediate](../intermediate/)** - Idiomatic Go patterns
- **[Concurrency](../concurrency/)** - Context is heavily used with goroutines
- **[Patterns](../patterns/)** - Functional options pattern in action

---

[← Back to Internal](../) | [↑ Back to Main](../../README.md)
