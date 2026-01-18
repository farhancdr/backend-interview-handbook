# 📚 Backend Fundamentals

> **Theory that Powers Practice**
>
> Deep dive into Operating Systems, Networking, and Database internals that every backend engineer must know.

This section covers the theoretical foundations that underpin all backend systems. While the [`internal/`](../internal) directory contains executable Go code, these fundamentals provide the conceptual knowledge that interviewers expect you to articulate clearly.

---

## 🗂️ Topics

### 🖥️ [Operating Systems](os/README.md)

Understanding how operating systems work is crucial for writing efficient backend code.

**8 Essential Topics:**
- Process vs Thread
- Concurrency Models
- Virtual Memory
- CPU Scheduling
- Deadlock
- Memory Management
- System Calls
- File Systems

**[→ Browse OS Topics](os/README.md)**

---

### 🌐 [Networking](networking/README.md)

Master the protocols and patterns that enable distributed systems.

**8 Essential Topics:**
- TCP, UDP & HTTP
- TLS/SSL
- Load Balancing (L4/L7)
- DNS Deep Dive
- Real-time Communication (WebSockets, SSE)
- Modern Protocols (HTTP/2, HTTP/3, gRPC, GraphQL)
- CDN & Edge Computing
- API Gateway

**[→ Browse Networking Topics](networking/README.md)**

---

### 🗄️ [Database](database/README.md)

Deep understanding of database internals and distributed data systems.

**9 Essential Topics:**
- ACID & Isolation Levels
- Indexing & Partitioning (B-Tree vs LSM)
- Replication
- Locking Mechanisms
- Query Optimization
- NoSQL vs SQL
- Connection Pooling
- Caching Strategies
- Distributed Transactions

**[→ Browse Database Topics](database/README.md)**

---

## 🎯 How to Use This Section

1. **Start with fundamentals** - Read the markdown files to understand the theory
2. **Connect to practice** - See how these concepts are implemented in [`internal/`](../internal)
3. **Interview prep** - Use these as talking points for system design discussions

### Example Learning Path

```
1. Read: fundamentals/networking/tcp_udp_http.md
2. Code: internal/concurrency/tcp_server.go (if exists)
3. Practice: Explain the 3-way handshake in your own words
```

---

## 🔗 Related Sections

- **[Go Implementations](../internal/)** - See these concepts in executable Go code
- **[System Design Patterns](../internal/system_design/)** - Apply these fundamentals to real systems
- **[Main README](../README.md)** - Back to project overview

---

**💡 Pro Tip:** Interviewers often ask "Why?" questions. These fundamentals help you explain the reasoning behind architectural decisions.
