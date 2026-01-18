# CPU Scheduling Algorithms

## The Problem
Multiple processes/threads want CPU time. The **scheduler** decides which one runs next.

## Goals
1. **Fairness**: Every process gets CPU time.
2. **Throughput**: Maximize number of processes completed per unit time.
3. **Turnaround Time**: Minimize time from submission to completion.
4. **Response Time**: Minimize time from request to first response (important for interactive systems).
5. **CPU Utilization**: Keep CPU busy.

## Preemptive vs Non-Preemptive

| Type | Description | Example |
| :--- | :--- | :--- |
| **Non-Preemptive** | Once a process starts, it runs until it voluntarily yields (blocks on I/O or finishes). | FCFS, Cooperative Multitasking |
| **Preemptive** | OS can forcibly stop a running process (timer interrupt) and switch to another. | Round Robin, Priority, CFS |

## Common Scheduling Algorithms

### 1. First-Come, First-Served (FCFS)
**Non-Preemptive**. Processes run in the order they arrive.

**Pros**: Simple.  
**Cons**: **Convoy Effect** - A long process blocks all short processes behind it.

### 2. Shortest Job First (SJF)
Run the process with the shortest expected CPU burst next.

**Pros**: Optimal average turnaround time.  
**Cons**: Requires knowing future CPU burst times (impossible). Can cause **starvation** (long jobs never run).

### 3. Round Robin (RR)
**Preemptive**. Each process gets a fixed **time quantum** (e.g., 10ms). After the quantum expires, the process is preempted and moved to the back of the queue.

**Pros**: Fair. Good response time.  
**Cons**: High context switching overhead if quantum is too small. Poor turnaround time if quantum is too large.

### 4. Priority Scheduling
Each process has a priority. Highest priority runs first.

**Pros**: Important processes run first.  
**Cons**: **Starvation** - Low-priority processes may never run.  
**Fix**: **Aging** - Gradually increase the priority of waiting processes.

### 5. Multi-Level Queue
Processes are divided into groups (e.g., foreground interactive, background batch). Each group has its own queue and scheduling algorithm.

### 6. Completely Fair Scheduler (CFS) - Linux Default
Tracks the **virtual runtime** (vruntime) of each process. The process with the lowest vruntime runs next.

**Key Idea**: Every process should get an equal share of CPU time over a period.

**Data Structure**: Red-Black Tree (O(log n) insertion/deletion).

## Context Switching Cost
Switching from Process A to Process B involves:
1. **Save** A's registers, program counter, stack pointer.
2. **Load** B's saved state.
3. **Flush TLB** (if different address space).
4. **Invalidate CPU caches** (partially).

**Cost**: ~1-10 microseconds (depends on architecture).

## Go Runtime Scheduler (M:N Model)
Go doesn't use OS threads directly. Instead, it uses a **work-stealing scheduler** that maps **M goroutines** onto **N OS threads**.

### Key Components
- **G (Goroutine)**: Lightweight thread (~2KB stack).
- **M (Machine)**: OS thread.
- **P (Processor)**: Scheduling context (local run queue).

### Scheduling Strategy
1. Each P has a local run queue of goroutines.
2. When a goroutine blocks (e.g., channel receive), the M detaches and picks up another goroutine.
3. **Work Stealing**: If P's queue is empty, it steals goroutines from other P's queues.

### Preemption in Go
- **Before Go 1.14**: Cooperative (goroutines yielded at function calls).
- **Go 1.14+**: **Asynchronous Preemption** - Runtime can preempt goroutines even in tight loops (prevents GC stalls).

## Interview Question: Why is Go's Scheduler Better?
**Answer**: OS schedulers are designed for processes/threads (heavyweight, 1-2MB stack). Go's scheduler is designed for goroutines (lightweight, 2KB stack). Go can schedule **millions** of goroutines efficiently because:
1. **No syscalls** for scheduling (all in userspace).
2. **Work stealing** balances load automatically.
3. **Integrated with GC** (can pause goroutines for GC without OS involvement).
