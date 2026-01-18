# Caching Strategies

## Why Caching?
Caching stores frequently accessed data in a fast storage layer (RAM) to reduce:
1. **Latency**: RAM is ~100x faster than disk, ~1000x faster than network.
2. **Database Load**: Fewer queries to the database.
3. **Cost**: Reduce expensive database operations.

## Cache-Aside (Lazy Loading)

**Pattern**: Application checks the cache first. If miss, fetch from database and populate cache.

### Flow
1. **Read**: Check cache.
   - **Hit**: Return cached data.
   - **Miss**: Fetch from database, store in cache, return data.
2. **Write**: Update database, invalidate cache (or update cache).

### Example (Go + Redis)
```go
func GetUser(userID int) (*User, error) {
    // Check cache
    cacheKey := fmt.Sprintf("user:%d", userID)
    cached, err := rdb.Get(ctx, cacheKey).Result()
    if err == nil {
        var user User
        json.Unmarshal([]byte(cached), &user)
        return &user, nil // Cache hit
    }
    
    // Cache miss: fetch from database
    var user User
    db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name, &user.Email)
    
    // Store in cache
    data, _ := json.Marshal(user)
    rdb.Set(ctx, cacheKey, data, 10*time.Minute)
    
    return &user, nil
}
```

### Pros
- **Simple**: Easy to implement.
- **Lazy**: Only cache data that's actually requested.

### Cons
- **Cache Miss Penalty**: First request is slow (must fetch from database).
- **Stale Data**: Cache may be out of sync with database.

## Write-Through

**Pattern**: Write to cache and database **simultaneously**.

### Flow
1. **Write**: Update cache and database (both must succeed).
2. **Read**: Always read from cache (cache is always up-to-date).

### Example
```go
func UpdateUser(user *User) error {
    // Update database
    _, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
    if err != nil {
        return err
    }
    
    // Update cache
    cacheKey := fmt.Sprintf("user:%d", user.ID)
    data, _ := json.Marshal(user)
    rdb.Set(ctx, cacheKey, data, 10*time.Minute)
    
    return nil
}
```

### Pros
- **Consistency**: Cache is always up-to-date.
- **No Cache Miss Penalty**: Data is always in cache.

### Cons
- **Slower Writes**: Must write to both cache and database.
- **Wasted Cache**: May cache data that's never read.

## Write-Back (Write-Behind)

**Pattern**: Write to cache first, then **asynchronously** write to database.

### Flow
1. **Write**: Update cache immediately.
2. **Background**: Periodically flush cache to database.

### Example
```go
func UpdateUser(user *User) error {
    // Update cache immediately
    cacheKey := fmt.Sprintf("user:%d", user.ID)
    data, _ := json.Marshal(user)
    rdb.Set(ctx, cacheKey, data, 10*time.Minute)
    
    // Queue for async database write
    writeQueue <- user
    
    return nil
}

// Background worker
go func() {
    for user := range writeQueue {
        db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
    }
}()
```

### Pros
- **Fast Writes**: No database latency.
- **Batch Writes**: Can batch multiple writes to database.

### Cons
- **Data Loss Risk**: If cache crashes before flush, data is lost.
- **Complexity**: Requires background workers.

## Write-Around

**Pattern**: Write to database only (bypass cache). Invalidate cache on write.

### Flow
1. **Write**: Update database, invalidate cache.
2. **Read**: Cache-aside (check cache, fetch from database on miss).

### Example
```go
func UpdateUser(user *User) error {
    // Update database
    _, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
    if err != nil {
        return err
    }
    
    // Invalidate cache
    cacheKey := fmt.Sprintf("user:%d", user.ID)
    rdb.Del(ctx, cacheKey)
    
    return nil
}
```

### Pros
- **Simple**: No need to update cache on write.
- **Avoids Cache Pollution**: Don't cache data that's rarely read.

### Cons
- **Cache Miss Penalty**: Next read is slow (must fetch from database).

## Comparison Table

