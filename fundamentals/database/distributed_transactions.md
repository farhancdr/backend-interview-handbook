# Distributed Transactions

## The Problem
In a distributed system with multiple databases/services, how do you ensure **all-or-nothing** semantics (atomicity)?

### Example: E-Commerce Order
```
1. Deduct inventory (Inventory Service)
2. Charge payment (Payment Service)
3. Create order (Order Service)
```

**Problem**: What if step 2 fails? Inventory is already deducted (inconsistent state).

## Two-Phase Commit (2PC)

**2PC** is a protocol to ensure atomicity across multiple databases.

### Roles
- **Coordinator**: Orchestrates the transaction.
- **Participants**: Databases/services involved in the transaction.

### Phases

#### Phase 1: Prepare (Voting)
1. **Coordinator** sends `PREPARE` to all participants.
2. **Participants** execute the transaction locally (but don't commit).
3. **Participants** respond with `YES` (ready to commit) or `NO` (abort).

#### Phase 2: Commit/Abort
1. **If all participants vote YES**:
   - Coordinator sends `COMMIT` to all participants.
   - Participants commit the transaction.
2. **If any participant votes NO**:
   - Coordinator sends `ABORT` to all participants.
   - Participants rollback the transaction.

### Example Flow
```
Coordinator → Participant A: PREPARE
Coordinator → Participant B: PREPARE

Participant A → Coordinator: YES
Participant B → Coordinator: YES

Coordinator → Participant A: COMMIT
Coordinator → Participant B: COMMIT
```

### Pros
- **Atomicity**: All participants commit or all abort (strong consistency).

### Cons
- **Blocking**: If the coordinator crashes after `PREPARE`, participants are **blocked** (can't commit or abort).
- **Slow**: 2 round trips (high latency).
- **Single Point of Failure**: Coordinator failure blocks the system.

### When to Use
- **Rare**: Only when strong consistency is absolutely required (e.g., financial transactions).
- **Not Scalable**: Doesn't work well in high-throughput systems.

## Saga Pattern

**Saga** = A sequence of **local transactions**, each with a **compensating transaction** (undo).

### Types of Sagas

#### 1. Choreography (Event-Driven)
Each service listens for events and triggers the next step.

**Example**:
```
1. Order Service: Create order → Emit "OrderCreated" event
2. Inventory Service: Deduct inventory → Emit "InventoryDeducted" event
3. Payment Service: Charge payment → Emit "PaymentCharged" event
```

**If Payment Fails**:
```
Payment Service: Emit "PaymentFailed" event
Inventory Service: Restore inventory (compensating transaction)
Order Service: Cancel order (compensating transaction)
```

**Pros**: Decoupled (no central coordinator).  
**Cons**: Hard to debug (event flow is complex).

#### 2. Orchestration (Centralized)
A central orchestrator coordinates the saga.

**Example**:
```
Orchestrator:
  1. Call Inventory Service: Deduct inventory
  2. Call Payment Service: Charge payment
  3. Call Order Service: Create order
```

**If Payment Fails**:
```
Orchestrator:
  1. Call Inventory Service: Restore inventory (compensating transaction)
  2. Call Order Service: Cancel order (compensating transaction)
```

**Pros**: Easy to debug (centralized logic).  
**Cons**: Single point of failure (orchestrator).

### Comparison: Choreography vs Orchestration

| Feature | Choreography | Orchestration |
| :--- | :--- | :--- |
| **Coordination** | Decentralized (events) | Centralized (orchestrator) |
| **Coupling** | Loose | Tight |
| **Debugging** | Hard (event flow) | Easy (centralized logic) |
| **Failure Handling** | Complex (each service must handle events) | Simple (orchestrator handles) |

## Eventual Consistency

**Eventual Consistency** = The system will eventually reach a consistent state, but may be temporarily inconsistent.

### Example: Social Media Feed
```
User posts a photo → Photo Service stores photo
                   → Feed Service updates feed (async)
```

**Temporary Inconsistency**: Photo is stored, but feed is not updated yet (users don't see the photo immediately).

**Eventually Consistent**: Feed is updated after a few seconds.

### Pros
- **High Availability**: No blocking (services can fail independently).
- **Low Latency**: No need to wait for all services.

### Cons
- **Complexity**: Application must handle inconsistent states.

## Idempotency

**Idempotency** = Performing the same operation multiple times has the same effect as performing it once.

### Why It Matters
In distributed systems, **retries** are common (network failures, timeouts). Idempotency ensures retries don't cause duplicate operations.

### Example: Non-Idempotent
```
POST /transfer
{
  "from": "Alice",
  "to": "Bob",
  "amount": 100
}
```

**Problem**: If the request is retried, $100 is transferred twice.

### Example: Idempotent
```
POST /transfer
{
  "idempotency_key": "abc123",
  "from": "Alice",
  "to": "Bob",
  "amount": 100
}
```

**Solution**: Server checks if `idempotency_key` has been processed. If yes, return the previous result (don't transfer again).

### Implementation (Go)
```go
var processedKeys = make(map[string]bool)
var mu sync.Mutex

func Transfer(idempotencyKey string, from, to string, amount int) error {
    mu.Lock()
    defer mu.Unlock()
    
    // Check if already processed
    if processedKeys[idempotencyKey] {
        return nil // Already processed, return success
    }
    
    // Process transfer
    // ... (deduct from Alice, add to Bob)
    
    // Mark as processed
    processedKeys[idempotencyKey] = true
    return nil
}
```

## Go Context: Saga with Orchestration

### Example: Order Saga
```go
package main

import "fmt"

type Orchestrator struct{}

func (o *Orchestrator) CreateOrder(orderID int, userID int, amount int) error {
    // Step 1: Deduct inventory
    if err := o.deductInventory(orderID); err != nil {
        return err
    }
    
    // Step 2: Charge payment
    if err := o.chargePayment(userID, amount); err != nil {
        // Compensate: Restore inventory
        o.restoreInventory(orderID)
        return err
    }
    
    // Step 3: Create order
    if err := o.createOrder(orderID, userID); err != nil {
        // Compensate: Refund payment + restore inventory
        o.refundPayment(userID, amount)
        o.restoreInventory(orderID)
        return err
    }
    
    return nil
}

func (o *Orchestrator) deductInventory(orderID int) error {
    fmt.Println("Deducting inventory for order", orderID)
    return nil
}

func (o *Orchestrator) restoreInventory(orderID int) {
    fmt.Println("Restoring inventory for order", orderID)
}

func (o *Orchestrator) chargePayment(userID int, amount int) error {
    fmt.Println("Charging payment for user", userID, "amount", amount)
    return fmt.Errorf("payment failed") // Simulate failure
}

func (o *Orchestrator) refundPayment(userID int, amount int) {
    fmt.Println("Refunding payment for user", userID, "amount", amount)
}

func (o *Orchestrator) createOrder(orderID int, userID int) error {
    fmt.Println("Creating order", orderID, "for user", userID)
    return nil
}

func main() {
    orch := &Orchestrator{}
    err := orch.CreateOrder(123, 456, 100)
    if err != nil {
        fmt.Println("Order failed:", err)
    }
}
```

## Interview Questions

### Q: What is Two-Phase Commit (2PC)?
**A**: A protocol to ensure atomicity across multiple databases. **Phase 1**: Coordinator asks participants to prepare (vote YES/NO). **Phase 2**: If all vote YES, commit; otherwise, abort. **Drawback**: Blocking (participants wait for coordinator).

### Q: What is the Saga pattern?
**A**: A sequence of local transactions, each with a compensating transaction (undo). If a step fails, execute compensating transactions to rollback. **Types**: Choreography (event-driven) vs Orchestration (centralized).

### Q: What is idempotency and why is it important?
**A**: Performing the same operation multiple times has the same effect as performing it once. Important in distributed systems because **retries** are common (network failures). Use **idempotency keys** to prevent duplicate operations.

### Q: When would you use 2PC vs Saga?
**A**: 
- **2PC**: When strong consistency is required (e.g., financial transactions). Rare, not scalable.
- **Saga**: When eventual consistency is acceptable (e.g., e-commerce orders). More scalable, but requires compensating transactions.
