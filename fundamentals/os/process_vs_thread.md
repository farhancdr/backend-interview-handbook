# Process vs Thread

## The Core Difference
A **Process** is an instance of a program in execution. It is an independent entity to which system resources (CPU time, memory, etc.) are allocated.

A **Thread** is the smallest sequence of programmed instructions that can be managed independently by a scheduler. Threads exist *within* a process.

| Feature | Process | Thread |
| :--- | :--- | :--- |
| **Memory** | Independent memory space (Stack + Heap separate). | Shared memory space (Shared Heap, separate Stack). |
| **Communication** | IPC (Inter-Process Communication) needed (Pipes, Sockets). | Direct variable access via shared memory. |
| **Creation Cost** | High (Heavyweight). Requires OS to allocate new memory. | Low (Lightweight). Just a new stack. |
| **Failure** | If one crashes, others usually unaffected. | If one crashes (e.g., panic), the entire process dies. |

## Context Switching
Switching between threads of the same process is faster than switching between processes because:
1.  **Virtual Memory**: Threads share the same page directory (memory map), so the TLB (Translation Lookaside Buffer) doesn't need to be flushed.
2.  **Cache**: Process switch invalidates CPU caches more aggressively.

## Go Context
In Go, a **Goroutine** is an even lighter weight abstraction:
-   **OS Thread**: ~1-2MB stack.
-   **Goroutine**: ~2KB stack (grows dynamically).
-   **Scheduling**: OS threads are scheduled by the OS Kernel. Goroutines are scheduled by the Go Runtime (M:N scheduler) onto OS threads.
