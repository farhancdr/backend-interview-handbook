# Virtual Memory & Paging

## What is Virtual Memory?
Virtual Memory is an abstraction that gives each process the illusion of having its own **private, contiguous address space**, even though physical RAM is shared and fragmented.

### Why Virtual Memory?
1. **Isolation**: Process A cannot access Process B's memory (security).
2. **Simplicity**: Programs can be written assuming they have all the memory (no manual memory management).
3. **Overcommitment**: Total virtual memory across all processes can exceed physical RAM (swap to disk).

## Paging
The OS divides virtual memory into fixed-size **pages** (typically 4KB). Physical memory is divided into **frames** of the same size.

### Page Table
A **Page Table** maps virtual page numbers to physical frame numbers.

| Virtual Page | Physical Frame | Valid Bit | Dirty Bit |
| :--- | :--- | :--- | :--- |
| 0 | 5 | 1 | 0 |
| 1 | 2 | 1 | 1 |
| 2 | - | 0 | - |

- **Valid Bit**: 1 = page is in RAM. 0 = page fault (needs to be loaded from disk).
- **Dirty Bit**: 1 = page has been modified (must be written back to disk if evicted).

## Translation Lookaside Buffer (TLB)
Page table lookups are slow (require memory access). The **TLB** is a hardware cache of recent virtual-to-physical translations.

**TLB Hit**: Translation found in TLB → Fast (1 cycle).  
**TLB Miss**: Must walk the page table → Slow (~100 cycles).

### Context Switching Impact
When switching processes, the TLB is often flushed (because virtual addresses are process-specific). This is why thread switching is faster than process switching.

## Page Faults
A **Page Fault** occurs when a process accesses a virtual page that is not in physical memory.

### Types of Page Faults
1. **Minor Page Fault**: Page is in memory but not mapped (e.g., copy-on-write).
2. **Major Page Fault**: Page must be loaded from disk (swap).

### Page Replacement Algorithms
When RAM is full and a new page is needed, the OS must evict a page:
- **LRU (Least Recently Used)**: Evict the page that hasn't been used for the longest time.
- **FIFO (First In First Out)**: Evict the oldest page.
- **Clock Algorithm**: Approximation of LRU using a reference bit.

## Swapping
When physical memory is full, the OS moves inactive pages to **swap space** (disk).

**Thrashing**: When the system spends more time swapping pages than executing code (too many processes, not enough RAM).

## Memory-Mapped Files
A file can be mapped into a process's virtual address space using `mmap()`. Reading/writing to that memory region reads/writes the file.

**Benefits**:
- Efficient file I/O (no explicit read/write syscalls).
- Shared memory between processes (map the same file).

## Go Context
Go's runtime abstracts away most of this, but understanding virtual memory helps explain:
- **Escape Analysis**: Go decides whether to allocate on the stack (fast, no page faults) or heap (slower, may page fault).
- **GC Pressure**: Allocating too much on the heap increases GC work and potential paging.
- **Memory Limits**: `GOMEMLIMIT` controls Go's heap size, preventing excessive swapping.

### Example: Checking Page Faults in Go
```go
// Use runtime.ReadMemStats to see page faults
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("Page Faults: %d\n", m.Faults)
```
