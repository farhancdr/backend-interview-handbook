# System Calls & Interrupts

## What is a System Call?
A **system call** is the interface between a user program and the operating system kernel. It allows programs to request services like file I/O, process creation, and network communication.

## User Mode vs Kernel Mode

| Mode | Description | Privileges | Examples |
| :--- | :--- | :--- | :--- |
| **User Mode** | Normal application code runs here. | Restricted (cannot access hardware directly). | Your Go program. |
| **Kernel Mode** | OS kernel runs here. | Full access to hardware, memory, I/O. | System call handlers, device drivers. |

**Why Two Modes?**: **Security and Stability**. If user programs could access hardware directly, a bug could crash the entire system.

## How System Calls Work

### The Flow
1. **User program** calls a library function (e.g., `read()`).
2. **Library** sets up arguments and invokes a **trap instruction** (e.g., `syscall` on x86-64).
3. **CPU switches to kernel mode** and jumps to the system call handler.
4. **Kernel** executes the requested operation (e.g., read from disk).
5. **Kernel returns** to user mode and resumes the program.

### Cost
System calls are **expensive** (~100-1000 CPU cycles) due to:
- **Mode switch** (user → kernel → user).
- **Context saving/restoring** (registers, stack).
- **TLB flush** (if address space changes).

**Optimization**: Batch operations (e.g., `writev()` instead of multiple `write()` calls).

## Common System Calls

### File Operations
- `open(path, flags)`: Open a file, returns file descriptor.
- `read(fd, buffer, size)`: Read from file.
- `write(fd, buffer, size)`: Write to file.
- `close(fd)`: Close file descriptor.

### Process Management
- `fork()`: Create a new process (copy of current process).
- `exec(path, args)`: Replace current process with a new program.
- `wait(pid)`: Wait for child process to terminate.
- `exit(status)`: Terminate the current process.

### Memory Management
- `mmap(addr, length, prot, flags, fd, offset)`: Map file or device into memory.
- `munmap(addr, length)`: Unmap memory.
- `brk(addr)`: Change the size of the heap.

### Networking
- `socket(domain, type, protocol)`: Create a socket.
- `bind(sockfd, addr, addrlen)`: Bind socket to address.
- `listen(sockfd, backlog)`: Listen for connections.
- `accept(sockfd, addr, addrlen)`: Accept a connection.

## Interrupts vs Traps

| Type | Trigger | Purpose | Example |
| :--- | :--- | :--- | :--- |
| **Interrupt** | **Asynchronous** (external event). | Notify CPU of hardware events. | Keyboard press, network packet arrival. |
| **Trap** | **Synchronous** (program instruction). | Request OS service. | System call, division by zero. |

### Interrupt Handling
1. **Hardware** sends interrupt signal to CPU.
2. **CPU** saves current state and jumps to **Interrupt Service Routine (ISR)**.
3. **ISR** handles the interrupt (e.g., read keyboard buffer).
4. **CPU** resumes the interrupted program.

## Go's syscall Package

Go provides the `syscall` package for low-level OS interactions.

### Example: Open a File
```go
package main

import (
    "fmt"
    "syscall"
)

func main() {
    fd, err := syscall.Open("/tmp/test.txt", syscall.O_RDONLY, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer syscall.Close(fd)

    buf := make([]byte, 100)
    n, err := syscall.Read(fd, buf)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
}
```

### Example: Fork (Unix-like systems)
```go
pid, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
if pid == 0 {
    fmt.Println("Child process")
} else {
    fmt.Println("Parent process, child PID:", pid)
}
```

**Note**: Go's runtime doesn't support `fork()` well (goroutines don't survive fork). Use `os/exec` instead.

## Why Go Abstracts System Calls

Go's standard library (`os`, `net`, `io`) wraps system calls for:
1. **Portability**: Same code works on Linux, Windows, macOS.
2. **Safety**: Prevents common errors (e.g., forgetting to close file descriptors).
3. **Concurrency**: Integrates with Go's scheduler (non-blocking I/O).

### Example: Blocking vs Non-Blocking I/O
**System Call**: `read()` blocks the OS thread until data is available.  
**Go**: Uses **non-blocking I/O** + **epoll/kqueue** (event loop). When `read()` would block, Go parks the goroutine and runs another.

## Interview Questions

### Q: What happens when you call `os.Open()` in Go?
**A**: 
1. Go calls `syscall.Open()` (which invokes the `open` system call).
2. CPU switches to kernel mode.
3. Kernel opens the file and returns a file descriptor.
4. Go wraps the file descriptor in an `os.File` struct.

### Q: Why are system calls slow?
**A**: Mode switching (user → kernel → user), context saving/restoring, and potential TLB flushes.

### Q: How does Go achieve high concurrency despite blocking system calls?
**A**: Go uses **non-blocking I/O** (epoll/kqueue) and parks goroutines when they would block, allowing other goroutines to run on the same OS thread.
