# Database Replication

## What is Replication?
**Replication** = Copying data from one database (primary) to one or more databases (replicas) to improve:
1. **Availability**: If the primary fails, a replica can take over.
2. **Read Scalability**: Distribute read queries across replicas.
3. **Disaster Recovery**: Replicas in different data centers.

## Master-Slave (Primary-Replica) Replication

### Architecture
- **Primary (Master)**: Handles all writes.
- **Replicas (Slaves)**: Handle reads. Receive updates from the primary.

```
Client (Write) → Primary
Client (Read) → Replica 1, Replica 2, Replica 3
```

### How It Works
1. **Write** goes to the primary.
2. **Primary** logs the change (write-ahead log, binlog).
3. **Replicas** pull the log and apply the changes.

### Pros
- **Read Scalability**: Add more replicas to handle more reads.
- **Simple**: One source of truth (primary).

### Cons
- **Single Point of Failure**: If the primary fails, writes stop (until failover).
- **Replication Lag**: Replicas may be slightly behind the primary (eventual consistency).

## Multi-Master Replication

### Architecture
- **Multiple Primaries**: All nodes accept writes.
- **Conflict Resolution**: Required when two nodes update the same row.

```
Client (Write) → Primary 1 or Primary 2
```

### Conflict Resolution Strategies
1. **Last Write Wins (LWW)**: Use timestamp to determine the winner.
2. **Application-Level**: Let the application decide (e.g., merge changes).
3. **CRDT (Conflict-Free Replicated Data Types)**: Data structures designed to merge automatically.

### Pros
- **High Availability**: No single point of failure.
- **Write Scalability**: Distribute writes across multiple nodes.

### Cons
- **Complexity**: Conflict resolution is hard.
- **Eventual Consistency**: Nodes may have different data temporarily.

## Synchronous vs Asynchronous Replication

| Type | Description | Pros | Cons |
| :--- | :--- | :--- | :--- |
| **Synchronous** | Primary waits for replicas to confirm before committing. | **Strong consistency** (replicas always up-to-date). | **Slower writes** (network latency). |
| **Asynchronous** | Primary commits immediately, replicas catch up later. | **Fast writes**. | **Replication lag** (replicas may be stale). |
| **Semi-Synchronous** | Primary waits for at least one replica. | Balance between speed and consistency. | Still some lag on other replicas. |

### Example: MySQL Replication Modes
- **Asynchronous** (default): Fast writes, but replicas may lag.
- **Semi-Synchronous**: Primary waits for at least one replica to acknowledge.

## Replication Lag

**Replication Lag** = Time delay between a write on the primary and when it appears on the replica.

### Causes
1. **Network Latency**: Slow network between primary and replica.
2. **High Write Load**: Replicas can't keep up with the primary.
3. **Long-Running Queries**: Replica is busy with a slow query.

### Problems
1. **Read-After-Write Inconsistency**: User writes data, then reads from a replica (data not there yet).
2. **Monotonic Read Violation**: User reads from Replica 1 (sees new data), then Replica 2 (sees old data).

### Solutions
1. **Read from Primary**: For critical reads (e.g., user's own data).
2. **Session Consistency**: Route all reads for a session to the same replica.
3. **Lag Monitoring**: Alert if lag exceeds threshold.

## Failover

**Failover** = Promoting a replica to primary when the primary fails.

### Automatic Failover
1. **Detect Failure**: Monitor primary (heartbeat, health checks).
2. **Elect New Primary**: Choose a replica (usually the one with the least lag).
3. **Promote Replica**: Make it the new primary.
4. **Redirect Writes**: Update DNS or load balancer to point to the new primary.

### Challenges
1. **Split-Brain**: Two nodes think they're the primary (data divergence).
2. **Data Loss**: If using asynchronous replication, uncommitted writes on the old primary are lost.

### Solutions
1. **Quorum**: Require majority vote to elect a new primary.
2. **Fencing**: Prevent the old primary from accepting writes (e.g., cut network access).

## Go Context: Simulating Read Replicas

### Database Connection Routing
```go
package main

import (
    "database/sql"
    "fmt"
    "math/rand"
)

type DB struct {
    primary  *sql.DB
    replicas []*sql.DB
}

func (db *DB) Write(query string, args ...interface{}) error {
    _, err := db.primary.Exec(query, args...)
    return err
}

func (db *DB) Read(query string, args ...interface{}) (*sql.Rows, error) {
    // Load balance reads across replicas
    replica := db.replicas[rand.Intn(len(db.replicas))]
    return replica.Query(query, args...)
}

func main() {
    primary, _ := sql.Open("mysql", "user:pass@tcp(primary:3306)/db")
    replica1, _ := sql.Open("mysql", "user:pass@tcp(replica1:3306)/db")
    replica2, _ := sql.Open("mysql", "user:pass@tcp(replica2:3306)/db")
    
    db := &DB{
        primary:  primary,
        replicas: []*sql.DB{replica1, replica2},
    }
    
    // Write to primary
    db.Write("INSERT INTO users (name) VALUES (?)", "Alice")
    
    // Read from replica
    rows, _ := db.Read("SELECT name FROM users")
    defer rows.Close()
    for rows.Next() {
        var name string
        rows.Scan(&name)
        fmt.Println(name)
    }
}
```

## Interview Questions

### Q: What's the difference between synchronous and asynchronous replication?
**A**: 
- **Synchronous**: Primary waits for replicas to confirm (strong consistency, slower writes).
- **Asynchronous**: Primary commits immediately (fast writes, replication lag).

### Q: How do you handle replication lag?
**A**: 
1. **Read from primary** for critical data (e.g., user's own data).
2. **Session consistency** (route all reads for a session to the same replica).
3. **Monitor lag** and alert if it exceeds threshold.

### Q: What is split-brain in failover?
**A**: When two nodes both think they're the primary (e.g., network partition). This causes data divergence. **Fix**: Use **quorum** (majority vote) or **fencing** (prevent old primary from accepting writes).

### Q: When would you use multi-master replication?
**A**: When you need **high availability** (no single point of failure) and **write scalability** (distribute writes across nodes). But it requires **conflict resolution** (complex).
