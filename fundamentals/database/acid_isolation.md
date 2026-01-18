# ACID & Isolation Levels

## ACID Properties
Transactions in a database must satisfy these four properties:

1.  **Atomicity**: "All or Nothing". Either every statement in the transaction succeeds, or the entire transaction fails. No partial updates.
2.  **Consistency**: The database moves from one valid state to another. Constraints (Foreign Keys, Unique) are always enforced.
3.  **Isolation**: Concurrent transactions results in the same state as if they were executed sequentially.
4.  **Durability**: Once a transaction is committed, it remains committed even if the power goes out (Write-Ahead Logging).

## Isolation Levels (The "I" in ACID)
Low isolation = High Performance + Data Issues.
High isolation = Low Performance + Data Integrity.

| Level | Dirty Read? | Non-Repeatable Read? | Phantom Read? |
| :--- | :--- | :--- | :--- |
| **Read Uncommitted** | ✅ Yes | ✅ Yes | ✅ Yes |
| **Read Committed** (Postgres Default) | ❌ No | ✅ Yes | ✅ Yes |
| **Repeatable Read** (MySQL Default) | ❌ No | ❌ No | ✅ Yes |
| **Serializable** | ❌ No | ❌ No | ❌ No |

### Definitions
-   **Dirty Read**: Reading uncommitted data from another transaction.
-   **Non-Repeatable Read**: Reading the same row twice gets different values (because someone updated it).
-   **Phantom Read**: A range query (e.g., `WHERE age > 10`) gets different number of rows (because someone inserted/deleted).
