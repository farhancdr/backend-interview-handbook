# Project Overview

## Philosophy
The **Backend Interview Handbook** is designed to be the ultimate executable reference for backend engineers. It rejects the "passive reading" model of typical interview prep. Instead, it embraces "executable learning" through Go:

> **If you can't run it/test it/break it, you don't know it.**

Every concept—from basic slice internals to complex distributed system patterns—is implemented as a runnable test case. This allows users to verify their understanding by running code, not just reading text.

## Core Pillars

1.  **Isolation**: Each package is independent. You can jump into `internal/concurrency` without reading `internal/basics` first.
2.  **Executability**: `go test` is the primary interface. Documentation lives in code comments.
3.  **Interview Relevance**: We prioritize topics that actually come up in backend interviews (Context, Goroutine leaks, System Design) over academic trivia.
4.  **Zero Dependencies**: We use only the Go standard library to prove that you don't need frameworks to build robust systems.

## Project Structure

The project is organized into `internal/` packages to simulate a real Go project layout, organized by domain:

### 🟢 Foundation Layers
*   **`internal/basics`**: The "must-knows". Arrays, Slices, Maps, Structs. Pitfalls like `nil` slice vs empty slice.
*   **`internal/intermediate`**: Interfaces, Defer mechanics, Error wrapping.
*   **`internal/advanced`**: Context, Generics, Reflection, Unsafe pointer usage (Zero-Copy), Memory Alignment optimization.

### 🔵 Concurrency & Runtime
*   **`internal/concurrency`**: The crown jewel. Channels, Select, Worker Pools, Context timeouts.
*   **`internal/memory`**: Memory layout, Escape Analysis examples, Benchmarking patterns.
*   **`internal/internals`**: Deep dives into how Go works under the hood (GMP scheduler, GC).

### 🟡 Backend Engineering
*   **`internal/patterns`**: Architecture patterns. Repository, Service, Middleware, Options, Dependency Injection.
*   **`internal/system_design`**: Distributed system primitives. Rate Limiter, Cache, Pub-Sub, Idempotency.

### 🔴 DSA (Data Structures & Algorithms)
*   **`internal/ds`**: Production-quality implementations of Stack, Queue, Heap, BST, LRU Cache.
*   **`internal/algo`**: Interview staples. Binary Search, DFS/BFS, Sliding Window, Backtracking.
*   **`internal/leetcode`**: **(NEW)** A sandbox for practicing LeetCode-style problems with pre-written tests to verify your solutions.

### 🟣 CS Theory (Fundamentals)
*   **`fundamentals/os`**: Process vs Thread, Concurrency Models, Virtual Memory, CPU Scheduling, Deadlock, Memory Management, System Calls, File Systems.
*   **`fundamentals/networking`**: TCP vs UDP, HTTP Evolution, Load Balancing (L4/L7), DNS, WebSockets/SSE, TLS/SSL, gRPC/GraphQL, CDN, API Gateway.
*   **`fundamentals/database`**: ACID, Isolation Levels, Indexing (B-Tree vs LSM), CAP, Replication, Locking, Query Optimization, NoSQL vs SQL, Connection Pooling, Caching Strategies, Distributed Transactions.

## Workflow for Users

1.  **Select a Topic**: Pick a package based on study needs (e.g., `concurrency`).
2.  **Run the Test**: `go test -v ./internal/concurrency/`.
3.  **Read the Source**: Open the `.go` file. Read the "Why interviewers ask this" header.
4.  **Experiment**: Change a buffer size, remove a lock, or modify a timeout. Run the test again to see it fail.
5.  **Internalize**: The failure confirms your understanding of *why* the code was written that way.

## Target Audience
*   **Mid-Senior Engineers**: Refreshing on internals and system design patterns.
*   **Junior Engineers**: Moving beyond syntax to "idiomatic Go".
*   **Interviewers**: Looking for good questions and "what to look for" in candidate answers.
