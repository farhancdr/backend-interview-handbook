# CDN & Edge Computing

## What is a CDN (Content Delivery Network)?
A **CDN** is a geographically distributed network of servers that cache and serve content (images, videos, CSS, JS) from locations close to users.

### Why CDN?
1. **Reduce Latency**: Serve content from the nearest edge server (e.g., US users get content from a US server, not from Asia).
2. **Reduce Load**: Offload traffic from the origin server.
3. **Improve Availability**: If the origin server is down, the CDN can still serve cached content.
4. **DDoS Protection**: CDN absorbs malicious traffic.

## How CDN Works

### Flow
1. **User** requests `https://example.com/image.png`.
2. **DNS** resolves `example.com` to the **nearest CDN edge server** (using GeoDNS).
3. **Edge Server** checks its cache:
   - **Cache Hit**: Serve the cached image.
   - **Cache Miss**: Fetch from the **origin server**, cache it, and serve it.
4. **Subsequent requests** for the same image are served from the cache.

### Example: Cloudflare CDN
```
User (New York) → Cloudflare Edge (New York) → Origin Server (California)
User (London) → Cloudflare Edge (London) → Origin Server (California)
```

## Cache Invalidation Strategies

### 1. Time-Based (TTL)
Set a **Time to Live (TTL)** for cached content. After TTL expires, the edge server fetches fresh content.

**Example**: `Cache-Control: max-age=3600` (cache for 1 hour).

**Pros**: Simple.  
**Cons**: Stale content may be served until TTL expires.

### 2. Purge/Invalidate
Manually invalidate cached content when it changes.

**Example**: After deploying a new version, purge `/static/app.js` from the CDN.

**Pros**: Immediate updates.  
**Cons**: Requires manual intervention or automation.

### 3. Versioned URLs
Append a version or hash to the URL (e.g., `/static/app.v2.js` or `/static/app.abc123.js`).

**Pros**: No cache invalidation needed (new URL = new cache entry).  
**Cons**: Requires build-time URL rewriting.

### 4. Stale-While-Revalidate
Serve stale content while fetching fresh content in the background.

**Example**: `Cache-Control: max-age=3600, stale-while-revalidate=86400`

**Pros**: Always fast (serve from cache) + eventually consistent.

## Edge Computing

**Edge Computing** = Running code at the CDN edge (close to users), not just caching static files.

### Use Cases
1. **A/B Testing**: Route users to different versions based on cookies.
2. **Authentication**: Verify JWT tokens at the edge (no need to hit the origin server).
3. **Geo-Blocking**: Block requests from certain countries.
4. **Image Resizing**: Resize images on-the-fly based on device (mobile vs desktop).
5. **API Gateway**: Rate limiting, request transformation.

### Examples
- **Cloudflare Workers**: Run JavaScript at the edge.
- **AWS Lambda@Edge**: Run Node.js/Python at CloudFront edge locations.
- **Fastly Compute@Edge**: Run WebAssembly at the edge.

### Example: Cloudflare Worker (JavaScript)
```javascript
addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request))
})

async function handleRequest(request) {
  const url = new URL(request.url)
  
  // Redirect /old to /new
  if (url.pathname === '/old') {
    return Response.redirect('https://example.com/new', 301)
  }
  
  // Fetch from origin
  return fetch(request)
}
```

## Geo-Routing

**Geo-Routing** = Route users to the nearest server based on their geographic location.

### How It Works
1. **GeoDNS**: DNS returns different IP addresses based on the client's location.
2. **Anycast**: Multiple servers share the same IP address. Routers send traffic to the nearest server.

### Example: Anycast
```
IP: 1.1.1.1 (Cloudflare DNS)
- Server in New York
- Server in London
- Server in Tokyo

User in New York → Routed to New York server
User in Tokyo → Routed to Tokyo server
```

## Cache Stampede Prevention

**Problem**: When a popular cached item expires, many requests simultaneously hit the origin server (cache stampede).

### Solutions

#### 1. Stale-While-Revalidate
Serve stale content while one request fetches fresh content.

#### 2. Request Coalescing
If multiple requests for the same resource arrive, only one request goes to the origin. Others wait for the result.

#### 3. Probabilistic Early Expiration
Randomly expire cache entries slightly before TTL to spread out origin requests.

## Go Context: Building a Simple CDN Cache

### In-Memory Cache with TTL
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type CacheItem struct {
    Value      string
    Expiration time.Time
}

type Cache struct {
    items map[string]CacheItem
    mu    sync.RWMutex
}

func NewCache() *Cache {
    return &Cache{items: make(map[string]CacheItem)}
}

func (c *Cache) Set(key, value string, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.items[key] = CacheItem{
        Value:      value,
        Expiration: time.Now().Add(ttl),
    }
}

func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    item, found := c.items[key]
    if !found || time.Now().After(item.Expiration) {
        return "", false
    }
    return item.Value, true
}

func main() {
    cache := NewCache()
    cache.Set("user:123", "Alice", 5*time.Second)
    
    if val, found := cache.Get("user:123"); found {
        fmt.Println("Cache hit:", val)
    } else {
        fmt.Println("Cache miss")
    }
    
    time.Sleep(6 * time.Second)
    if val, found := cache.Get("user:123"); found {
        fmt.Println("Cache hit:", val)
    } else {
        fmt.Println("Cache miss (expired)")
    }
}
```

## Interview Questions

### Q: How does a CDN reduce latency?
**A**: By serving content from edge servers geographically close to users (reduces network round-trip time).

### Q: What's the difference between CDN and Edge Computing?
**A**: 
- **CDN**: Caches static content (images, videos, CSS).
- **Edge Computing**: Runs code at the edge (authentication, A/B testing, image resizing).

### Q: How do you invalidate cached content in a CDN?
**A**: 
1. **TTL**: Set expiration time.
2. **Purge**: Manually invalidate.
3. **Versioned URLs**: Use unique URLs for each version (e.g., `/app.v2.js`).

### Q: What is cache stampede and how do you prevent it?
**A**: When a popular cached item expires, many requests hit the origin simultaneously. **Fix**: Use **stale-while-revalidate** (serve stale content while fetching fresh) or **request coalescing** (only one request fetches, others wait).
