# Indexing & Partitioning

## Indexing
Indices speed up reads but slow down writes (because the index must be updated on insert).

### B-Tree (Balanced Tree)
-   **Structure**: Balanced tree where leaf nodes contain data pointers.
-   **Use Case**: Standard SQL indices (Postgres, MySQL).
-   **Pros**: Great for range queries (`>`), equality, and sorting. Good read/write balance.

### LSM Tree (Log-Structured Merge-Tree)
-   **Structure**: Writes go to an in-memory buffer (MemTable), then flushed to disk as immutable sorted files (SSTables). Background compaction merges files.
-   **Use Case**: Write-heavy databases (Cassandra, RocksDB, LevelDB).
-   **Pros**: Extremely fast writes (append-only).
-   **Cons**: Slower reads (check MemTable + multiple SSTables).

## Scaling: Sharding vs Partitioning

### Partitioning (Vertical Sorting)
Splitting a *single* table into smaller chunks *on the same server* (usually).
-   Example: `Orders` table partitioned by `year`. `Orders_2023`, `Orders_2024`.
-   **Goal**: Manageability and performance (scan only one partition).

### Sharding (Horizontal Scaling)
Splitting data across *multiple servers*.
-   Example: Users 1-1M -> Server A. Users 1M-2M -> Server B.
-   **Goal**: Unlimited scale.
-   **Challenge**: Cross-shard joins are expensive/impossible. Distributed transactions are hard.

## CAP Theorem
In a distributed system, you can only pick 2:
1.  **Consistency**: Every read receives the most recent write.
2.  **Availability**: Every request receives a (non-error) response.
3.  **Partition Tolerance**: System continues to operate despite network messages dropping.

**Reality**: P is mandatory in distributed systems (networks fail). So you choose:
-   **CP** (Consistency > Availability): MongoDB, HBase. (Refuse writes if disconnected).
-   **AP** (Availability > Consistency): Cassandra, DynamoDB. (Accept writes, reconcile later -> Eventual Consistency).
