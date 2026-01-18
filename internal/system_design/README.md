# 🏗️ System Design Primitives

> **Building blocks for scalable backend systems**

This package implements fundamental system design patterns that appear in every large-scale backend system. Master these to excel in system design interviews.

---

## 📖 Topics

| Topic | File | Key Concepts |
|:------|:-----|:-------------|
| **Rate Limiter** | [rate_limiter.go](rate_limiter.go) | Token bucket, sliding window, distributed rate limiting |
| **Cache** | [cache.go](cache.go) | LRU eviction, TTL, cache invalidation, write strategies |
| **Pub-Sub** | [pubsub.go](pubsub.go) | Message broker, topic-based routing, fan-out |
| **Idempotency** | [idempotency.go](idempotency.go) | Idempotent operations, request deduplication, idempotency keys |
| **Pagination** | [pagination.go](pagination.go) | Cursor-based, offset-based, keyset pagination |
| **Orchestrator** | [orchestrator.go](orchestrator.go) | Workflow coordination, saga pattern, compensation |

---

## 🚀 Quick Start

```bash
# Run all tests
go test -v ./system_design/

# Run specific primitive
go test -v ./system_design/ -run TestRateLimiter

# Run with race detector
go test -race ./system_design/
```

---

## 🎓 Learning Guide

### Recommended Order
1. **Cache** - Most fundamental primitive
2. **Rate Limiter** - Protect your APIs
3. **Pagination** - Handle large datasets
4. **Idempotency** - Ensure correctness
5. **Pub-Sub** - Decouple components
6. **Orchestrator** - Coordinate workflows

### Common Interview Questions

**Rate Limiter:**
- "Design a rate limiter for an API"
- "Explain token bucket vs leaky bucket"
- "How do you implement distributed rate limiting?"

**Cache:**
- "Design a cache with TTL support"
- "Explain cache eviction policies (LRU, LFU, FIFO)"
- "What is cache stampede and how do you prevent it?"

**Pub-Sub:**
- "Design a message queue"
- "How do you ensure message delivery?"
- "Explain at-most-once vs at-least-once vs exactly-once delivery"

**Idempotency:**
- "How do you make a payment API idempotent?"
- "What is an idempotency key?"
- "How long should you store idempotency keys?"

---

## 💡 Key Takeaways

### Rate Limiter (Token Bucket)
```go
// Algorithm:
// 1. Tokens refill at constant rate
// 2. Request consumes tokens
// 3. Reject if insufficient tokens

// Use cases:
// - API rate limiting
// - DDoS protection
// - Resource throttling

// Trade-offs:
// - Token bucket: Allows bursts
// - Leaky bucket: Smooth rate
// - Fixed window: Simple but has boundary issues
// - Sliding window: Accurate but complex
```

### Cache Strategies
```go
// Write Strategies:
// 1. Write-Through: Write to cache + DB synchronously
// 2. Write-Behind: Write to cache, async to DB
// 3. Write-Around: Write to DB, invalidate cache

// Read Strategies:
// 1. Cache-Aside: App checks cache, loads from DB on miss
// 2. Read-Through: Cache loads from DB automatically

// Eviction Policies:
// - LRU (Least Recently Used): Best for most use cases
// - LFU (Least Frequently Used): Good for hot data
// - FIFO: Simple but not optimal
// - TTL: Time-based expiration
```

### Pub-Sub Pattern
```go
// Components:
// - Publisher: Sends messages
// - Subscriber: Receives messages
// - Topic: Message category
// - Broker: Routes messages

// Benefits:
// - Decoupling
// - Scalability
// - Asynchronous communication

// Challenges:
// - Message ordering
// - Delivery guarantees
// - Backpressure handling
```

### Idempotency
```go
// Idempotent: Same operation, same result
// Examples:
// - PUT /users/123 (idempotent)
// - POST /users (NOT idempotent without idempotency key)
// - DELETE /users/123 (idempotent)

// Implementation:
// 1. Client generates idempotency key (UUID)
// 2. Server stores key + result
// 3. Duplicate request returns cached result

// TTL: Store keys for 24 hours (configurable)
```

### Pagination Strategies
```go
// 1. Offset-Based
// - Simple: ?page=2&limit=10
// - Problem: Inconsistent with concurrent writes
// - Use: Small datasets, admin panels

// 2. Cursor-Based
// - Stable: ?cursor=xyz&limit=10
// - Handles concurrent writes
// - Use: Infinite scroll, APIs

// 3. Keyset Pagination
// - Efficient: ?after_id=123&limit=10
// - Requires indexed column
// - Use: Large datasets, high performance
```

---

## ⚠️ Common Pitfalls

1. **Rate Limiter**
   - Not handling distributed systems (use Redis)
   - Fixed window boundary issues
   - Not considering burst traffic

2. **Cache**
   - Cache stampede (thundering herd)
   - Stale data issues
   - Memory leaks (no eviction)

3. **Pub-Sub**
   - Not handling slow consumers
   - Message loss on subscriber crash
   - Unbounded queue growth

4. **Idempotency**
   - Storing keys forever (memory leak)
   - Not handling concurrent duplicate requests
   - Using weak idempotency keys

5. **Pagination**
   - Offset-based with large offsets (slow)
   - Not handling deleted items
   - Inconsistent ordering

---

## 🏗️ System Design Applications

### Design Instagram
- **Cache**: User profiles, feed posts
- **Rate Limiter**: Upload API, like API
- **Pub-Sub**: Notification system
- **Pagination**: Feed scrolling

### Design URL Shortener
- **Cache**: Hot URLs (80/20 rule)
- **Rate Limiter**: Prevent abuse
- **Idempotency**: Duplicate URL creation

### Design Payment System
- **Idempotency**: Prevent double charges
- **Rate Limiter**: Fraud prevention
- **Orchestrator**: Multi-step payment flow

---

## 🔗 Related Topics

- **[Patterns](../patterns/)** - Circuit breaker, retry patterns
- **[Concurrency](../concurrency/)** - Thread-safe implementations
- **[Data Structures](../ds/)** - LRU cache, heap for priority queues
- **[Database Fundamentals](../../fundamentals/database/)** - Caching strategies, replication

---

[← Back to Internal](../) | [↑ Back to Main](../../README.md)
