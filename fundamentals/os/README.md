# 🖥️ Operating Systems Fundamentals

> **Understand the foundation that runs your code**

Operating systems knowledge is crucial for writing efficient backend systems. This section covers process management, memory, concurrency, and more.

---

## 📖 Topics

### Process & Thread Management

| Topic | Description | Key Interview Questions |
|:------|:------------|:------------------------|
| **[Process vs Thread](process_vs_thread.md)** | Differences, context switching, scheduling | "What's the difference between a process and a thread?" |
| **[Concurrency Models](concurrency_models.md)** | Multi-threading, event loops, actor model | "Explain different concurrency models" |
| **[CPU Scheduling](cpu_scheduling.md)** | Scheduling algorithms, preemption, priorities | "Explain Round-Robin vs Priority scheduling" |
| **[Deadlock](deadlock.md)** | Conditions, prevention, detection, recovery | "How do you prevent deadlocks?" |

### Memory Management

| Topic | Description | Key Interview Questions |
|:------|:------------|:------------------------|
| **[Memory Management](memory_management.md)** | Heap vs stack, allocation strategies, fragmentation | "Explain heap vs stack memory" |
| **[Virtual Memory](virtual_memory.md)** | Paging, segmentation, page faults, TLB | "How does virtual memory work?" |

### System Interaction

| Topic | Description | Key Interview Questions |
|:------|:------------|:------------------------|
| **[System Calls](system_calls.md)** | Kernel mode vs user mode, common syscalls | "What happens during a system call?" |
| **[File Systems](file_systems.md)** | Inodes, journaling, file descriptors | "Explain how file systems organize data" |

---

## 🎯 Interview Preparation Guide

### Must-Know Concepts

1. **Process vs Thread** - Understand memory isolation and context switching costs
2. **Deadlock Conditions** - Memorize the 4 Coffman conditions
3. **Virtual Memory** - Explain page tables and TLB
4. **CPU Scheduling** - Know common algorithms (FCFS, SJF, Round-Robin, Priority)
5. **System Calls** - Understand the kernel/user mode transition
6. **Memory Hierarchy** - Registers → Cache → RAM → Disk

### Common Interview Patterns

**System Design Questions:**
- "Design a task scheduler" → Requires CPU scheduling knowledge
- "Design a memory allocator" → Requires memory management understanding
- "Explain how your web server handles 10,000 concurrent connections" → Requires concurrency model knowledge

**Technical Deep Dives:**
- "What happens when you run a program?"
- "Explain the fork() system call"
- "How does the OS handle a page fault?"
- "What causes thrashing?"

---

## 🔗 Related Topics

- **[Go Concurrency](../../internal/concurrency/)** - See how Go implements concurrency with goroutines
- **[Memory Patterns](../../internal/memory/)** - Understand Go's memory management and GC
- **[Go Internals](../../internal/internals/)** - Learn about Go's GMP scheduler

---

## 📚 Study Tips

1. **Draw diagrams** - Visualize process states, memory layouts, page tables
2. **Trace execution** - Walk through system call flows step-by-step
3. **Compare trade-offs** - Threads vs processes, paging vs segmentation
4. **Relate to Go** - Connect OS concepts to Go's runtime (goroutines, GC, scheduler)

### Hands-on Practice

```bash
# View running processes
ps aux

# Monitor system resources
top / htop

# Check memory usage
free -h

# View open file descriptors
lsof -p <pid>

# Trace system calls
strace <command>
```

---

## 🌟 Pro Tips

### For Interviews

- **Use analogies** - Compare processes to separate houses, threads to rooms in a house
- **Explain trade-offs** - More threads = more parallelism but higher context switching cost
- **Connect to real systems** - "This is why Nginx uses event loops instead of thread-per-connection"

### Common Pitfalls

- **Confusing concurrency with parallelism** - Concurrency is about structure, parallelism is about execution
- **Ignoring context switching costs** - Creating too many threads can hurt performance
- **Not understanding memory barriers** - Critical for multi-threaded programming

---

## 🔄 Connection to Go

Go abstracts many OS concepts but understanding them helps you write better code:

- **Goroutines** → Lightweight threads managed by Go runtime
- **Channels** → Safe communication between goroutines (prevents race conditions)
- **GC** → Automatic memory management (but you still need to understand heap vs stack)
- **GMP Scheduler** → Go's implementation of M:N threading model

---

[← Back to Fundamentals](../) | [↑ Back to Main](../../README.md)
