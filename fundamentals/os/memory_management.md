# Memory Management: Stack vs Heap

## Stack vs Heap

| Feature | Stack | Heap |
| :--- | :--- | :--- |
| **Allocation** | Automatic (function call pushes frame). | Manual (malloc/new) or GC-managed. |
| **Deallocation** | Automatic (function return pops frame). | Manual (free/delete) or GC-managed. |
| **Speed** | **Very Fast** (just move stack pointer). | Slower (requires finding free block). |
| **Size** | Fixed per thread (~1-8MB). | Limited by virtual memory (GBs). |
| **Lifetime** | Scoped to function call. | Persists until explicitly freed or GC'd. |
| **Fragmentation** | None (LIFO structure). | **Yes** (external/internal fragmentation). |
| **Thread Safety** | Each thread has its own stack. | Shared across threads (needs synchronization). |

## Stack Memory

### How It Works
Each function call pushes a **stack frame** containing:
1. **Local variables**.
2. **Return address** (where to jump back after function ends).
3. **Saved registers**.

When the function returns, the frame is popped.

### Stack Overflow
Occurs when the stack grows beyond its limit (e.g., infinite recursion).

```go
func recurse() {
    recurse() // Stack overflow!
}
```

**Go**: Default stack size is **2KB** (grows dynamically up to ~1GB on 64-bit systems).

## Heap Memory

### How It Works
The heap is a large pool of memory managed by the allocator (or GC in Go).

**Allocation**: Find a free block of sufficient size.  
**Deallocation**: Mark the block as free (or let GC reclaim it).

### Fragmentation

#### External Fragmentation
Free memory is scattered in small chunks, so a large allocation fails even though total free memory is sufficient.

**Example**: Free blocks of 10KB, 5KB, 10KB. Request for 20KB fails.

**Fix**: **Compaction** (move allocated blocks together) or use **Buddy Allocator**.

#### Internal Fragmentation
Allocated block is larger than requested (wasted space).

**Example**: Request 10 bytes, allocator gives 16 bytes (due to alignment).

## Memory Allocators

### 1. Free List Allocator
Maintains a linked list of free blocks.

**Strategies**:
- **First Fit**: Use the first block that fits.
- **Best Fit**: Use the smallest block that fits (minimizes waste).
- **Worst Fit**: Use the largest block (leaves large free blocks).

### 2. Buddy Allocator
Splits memory into power-of-2 sized blocks. When freeing, merges adjacent "buddy" blocks.

**Pros**: Fast allocation/deallocation. Reduces external fragmentation.  
**Cons**: Internal fragmentation (always allocates power-of-2 sizes).

### 3. Slab Allocator
Pre-allocates fixed-size chunks for frequently allocated objects (e.g., file descriptors).

**Pros**: Very fast. No fragmentation for fixed-size objects.  
**Used by**: Linux kernel, Go's runtime (for small objects).

## Garbage Collection Basics

### Mark and Sweep
1. **Mark**: Starting from roots (globals, stack), mark all reachable objects.
2. **Sweep**: Free all unmarked objects.

**Drawback**: "Stop-the-world" pauses (all threads stop during GC).

### Generational GC
**Observation**: Most objects die young.

**Strategy**: Divide heap into generations (young, old). GC young generation frequently, old generation rarely.

### Go's GC
- **Concurrent Mark-Sweep**: GC runs concurrently with application (minimal pauses).
- **Tri-color Marking**: White (unvisited), Gray (visited, children not scanned), Black (visited, children scanned).
- **Write Barrier**: Ensures correctness when mutator (application) modifies pointers during GC.

## Escape Analysis in Go

Go's compiler decides whether a variable should be allocated on the **stack** or **heap**.

### Rules
- **Stack**: Variable's lifetime is scoped to the function.
- **Heap**: Variable "escapes" (e.g., returned pointer, stored in global, sent to channel).

### Example
```go
func stackAlloc() int {
    x := 42 // Allocated on stack
    return x
}

func heapAlloc() *int {
    x := 42 // Escapes to heap (pointer returned)
    return &x
}
```

**Check Escape Analysis**:
```bash
go build -gcflags="-m" main.go
```

## Interview Questions

### Q: Why is stack allocation faster than heap allocation?
**A**: Stack allocation is just incrementing the stack pointer (1 instruction). Heap allocation requires finding a free block, updating metadata, and potentially triggering GC.

### Q: What causes memory leaks in Go (despite having GC)?
**A**: 
1. **Goroutine leaks** (goroutines blocked forever hold references).
2. **Global variables** holding large slices/maps.
3. **Unclosed resources** (file handles, network connections).

### Q: How does Go's GC differ from Java's?
**A**: Go uses **concurrent mark-sweep** (low-latency, short pauses). Java uses **generational GC** (higher throughput, longer pauses). Go prioritizes **latency**, Java prioritizes **throughput**.
