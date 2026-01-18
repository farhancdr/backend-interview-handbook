# Locking Mechanisms & Transactions

## Pessimistic vs Optimistic Locking

| Type | Description | When to Use | Example |
| :--- | :--- | :--- | :--- |
| **Pessimistic** | Lock the row **before** reading (assume conflicts will happen). | **High contention** (many concurrent updates). | `SELECT FOR UPDATE` |
| **Optimistic** | Don't lock. Check if data changed **before** committing. | **Low contention** (few concurrent updates). | Version number or timestamp check. |

### Pessimistic Locking Example (SQL)
```sql
BEGIN;
SELECT * FROM accounts WHERE id = 123 FOR UPDATE; -- Lock the row
UPDATE accounts SET balance = balance - 100 WHERE id = 123;
COMMIT;
```

**How It Works**: `FOR UPDATE` locks the row. Other transactions trying to read/update the same row will **block** until the lock is released.

### Optimistic Locking Example (SQL)
```sql
-- Read the current version
SELECT balance, version FROM accounts WHERE id = 123;
-- balance = 1000, version = 5

-- Update only if version hasn't changed
UPDATE accounts 
SET balance = 900, version = 6 
WHERE id = 123 AND version = 5;

-- If 0 rows updated, someone else modified it (retry)
```

**How It Works**: No locks. Before committing, check if the version number changed. If it did, someone else updated the row (conflict).

### Comparison

| Feature | Pessimistic | Optimistic |
| :--- | :--- | :--- |
| **Locking** | Locks immediately | No locks (check before commit) |
| **Contention** | High (many transactions block) | Low (conflicts are rare) |
| **Performance** | Slower (blocking) | Faster (no blocking) |
| **Use Case** | High contention (e.g., inventory, banking) | Low contention (e.g., user profiles) |

## Row-Level vs Table-Level Locks

| Type | Granularity | Concurrency | Use Case |
| :--- | :--- | :--- | :--- |
| **Row-Level** | Lock individual rows | **High** (other rows can be modified) | Modern databases (Postgres, MySQL InnoDB) |
| **Table-Level** | Lock entire table | **Low** (entire table is locked) | Legacy databases, DDL operations |

### Example: Row-Level Lock (Postgres)
```sql
BEGIN;
SELECT * FROM users WHERE id = 123 FOR UPDATE; -- Lock row 123
-- Other transactions can still update row 124
COMMIT;
```

### Example: Table-Level Lock
```sql
LOCK TABLE users IN EXCLUSIVE MODE;
-- No other transaction can read or write to the users table
```

## Deadlock Detection in Databases

**Deadlock** = Two transactions waiting for each other's locks.

### Example
```
Transaction 1:
  LOCK row A
  (wait for row B)

Transaction 2:
  LOCK row B
  (wait for row A)
```

### Detection
Databases maintain a **wait-for graph**. If a cycle is detected, one transaction is **aborted** (rolled back).

### Prevention
1. **Lock Ordering**: Always acquire locks in the same order.
2. **Timeout**: Abort transactions that wait too long.

### Example: Deadlock in Postgres
```sql
-- Transaction 1
BEGIN;
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
UPDATE accounts SET balance = balance + 100 WHERE id = 2; -- Waits if T2 locked row 2

-- Transaction 2
BEGIN;
UPDATE accounts SET balance = balance - 50 WHERE id = 2;
UPDATE accounts SET balance = balance + 50 WHERE id = 1; -- Waits if T1 locked row 1

-- Deadlock! Postgres aborts one transaction.
```

**Error**:
```
ERROR: deadlock detected
DETAIL: Process 1234 waits for ShareLock on transaction 5678; blocked by process 5678.
```

## SELECT FOR UPDATE

**Purpose**: Lock rows for update (pessimistic locking).

### Syntax
```sql
SELECT * FROM users WHERE id = 123 FOR UPDATE;
```

### Variants

| Variant | Description |
| :--- | :--- |
| `FOR UPDATE` | Exclusive lock (blocks reads and writes). |
| `FOR SHARE` | Shared lock (blocks writes, allows reads). |
| `FOR UPDATE NOWAIT` | Fail immediately if row is locked (don't wait). |
| `FOR UPDATE SKIP LOCKED` | Skip locked rows (useful for job queues). |

### Example: Job Queue (SKIP LOCKED)
```sql
BEGIN;
SELECT * FROM jobs WHERE status = 'pending' 
ORDER BY created_at 
LIMIT 1 
FOR UPDATE SKIP LOCKED; -- Skip jobs locked by other workers

UPDATE jobs SET status = 'processing' WHERE id = <job_id>;
COMMIT;
```

**Use Case**: Multiple workers processing jobs. Each worker picks a job without blocking others.

## Two-Phase Locking (2PL)

**2PL** = A protocol to ensure serializability (transactions appear to execute sequentially).

### Phases
1. **Growing Phase**: Acquire locks (cannot release any lock).
2. **Shrinking Phase**: Release locks (cannot acquire any lock).

**Rule**: Once you release a lock, you cannot acquire any more locks.

### Example
```
Transaction:
  LOCK row A
  LOCK row B
  (Growing phase ends)
  UNLOCK row A
  UNLOCK row B
  (Shrinking phase)
```

**Problem**: **Deadlocks** can still occur (two transactions in growing phase waiting for each other).

## Go Context: Optimistic Locking

### Using Version Numbers
```go
package main

import (
    "database/sql"
    "fmt"
)

type Account struct {
    ID      int
    Balance int
    Version int
}

func UpdateBalance(db *sql.DB, accountID, amount int) error {
    for {
        // Read current state
        var acc Account
        err := db.QueryRow("SELECT id, balance, version FROM accounts WHERE id = ?", accountID).
            Scan(&acc.ID, &acc.Balance, &acc.Version)
        if err != nil {
            return err
        }
        
        newBalance := acc.Balance + amount
        newVersion := acc.Version + 1
        
        // Optimistic update
        result, err := db.Exec(
            "UPDATE accounts SET balance = ?, version = ? WHERE id = ? AND version = ?",
            newBalance, newVersion, accountID, acc.Version,
        )
        if err != nil {
            return err
        }
        
        rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 0 {
            // Conflict! Retry
            fmt.Println("Conflict detected, retrying...")
            continue
        }
        
        // Success
        return nil
    }
}
```

## Interview Questions

### Q: When would you use pessimistic vs optimistic locking?
**A**: 
- **Pessimistic**: High contention (many concurrent updates, e.g., inventory, banking).
- **Optimistic**: Low contention (few concurrent updates, e.g., user profiles).

### Q: What is SELECT FOR UPDATE?
**A**: A SQL statement that locks rows for update (pessimistic locking). Other transactions trying to read/update the same rows will block until the lock is released.

### Q: How do databases detect deadlocks?
**A**: Databases maintain a **wait-for graph**. If a cycle is detected (Transaction A waits for B, B waits for A), one transaction is aborted.

### Q: What's the difference between FOR UPDATE and FOR SHARE?
**A**: 
- **FOR UPDATE**: Exclusive lock (blocks reads and writes).
- **FOR SHARE**: Shared lock (blocks writes, allows reads).

### Q: What is SKIP LOCKED used for?
**A**: Skip rows that are already locked (useful for job queues where multiple workers pick jobs without blocking each other).
