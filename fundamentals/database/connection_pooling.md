# Database Connection Pooling

## What is Connection Pooling?
A **connection pool** is a cache of database connections that can be reused, rather than creating a new connection for every request.

### Why Connection Pooling?
1. **Expensive**: Creating a TCP connection + authentication is slow (~10-100ms).
2. **Limited**: Databases have a maximum number of connections (e.g., Postgres default: 100).
3. **Reuse**: Most requests are short-lived (query takes <10ms, but connection setup takes 50ms).

### Without Pooling
```
Request 1: Create connection → Query → Close connection
Request 2: Create connection → Query → Close connection
Request 3: Create connection → Query → Close connection
```

**Problem**: Slow (connection overhead per request).

### With Pooling
```
Request 1: Get connection from pool → Query → Return to pool
Request 2: Get connection from pool → Query → Return to pool
Request 3: Get connection from pool → Query → Return to pool
```

**Benefit**: Fast (reuse existing connections).

## Connection Lifecycle

### States
1. **Idle**: Connection is in the pool, waiting to be used.
2. **In Use**: Connection is being used by a request.
3. **Closed**: Connection is closed (e.g., timeout, error).

### Flow
1. **Acquire**: Get a connection from the pool (or create a new one if pool is empty).
2. **Use**: Execute query.
3. **Release**: Return connection to the pool.
4. **Cleanup**: Close idle connections after timeout.

## Pool Sizing Strategies

### Formula (Rule of Thumb)
```
Pool Size = (Number of CPU Cores × 2) + Effective Spindle Count
```

**Example**: 4 cores + 1 disk = 4 × 2 + 1 = **9 connections**.

### Considerations
1. **Too Small**: Requests wait for connections (high latency).
2. **Too Large**: Database is overwhelmed (context switching, memory usage).
3. **Database Limit**: Don't exceed the database's max connections (e.g., Postgres: 100).

### Recommended Settings
- **Min Connections**: 2-5 (keep some connections warm).
- **Max Connections**: 10-20 (depends on workload).
- **Idle Timeout**: 5-10 minutes (close idle connections).
- **Max Lifetime**: 30-60 minutes (prevent stale connections).

## Go's database/sql Package

Go's `database/sql` package has a **built-in connection pool**.

### Default Settings
```go
db, _ := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
// Default: MaxOpenConns = unlimited, MaxIdleConns = 2
```

### Configure Pool
```go
db.SetMaxOpenConns(25)        // Max open connections (in use + idle)
db.SetMaxIdleConns(5)         // Max idle connections
db.SetConnMaxLifetime(5 * time.Minute)  // Max connection lifetime
db.SetConnMaxIdleTime(1 * time.Minute)  // Max idle time before closing
```

### Example: Full Configuration
```go
package main

import (
    "database/sql"
    "fmt"
    "time"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    
    // Connection pool settings
    db.SetMaxOpenConns(25)                 // Max 25 connections
    db.SetMaxIdleConns(5)                  // Keep 5 idle connections
    db.SetConnMaxLifetime(5 * time.Minute) // Close connections after 5 minutes
    db.SetConnMaxIdleTime(1 * time.Minute) // Close idle connections after 1 minute
    
    // Test connection
    if err := db.Ping(); err != nil {
        panic(err)
    }
    
    fmt.Println("Connected to database with connection pool")
}
```

## Monitoring Pool Health

### Check Pool Stats
```go
stats := db.Stats()
fmt.Printf("Open Connections: %d\n", stats.OpenConnections)
fmt.Printf("In Use: %d\n", stats.InUse)
fmt.Printf("Idle: %d\n", stats.Idle)
fmt.Printf("Wait Count: %d\n", stats.WaitCount)       // Requests that waited for a connection
fmt.Printf("Wait Duration: %s\n", stats.WaitDuration) // Total time spent waiting
fmt.Printf("Max Idle Closed: %d\n", stats.MaxIdleClosed)
fmt.Printf("Max Lifetime Closed: %d\n", stats.MaxLifetimeClosed)
```

### Metrics to Watch
1. **WaitCount**: High → Pool is too small (increase MaxOpenConns).
2. **WaitDuration**: High → Requests are waiting too long (increase pool size).
3. **MaxIdleClosed**: High → Too many idle connections (reduce MaxIdleConns or IdleTimeout).

## Common Pitfalls

### 1. Not Closing Rows
**Bad**:
```go
rows, _ := db.Query("SELECT * FROM users")
// Forgot to close rows → Connection is not returned to pool
```

**Good**:
```go
rows, _ := db.Query("SELECT * FROM users")
defer rows.Close() // Always close rows
```

### 2. Not Using Prepared Statements
**Bad**:
```go
for _, email := range emails {
    db.Exec("SELECT * FROM users WHERE email = '" + email + "'") // SQL injection risk
}
```

**Good**:
```go
stmt, _ := db.Prepare("SELECT * FROM users WHERE email = $1")
defer stmt.Close()
for _, email := range emails {
    stmt.Query(email) // Reuse prepared statement
}
```

### 3. Pool Size Too Large
**Problem**: Database is overwhelmed (e.g., Postgres max connections = 100, but you have 10 app servers with 20 connections each = 200 connections).

**Fix**: Reduce MaxOpenConns or use a connection pooler (PgBouncer).

## Connection Poolers (External)

### PgBouncer (Postgres)
A lightweight connection pooler that sits between the application and Postgres.

**Modes**:
1. **Session**: One connection per session (default).
2. **Transaction**: One connection per transaction (more efficient).
3. **Statement**: One connection per statement (most efficient, but breaks transactions).

**Benefit**: 1000 app connections → 20 Postgres connections (reduces database load).

## Interview Questions

### Q: Why is connection pooling important?
**A**: Creating a database connection is expensive (~10-100ms). Connection pooling reuses connections, reducing latency and database load.

### Q: How do you size a connection pool?
**A**: Rule of thumb: `(CPU Cores × 2) + Disk Count`. Monitor `WaitCount` and `WaitDuration` to adjust. Don't exceed the database's max connections.

### Q: What happens if you don't close rows in Go?
**A**: The connection is not returned to the pool (connection leak). Eventually, the pool runs out of connections and requests block.

### Q: What is PgBouncer?
**A**: A connection pooler for Postgres that sits between the application and database. It allows many app connections to share a smaller pool of database connections (reduces database load).

### Q: What's the difference between MaxOpenConns and MaxIdleConns?
**A**: 
- **MaxOpenConns**: Max total connections (in use + idle).
- **MaxIdleConns**: Max idle connections (waiting in the pool). Idle connections above this limit are closed.
