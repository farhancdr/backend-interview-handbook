# NoSQL vs SQL: When to Use Each

## SQL (Relational Databases)

### Characteristics
- **Schema**: Fixed schema (tables, columns, types).
- **ACID**: Strong consistency, transactions.
- **Query Language**: SQL (Structured Query Language).
- **Relationships**: Foreign keys, joins.

### Examples
- **Postgres**, **MySQL**, **Oracle**, **SQL Server**

### Pros
- **ACID Transactions**: Strong consistency (critical for banking, e-commerce).
- **Complex Queries**: Joins, aggregations, subqueries.
- **Mature Ecosystem**: Well-understood, battle-tested.

### Cons
- **Rigid Schema**: Schema changes require migrations (ALTER TABLE).
- **Vertical Scaling**: Hard to scale horizontally (sharding is complex).
- **Performance**: Joins can be slow for large datasets.

---

## NoSQL (Non-Relational Databases)

### Types of NoSQL Databases

#### 1. Document Stores (MongoDB, CouchDB)
**Data Model**: JSON-like documents.

**Example**:
```json
{
  "_id": "123",
  "name": "Alice",
  "email": "alice@example.com",
  "posts": [
    {"title": "Hello World", "content": "..."},
    {"title": "NoSQL is cool", "content": "..."}
  ]
}
```

**Pros**:
- **Flexible Schema**: Add fields without migrations.
- **Nested Data**: Store related data together (no joins).

**Cons**:
- **No Joins**: Must denormalize data (duplicate data).
- **Eventual Consistency**: Weak consistency by default.

**Use Case**: Content management, user profiles, catalogs.

#### 2. Key-Value Stores (Redis, DynamoDB, Memcached)
**Data Model**: Key → Value (value can be string, hash, list, set).

**Example**:
```
user:123 → {"name": "Alice", "email": "alice@example.com"}
session:abc → {"user_id": 123, "expires_at": "2024-01-01"}
```

**Pros**:
- **Very Fast**: O(1) lookups.
- **Simple**: No complex queries.

**Cons**:
- **No Queries**: Can only lookup by key (no range queries, no joins).

**Use Case**: Caching, session storage, rate limiting.

#### 3. Column-Family Stores (Cassandra, HBase)
**Data Model**: Rows with dynamic columns (wide-column).

**Example**:
```
Row Key: user:123
  name: Alice
  email: alice@example.com
  post:1: {"title": "Hello"}
  post:2: {"title": "World"}
```

**Pros**:
- **Write-Heavy**: Optimized for writes (LSM tree).
- **Horizontal Scaling**: Easy to shard (partition key).

**Cons**:
- **Complex Queries**: No joins, limited aggregations.

**Use Case**: Time-series data, IoT, analytics.

#### 4. Graph Databases (Neo4j, Amazon Neptune)
**Data Model**: Nodes and edges (relationships).

**Example**:
```
(Alice) -[FRIENDS_WITH]-> (Bob)
(Alice) -[LIKES]-> (Post:123)
```

**Pros**:
- **Relationship Queries**: Fast traversal (e.g., "friends of friends").

**Cons**:
- **Specialized**: Not general-purpose.

**Use Case**: Social networks, recommendation engines, fraud detection.

---

## SQL vs NoSQL Comparison

| Feature | SQL | NoSQL |
| :--- | :--- | :--- |
| **Schema** | Fixed (tables, columns) | Flexible (schema-less or dynamic) |
| **Transactions** | **ACID** (strong consistency) | **BASE** (eventual consistency) |
| **Scalability** | Vertical (hard to shard) | **Horizontal** (easy to shard) |
| **Queries** | **Complex** (joins, aggregations) | Simple (key-value, limited joins) |
| **Consistency** | **Strong** | Eventual (configurable) |
| **Use Case** | Banking, e-commerce, ERP | Social media, IoT, real-time analytics |

---

## Eventual Consistency vs Strong Consistency

| Type | Description | Example |
| :--- | :--- | :--- |
| **Strong Consistency** | Reads always return the latest write. | SQL databases (ACID). |
| **Eventual Consistency** | Reads may return stale data, but will eventually be consistent. | NoSQL databases (Cassandra, DynamoDB). |

### Example: Eventual Consistency
```
Write: SET user:123 = "Alice"
Read (immediately): GET user:123 → "Bob" (stale data)
Read (after 100ms): GET user:123 → "Alice" (eventually consistent)
```

**Trade-off**: Eventual consistency allows **higher availability** and **lower latency** (no need to wait for all replicas to confirm).

---

## When to Use SQL

1. **ACID Transactions**: Banking, e-commerce (need strong consistency).
2. **Complex Queries**: Reporting, analytics (joins, aggregations).
3. **Structured Data**: Data fits well into tables (users, orders, products).
4. **Mature Ecosystem**: Need robust tooling (ORMs, migrations, backups).

---

## When to Use NoSQL

1. **Flexible Schema**: Rapidly changing data model (startups, prototyping).
2. **Horizontal Scaling**: Need to handle massive scale (billions of rows).
3. **High Write Throughput**: Time-series data, IoT, logs.
4. **Denormalized Data**: Data is naturally nested (user profiles with posts).
5. **Eventual Consistency**: Can tolerate stale reads (social media feeds).

---

## Polyglot Persistence

**Polyglot Persistence** = Using multiple database types in the same system.

### Example Architecture
```
- **Postgres** (SQL): User accounts, orders (ACID transactions).
- **Redis** (Key-Value): Session storage, caching.
- **Elasticsearch** (Search): Full-text search.
- **Cassandra** (Column-Family): Time-series metrics.
```

**Benefit**: Use the right tool for each job.

**Challenge**: Increased complexity (multiple databases to manage).

---

## Go Context: SQL vs NoSQL

### SQL (Postgres)
```go
import "database/sql"

db, _ := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
rows, _ := db.Query("SELECT id, name FROM users WHERE email = $1", "alice@example.com")
defer rows.Close()
for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
    fmt.Println(id, name)
}
```

### NoSQL (MongoDB)
```go
import "go.mongodb.org/mongo-driver/mongo"

client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
collection := client.Database("mydb").Collection("users")
var user User
collection.FindOne(context.Background(), bson.M{"email": "alice@example.com"}).Decode(&user)
fmt.Println(user.Name)
```

### NoSQL (Redis)
```go
import "github.com/go-redis/redis/v8"

rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
val, _ := rdb.Get(context.Background(), "user:123").Result()
fmt.Println(val)
```

---

## Interview Questions

### Q: When would you use NoSQL over SQL?
**A**: 
- **Flexible schema** (rapidly changing data model).
- **Horizontal scaling** (need to handle billions of rows).
- **High write throughput** (time-series, IoT, logs).
- **Eventual consistency** is acceptable (social media feeds).

### Q: What is eventual consistency?
**A**: Reads may return stale data, but will eventually be consistent. This allows higher availability and lower latency (no need to wait for all replicas to confirm).

### Q: What is polyglot persistence?
**A**: Using multiple database types in the same system (e.g., Postgres for transactions, Redis for caching, Elasticsearch for search). Use the right tool for each job.

### Q: Can you give an example of when SQL is better than NoSQL?
**A**: **Banking** (need ACID transactions for money transfers). NoSQL's eventual consistency is unacceptable (can't have stale account balances).
