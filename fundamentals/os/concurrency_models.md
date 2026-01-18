# Concurrency Models

## 1. Preemptive Multitasking (OS Standard)
The OS decides when to stop one task and start another.
-   **Pros**: Fairness. A tight loop in one program won't freeze the system.
-   **Cons**: Context switching overhead. You have no control over *when* the switch happens (danger of race conditions).

## 2. Cooperative Multitasking (Old systems, some runtimes)
The task must explicitly say "I'm done for now" (yield) to let others run.
-   **Pros**: Very predictable. Zero uncontrolled context switches.
-   **Cons**: One bad task can freeze the entire system.

## 3. Go's Model (Hybrid M:N)
Go uses **M** goroutines mapped onto **N** OS threads.
-   **Preemption**: Since Go 1.14, the runtime signals goroutines to stop (asynchronous preemption) ensuring loops don't block GC.

## Deadlock vs Race Condition vs Starvation

### Deadlock
Thread A holds Lock 1 and waits for Lock 2.
Thread B holds Lock 2 and waits for Lock 1.
Result: **Neither proceeds.**

### Race Condition
Two threads access shared data concurrently, and at least one is writing.
Result: **Undefined behavior / Data Corruption.**
*Fix*: Mutex, Atomic, or Channel.

### Starvation
A thread is perpetually denied resources (CPU/locks) because other "greedy" high-priority threads keep taking them.
Result: **Thread works, but very slowly or never finishes.**