| Strategy | Write Speed | Read Speed | Consistency | Use Case |
| :--- | :--- | :--- | :--- | :--- |
| **Cache-Aside** | Fast (DB only) | Fast (cache hit) | **Eventual** | General-purpose |
| **Write-Through** | **Slow** (cache + DB) | **Very Fast** (always cached) | **Strong** | Read-heavy, critical data |
| **Write-Back** | **Very Fast** (cache only) | **Very Fast** (always cached) | **Weak** (risk of loss) | Write-heavy, can tolerate loss |
| **Write-Around** | Fast (DB only) | Slow (cache miss) | **Strong** | Write-heavy, infrequent reads |

## TTL (Time to Live) Strategies

### Fixed TTL
Set a fixed expiration time (e.g., 10 minutes).

```go
rdb.Set(ctx, "user:123", data, 10*time.Minute)
```

**Pros**: Simple.  
**Cons**: May serve stale data until TTL expires.

### Sliding TTL
Reset TTL on every access (keep frequently accessed data cached longer).

```go
rdb.Get(ctx, "user:123")
rdb.Expire(ctx, "user:123", 10*time.Minute) // Reset TTL
```

### No TTL (Manual Invalidation)
Cache forever, invalidate manually on updates.

```go
rdb.Set(ctx, "user:123", data, 0) // No expiration
// On update:
rdb.Del(ctx, "user:123")
```

## Cache Stampede Prevention

**Problem**: When a popular cached item expires, many requests simultaneously hit the database (thundering herd).

### Solutions

#### 1. Probabilistic Early Expiration
Randomly expire cache entries slightly before TTL.

```go
actualTTL := baseTTL - rand.Intn(60) // Expire 0-60 seconds early
rdb.Set(ctx, key, data, actualTTL)
```

#### 2. Locking (Single-Flight)
Only one request fetches from database, others wait.

```go
import "golang.org/x/sync/singleflight"

var group singleflight.Group

func GetUser(userID int) (*User, error) {
    key := fmt.Sprintf("user:%d", userID)
    
    // Only one request fetches from DB
    val, err, _ := group.Do(key, func() (interface{}, error) {
        // Check cache
        cached, err := rdb.Get(ctx, key).Result()
        if err == nil {
            var user User
            json.Unmarshal([]byte(cached), &user)
            return &user, nil
        }
        
        // Fetch from database
        var user User
        db.QueryRow("SELECT * FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name, &user.Email)
        
        // Cache it
        data, _ := json.Marshal(user)
        rdb.Set(ctx, key, data, 10*time.Minute)
        
        return &user, nil
    })
    
    return val.(*User), err
}
```

#### 3. Stale-While-Revalidate
Serve stale data while fetching fresh data in the background.

## Cache Invalidation

**"There are only two hard things in Computer Science: cache invalidation and naming things."** - Phil Karlton

### Strategies

#### 1. TTL-Based
Let cache expire automatically.

#### 2. Event-Based
Invalidate cache when data changes.

```go
func UpdateUser(user *User) error {
    db.Exec("UPDATE users SET name = $1 WHERE id = $2", user.Name, user.ID)
    rdb.Del(ctx, fmt.Sprintf("user:%d", user.ID)) // Invalidate
    return nil
}
```

#### 3. Tag-Based
Group related cache entries and invalidate by tag.

```go
// Cache user and their posts
rdb.Set(ctx, "user:123", userData, 10*time.Minute)
rdb.SAdd(ctx, "tag:user:123", "user:123", "posts:123")

// Invalidate all related data
members, _ := rdb.SMembers(ctx, "tag:user:123").Result()
for _, key := range members {
    rdb.Del(ctx, key)
}
```

## Interview Questions

### Q: What's the difference between Cache-Aside and Write-Through?
**A**: 
- **Cache-Aside**: Application manages cache (check cache, fetch from DB on miss, populate cache).
- **Write-Through**: Write to cache and database simultaneously (cache is always up-to-date).

### Q: What is cache stampede and how do you prevent it?
**A**: When a popular cached item expires, many requests hit the database simultaneously. **Fix**: Use **singleflight** (only one request fetches, others wait) or **probabilistic early expiration** (spread out expirations).

### Q: When would you use Write-Back caching?
**A**: Write-heavy workloads where you can tolerate data loss (e.g., analytics, logs). Writes are very fast (cache only), but risk losing data if cache crashes before flushing to database.

### Q: What are the two hard things in computer science?
**A**: **Cache invalidation** and **naming things** (and off-by-one errors).
