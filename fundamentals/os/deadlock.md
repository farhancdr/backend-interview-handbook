# Deadlock Detection & Prevention

## What is Deadlock?
A situation where two or more processes are **permanently blocked**, each waiting for a resource held by another.

### Classic Example: Dining Philosophers
5 philosophers sit at a table. Each needs 2 forks to eat. There are only 5 forks. If each picks up the left fork simultaneously, all wait forever for the right fork → **Deadlock**.

## Coffman Conditions (All 4 Must Hold for Deadlock)
1. **Mutual Exclusion**: At least one resource must be non-shareable (only one process can use it at a time).
2. **Hold and Wait**: A process holding at least one resource is waiting to acquire additional resources held by others.
3. **No Preemption**: Resources cannot be forcibly taken away; they must be released voluntarily.
4. **Circular Wait**: A circular chain of processes exists, where each process holds a resource needed by the next.

## Deadlock Strategies

### 1. Deadlock Prevention
**Break one of the Coffman conditions.**

| Condition | Prevention Strategy | Example |
| :--- | :--- | :--- |
| **Mutual Exclusion** | Make resources shareable. | Read-only files (multiple readers allowed). |
| **Hold and Wait** | Require processes to request all resources at once. | Acquire all locks before starting. |
| **No Preemption** | Allow forcibly taking resources. | Database transaction rollback. |
| **Circular Wait** | **Impose a total ordering on resources.** | Always acquire locks in the same order (Lock 1 → Lock 2). |

**Most Common**: **Lock Ordering** (prevents circular wait).

### 2. Deadlock Avoidance
Use algorithms like **Banker's Algorithm** to ensure the system never enters an unsafe state.

#### Banker's Algorithm
Before granting a resource request, check if the system will remain in a **safe state** (a state where there exists a sequence of process executions that allows all to complete).

**Drawback**: Requires knowing the maximum resource needs of all processes in advance (rarely practical).

### 3. Deadlock Detection & Recovery
Allow deadlocks to occur, but detect them and recover.

#### Detection
Maintain a **Resource Allocation Graph (RAG)**:
- **Nodes**: Processes and Resources.
- **Edges**: Process → Resource (request), Resource → Process (allocation).

**Deadlock exists** if there is a **cycle** in the graph.

#### Recovery
1. **Kill one or more processes** (abort the deadlocked processes).
2. **Preempt resources** (rollback a process and take its resources).

### 4. Ignore the Problem (Ostrich Algorithm)
Assume deadlocks are rare and let the user reboot the system.

**Used by**: Many operating systems (Windows, Linux) for application-level deadlocks.

## Deadlock in Go

### Example: Mutex Deadlock
```go
var mu1, mu2 sync.Mutex

// Goroutine 1
go func() {
    mu1.Lock()
    time.Sleep(10 * time.Millisecond)
    mu2.Lock() // Waits forever if Goroutine 2 holds mu2
    mu2.Unlock()
    mu1.Unlock()
}()

// Goroutine 2
go func() {
    mu2.Lock()
    time.Sleep(10 * time.Millisecond)
    mu1.Lock() // Waits forever if Goroutine 1 holds mu1
    mu1.Unlock()
    mu2.Unlock()
}()
```

**Fix**: **Lock Ordering** - Always acquire `mu1` before `mu2`.

### Example: Channel Deadlock
```go
ch := make(chan int)
ch <- 42 // Blocks forever (no receiver)
```

**Fix**: Use a buffered channel or spawn a goroutine to receive.

### Go's Deadlock Detector
Go's runtime detects **global deadlocks** (all goroutines are blocked):
```
fatal error: all goroutines are asleep - deadlock!
```

**Limitation**: Only detects when **all** goroutines are blocked. Partial deadlocks (some goroutines blocked) are not detected.

## Interview Questions

### Q: How do you prevent deadlock in a system with multiple locks?
**A**: **Lock Ordering** - Establish a global order for acquiring locks and always acquire them in that order.

### Q: What's the difference between deadlock and livelock?
**A**: 
- **Deadlock**: Processes are blocked and waiting forever.
- **Livelock**: Processes are actively running but making no progress (e.g., two people in a hallway both stepping aside in the same direction repeatedly).

### Q: Can you have a deadlock with just one resource?
**A**: No. Deadlock requires at least two resources (circular wait condition).
