# File Systems

## What is a File System?
A **file system** organizes and stores data on disk. It provides:
1. **Naming**: Files have human-readable names.
2. **Hierarchy**: Directories organize files.
3. **Metadata**: File size, permissions, timestamps.
4. **Persistence**: Data survives power loss.

## Inodes (Index Nodes)

An **inode** is a data structure that stores metadata about a file (but not the filename or data).

### Inode Contents
- **File size**
- **Owner (UID, GID)**
- **Permissions** (read, write, execute)
- **Timestamps** (created, modified, accessed)
- **Pointers to data blocks** (where the actual file content is stored)

### Key Insight
**Filenames are stored in directories, not inodes.** A directory is just a file that maps filenames to inode numbers.

### Example: `ls -i` (Show Inode Numbers)
```bash
$ ls -i
12345 file.txt
12346 dir/
```

## Hard Links vs Soft Links (Symbolic Links)

| Feature | Hard Link | Soft Link (Symlink) |
| :--- | :--- | :--- |
| **Definition** | Another name for the same inode. | A special file that points to another file's path. |
| **Inode** | Same inode as original. | Different inode (stores the target path). |
| **Deletion** | File persists until all hard links are deleted. | If target is deleted, symlink becomes "dangling". |
| **Cross-Filesystem** | No (inode numbers are filesystem-specific). | Yes. |
| **Directories** | Not allowed (would create cycles). | Allowed. |

### Example: Hard Link
```bash
$ echo "hello" > original.txt
$ ln original.txt hardlink.txt
$ ls -i
12345 original.txt
12345 hardlink.txt  # Same inode!
$ rm original.txt
$ cat hardlink.txt  # Still works!
hello
```

### Example: Soft Link
```bash
$ ln -s original.txt symlink.txt
$ ls -l
lrwxrwxrwx 1 user user 12 Jan 1 12:00 symlink.txt -> original.txt
$ rm original.txt
$ cat symlink.txt  # Error: No such file or directory
```

## File System Operations

### 1. Creating a File
1. Allocate a new **inode**.
2. Allocate **data blocks** for the file content.
3. Add an entry in the parent directory (filename → inode number).

### 2. Reading a File
1. Look up the filename in the directory to get the inode number.
2. Read the inode to get the data block pointers.
3. Read the data blocks.

### 3. Deleting a File
1. Remove the directory entry (filename → inode).
2. Decrement the inode's **link count**.
3. If link count reaches 0, free the inode and data blocks.

## Journaling File Systems

**Problem**: If the system crashes during a file operation (e.g., write), the file system can become **inconsistent** (metadata doesn't match data).

**Solution**: **Journaling** - Write changes to a **journal** (log) before applying them to the file system.

### How It Works
1. **Write** the operation to the journal (e.g., "create file X").
2. **Commit** the journal entry (mark it as complete).
3. **Apply** the operation to the file system.
4. **Delete** the journal entry.

If the system crashes:
- **Before commit**: Ignore the incomplete journal entry.
- **After commit**: Replay the journal entry on reboot.

### Types of Journaling
- **Metadata Journaling** (ext3, ext4 default): Only journal metadata (fast, but data may be lost).
- **Full Journaling** (ext4 with `data=journal`): Journal both metadata and data (slower, but safer).

### Examples
- **ext4** (Linux): Journaling file system.
- **NTFS** (Windows): Journaling file system.
- **APFS** (macOS): Copy-on-write (similar to journaling).

## File System Performance

### 1. Disk Seek Time
**Problem**: Mechanical hard drives have slow seek times (~5-10ms).

**Solution**: 
- **Contiguous Allocation**: Store file blocks sequentially (reduces seeks).
- **SSD**: No mechanical parts (seek time ~0.1ms).

### 2. Caching
The OS caches frequently accessed file data in RAM (**page cache**).

**Read**: Check cache first. If miss, read from disk and cache it.  
**Write**: Write to cache first (dirty page). Flush to disk later (write-back).

### 3. Prefetching
When reading a file sequentially, the OS prefetches the next blocks into the cache.

## Go Context

### 1. File Operations in Go
```go
// Open file (syscall: open)
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// Read file (syscall: read)
data, err := io.ReadAll(file)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

### 2. Hard Link in Go
```go
err := os.Link("original.txt", "hardlink.txt")
if err != nil {
    log.Fatal(err)
}
```

### 3. Soft Link in Go
```go
err := os.Symlink("original.txt", "symlink.txt")
if err != nil {
    log.Fatal(err)
}
```

### 4. File Metadata (Inode Info)
```go
info, err := os.Stat("file.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Size:", info.Size())
fmt.Println("Mode:", info.Mode())
fmt.Println("ModTime:", info.ModTime())
```

## Interview Questions

### Q: What happens when you delete a file that's still open by a process?
**A**: The directory entry is removed, but the inode and data blocks are **not freed** until the process closes the file (link count includes open file descriptors).

### Q: Why can't you create hard links to directories?
**A**: It would create cycles in the directory tree, making it impossible to determine when to free an inode (link count would never reach 0).

### Q: How does journaling improve reliability?
**A**: By logging operations before applying them, the file system can recover from crashes by replaying the journal (ensures consistency).
