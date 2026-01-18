# Query Optimization & Execution Plans

## Query Execution Plan (EXPLAIN)

An **execution plan** shows how the database will execute a query (which indexes to use, join order, etc.).

### Example: EXPLAIN in Postgres
```sql
EXPLAIN SELECT * FROM users WHERE email = 'alice@example.com';
```

**Output**:
```
Seq Scan on users  (cost=0.00..35.50 rows=1 width=100)
  Filter: (email = 'alice@example.com'::text)
```

**Interpretation**:
- **Seq Scan**: Full table scan (slow for large tables).
- **cost=0.00..35.50**: Estimated cost (startup..total).
- **rows=1**: Estimated number of rows returned.

### Example: With Index
```sql
CREATE INDEX idx_users_email ON users(email);
EXPLAIN SELECT * FROM users WHERE email = 'alice@example.com';
```

**Output**:
```
Index Scan using idx_users_email on users  (cost=0.29..8.30 rows=1 width=100)
  Index Cond: (email = 'alice@example.com'::text)
```

**Interpretation**: Now uses **Index Scan** (much faster).

## Index Selection

### When to Use Indexes
1. **WHERE clause**: Filter rows (e.g., `WHERE email = 'alice@example.com'`).
2. **JOIN**: Join on indexed columns (e.g., `JOIN orders ON users.id = orders.user_id`).
3. **ORDER BY**: Sort results (e.g., `ORDER BY created_at`).

### When NOT to Use Indexes
1. **Small tables**: Full table scan is faster than index lookup.
2. **High cardinality**: Columns with few unique values (e.g., `gender` with values 'M', 'F').
3. **Frequent updates**: Indexes slow down writes (must update index on every INSERT/UPDATE/DELETE).

### Covering Index
An index that contains **all columns** needed for a query (no need to access the table).

**Example**:
```sql
CREATE INDEX idx_users_email_name ON users(email, name);
SELECT name FROM users WHERE email = 'alice@example.com';
```

**Benefit**: Database only reads the index (faster).

## Query Rewriting Techniques

### 1. Avoid SELECT *
**Bad**:
```sql
SELECT * FROM users WHERE id = 123;
```

**Good**:
```sql
SELECT id, name, email FROM users WHERE id = 123;
```

**Why**: Reduces data transfer and allows covering indexes.

### 2. Use LIMIT
**Bad**:
```sql
SELECT * FROM users ORDER BY created_at DESC;
```

**Good**:
```sql
SELECT * FROM users ORDER BY created_at DESC LIMIT 10;
```

**Why**: Database can stop scanning after 10 rows.

### 3. Avoid Functions in WHERE Clause
**Bad**:
```sql
SELECT * FROM users WHERE YEAR(created_at) = 2023;
```

**Good**:
```sql
SELECT * FROM users WHERE created_at >= '2023-01-01' AND created_at < '2024-01-01';
```

**Why**: Functions prevent index usage (database must compute `YEAR()` for every row).

### 4. Use EXISTS Instead of IN (for subqueries)
**Bad**:
```sql
SELECT * FROM users WHERE id IN (SELECT user_id FROM orders);
```

**Good**:
```sql
SELECT * FROM users WHERE EXISTS (SELECT 1 FROM orders WHERE orders.user_id = users.id);
```

**Why**: `EXISTS` stops scanning as soon as a match is found (faster for large subqueries).

### 5. Avoid OR (use UNION)
**Bad**:
```sql
SELECT * FROM users WHERE name = 'Alice' OR email = 'alice@example.com';
```

**Good**:
```sql
SELECT * FROM users WHERE name = 'Alice'
UNION
SELECT * FROM users WHERE email = 'alice@example.com';
```

**Why**: `OR` prevents index usage (database must scan both conditions). `UNION` allows separate index scans.

## N+1 Query Problem

**Problem**: Fetching a list of objects, then fetching related objects one by one.

### Example: N+1 Problem
```go
// Fetch all users (1 query)
users := []User{}
db.Find(&users)

// Fetch posts for each user (N queries)
for _, user := range users {
    db.Where("user_id = ?", user.ID).Find(&user.Posts)
}
```

**Total**: 1 + N queries (if 100 users, 101 queries).

### Solution 1: Eager Loading (JOIN)
```go
db.Preload("Posts").Find(&users)
```

**SQL**:
```sql
SELECT * FROM users;
SELECT * FROM posts WHERE user_id IN (1, 2, 3, ...);
```

**Total**: 2 queries (much faster).

### Solution 2: Batch Loading
```go
// Fetch all users
users := []User{}
db.Find(&users)

// Fetch all posts in one query
userIDs := []int{}
for _, user := range users {
    userIDs = append(userIDs, user.ID)
}
posts := []Post{}
db.Where("user_id IN ?", userIDs).Find(&posts)

// Map posts to users
postsByUser := make(map[int][]Post)
for _, post := range posts {
    postsByUser[post.UserID] = append(postsByUser[post.UserID], post)
}
for i, user := range users {
    users[i].Posts = postsByUser[user.ID]
}
```

## Join Optimization

### Types of Joins
1. **Nested Loop Join**: For each row in table A, scan table B (slow for large tables).
2. **Hash Join**: Build a hash table for one table, probe with the other (fast for large tables).
3. **Merge Join**: Sort both tables, then merge (fast if already sorted).

### Example: Force Hash Join (Postgres)
```sql
SET enable_nestloop = off;
SELECT * FROM users JOIN orders ON users.id = orders.user_id;
```

## Go Context: Avoiding N+1 with GORM

### N+1 Problem
```go
var users []User
db.Find(&users)
for _, user := range users {
    db.Model(&user).Association("Posts").Find(&user.Posts) // N queries
}
```

### Solution: Preload
```go
var users []User
db.Preload("Posts").Find(&users) // 2 queries
```

### Solution: Joins (Single Query)
```go
var users []User
db.Joins("JOIN posts ON posts.user_id = users.id").Find(&users) // 1 query
```

## Interview Questions

### Q: What is EXPLAIN and why is it useful?
**A**: `EXPLAIN` shows the query execution plan (how the database will execute the query). It helps identify slow queries (e.g., full table scans) and optimize them (e.g., add indexes).

### Q: What is a covering index?
**A**: An index that contains all columns needed for a query (no need to access the table). This makes queries faster because the database only reads the index.

### Q: What is the N+1 query problem?
**A**: Fetching a list of objects (1 query), then fetching related objects one by one (N queries). **Fix**: Use **eager loading** (JOIN or Preload) to fetch all data in 1-2 queries.

### Q: Why should you avoid functions in WHERE clauses?
**A**: Functions prevent index usage (database must compute the function for every row). Instead, rewrite the query to use indexed columns directly.

### Q: When should you use EXISTS instead of IN?
**A**: For subqueries, `EXISTS` is faster because it stops scanning as soon as a match is found. `IN` must scan the entire subquery.
