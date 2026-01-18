# 🗄️ Database Fundamentals

> **Master the data layer that powers every backend system**

Understanding database internals is critical for backend interviews. This section covers everything from ACID guarantees to distributed transactions.

---

## 📖 Topics

### Core Concepts

| Topic | Description | Key Interview Questions |
|:------|:------------|:------------------------|
| **[ACID & Isolation Levels](acid_isolation.md)** | Transaction guarantees and isolation levels | "Explain the difference between Read Committed and Serializable" |
| **[Locking Mechanisms](locking_mechanisms.md)** | Pessimistic vs Optimistic locking, deadlocks | "How would you prevent deadlocks in a high-concurrency system?" |
| **[Indexing & Partitioning](indexing_partitioning.md)** | B-Tree vs LSM-Tree, sharding strategies | "When would you use a hash index vs a B-Tree index?" |

### Performance & Optimization

| Topic | Description | Key Interview Questions |
|:------|:------------|:------------------------|
| **[Query Optimization](query_optimization.md)** | Execution plans, index selection, query rewriting | "How do you optimize a slow JOIN query?" |
| **[Connection Pooling](connection_pooling.md)** | Managing database connections efficiently | "Why is connection pooling important?" |
| **[Caching Strategies](caching_strategies.md)** | Cache-aside, write-through, write-behind | "Explain cache invalidation strategies" |

### Distributed Systems

| Topic | Description | Key Interview Questions |
|:------|:------------|:------------------------|
| **[Replication](replication.md)** | Master-slave, multi-master, consensus | "What's the difference between synchronous and asynchronous replication?" |
| **[Distributed Transactions](distributed_transactions.md)** | 2PC, Saga pattern, eventual consistency | "How would you implement distributed transactions?" |
| **[NoSQL vs SQL](nosql_vs_sql.md)** | CAP theorem, consistency models, use cases | "When would you choose MongoDB over PostgreSQL?" |

---

## 🎯 Interview Preparation Guide

### Must-Know Concepts

1. **ACID Properties** - Be able to explain each property with examples
2. **Isolation Levels** - Know the trade-offs between consistency and performance
3. **Indexing** - Understand when to use different index types
4. **Replication** - Explain leader-follower vs leaderless replication
5. **CAP Theorem** - Articulate the trade-offs in distributed databases

### Common Interview Patterns

**System Design Questions:**
- "Design a URL shortener" → Requires understanding of indexing and sharding
- "Design Instagram" → Requires knowledge of replication and caching
- "Design a distributed cache" → Requires understanding of consistency models

**Behavioral Questions:**
- "Tell me about a time you optimized a slow query"
- "How would you handle a database outage?"
- "Explain a complex database migration you've done"

---

## 🔗 Related Topics

- **[System Design Patterns](../../internal/system_design/)** - See cache and pagination implementations in Go
- **[Data Structures](../../internal/ds/)** - Understand the data structures behind databases (B-Trees, Hash Maps)
- **[Networking Fundamentals](../networking/)** - Learn about protocols used in database replication

---

## 📚 Study Tips

1. **Read sequentially** - Start with ACID, then move to indexing, then distributed concepts
2. **Draw diagrams** - Visualize replication topologies and transaction flows
3. **Compare trade-offs** - Always think in terms of consistency vs availability vs performance
4. **Practice explaining** - Use the "explain it to a 5-year-old" technique

---

[← Back to Fundamentals](../) | [↑ Back to Main](../../README.md)
