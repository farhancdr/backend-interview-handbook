# Load Balancing

## What is Load Balancing?
Distributing incoming network traffic across multiple servers to ensure:
1. **High Availability**: If one server fails, others handle the load.
2. **Scalability**: Add more servers to handle increased traffic.
3. **Performance**: No single server is overwhelmed.

## L4 vs L7 Load Balancing

| Feature | L4 (Transport Layer) | L7 (Application Layer) |
| :--- | :--- | :--- |
| **OSI Layer** | Layer 4 (TCP/UDP) | Layer 7 (HTTP/HTTPS) |
| **Routing Decision** | Based on IP address and port. | Based on HTTP headers, URL path, cookies. |
| **Speed** | **Faster** (less inspection). | Slower (must parse HTTP). |
| **SSL Termination** | No (passes through). | Yes (can decrypt and inspect). |
| **Use Case** | Generic TCP/UDP traffic (databases, game servers). | Web applications (route `/api` to API servers, `/static` to CDN). |
| **Examples** | AWS NLB, HAProxy (TCP mode). | AWS ALB, Nginx, HAProxy (HTTP mode). |

### L4 Example
```
Client → Load Balancer (sees IP:Port) → Server A or B
```

### L7 Example
```
Client → Load Balancer (sees HTTP path /api/users) → API Server
Client → Load Balancer (sees HTTP path /static/image.png) → CDN
```

## Load Balancing Algorithms

### 1. Round Robin
Distribute requests sequentially to each server in order.

**Pros**: Simple. Fair if all servers are equal.  
**Cons**: Doesn't account for server load or capacity.

### 2. Least Connections
Send requests to the server with the fewest active connections.

**Pros**: Better for long-lived connections (WebSockets, database connections).  
**Cons**: Requires tracking connection counts.

### 3. IP Hash
Hash the client's IP address to determine which server to use.

**Pros**: **Sticky sessions** (same client always goes to the same server).  
**Cons**: Uneven distribution if clients are behind NAT (many clients share one IP).

### 4. Weighted Round Robin
Assign weights to servers based on capacity (e.g., Server A gets 70% of traffic, Server B gets 30%).

**Use Case**: Servers with different hardware specs.

### 5. Consistent Hashing
Hash both the client and servers onto a ring. Route client to the nearest server.

**Pros**: Minimal disruption when adding/removing servers (only 1/N keys are remapped).  
**Use Case**: Distributed caches (Redis, Memcached).

## Health Checks
Load balancers periodically check if servers are healthy.

### Types
1. **Passive**: Mark server as unhealthy after N failed requests.
2. **Active**: Send periodic health check requests (e.g., `GET /health` every 5 seconds).

### Example: Nginx Health Check
```nginx
upstream backend {
    server backend1.example.com;
    server backend2.example.com;
    
    # Active health check
    check interval=3000 rise=2 fall=3 timeout=1000;
}
```

## Sticky Sessions (Session Affinity)
Ensure a client always connects to the same server (important if session data is stored locally on the server).

### Methods
1. **IP Hash**: Hash client IP.
2. **Cookie-Based**: Load balancer sets a cookie with the server ID.

**Drawback**: Reduces load balancing effectiveness (one server may get overloaded).

**Better Solution**: Use a **shared session store** (Redis, database) so any server can handle any client.

## Failover
If a server fails, the load balancer stops routing traffic to it.

### Active-Passive Failover
- **Active**: Primary server handles all traffic.
- **Passive**: Standby server takes over if primary fails.

### Active-Active Failover
- All servers handle traffic simultaneously.
- If one fails, others absorb the load.

## Go Context: Building a Simple Load Balancer

### Round Robin Load Balancer
```go
package main

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "net/url"
    "sync"
)

type LoadBalancer struct {
    servers []*url.URL
    current int
    mu      sync.Mutex
}

func (lb *LoadBalancer) NextServer() *url.URL {
    lb.mu.Lock()
    defer lb.mu.Unlock()
    server := lb.servers[lb.current]
    lb.current = (lb.current + 1) % len(lb.servers)
    return server
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    target := lb.NextServer()
    proxy := httputil.NewSingleHostReverseProxy(target)
    proxy.ServeHTTP(w, r)
}

func main() {
    servers := []*url.URL{
        {Scheme: "http", Host: "localhost:8001"},
        {Scheme: "http", Host: "localhost:8002"},
    }
    lb := &LoadBalancer{servers: servers}
    http.ListenAndServe(":8000", lb)
}
```

## Interview Questions

### Q: When would you use L4 vs L7 load balancing?
**A**: 
- **L4**: When you need maximum performance and don't need to inspect application-layer data (e.g., load balancing database connections).
- **L7**: When you need content-based routing (e.g., route `/api` to API servers, `/static` to CDN).

### Q: How does consistent hashing help with caching?
**A**: When a cache server is added/removed, only ~1/N keys are remapped (instead of all keys with traditional hashing). This minimizes cache invalidation.

### Q: What's the problem with sticky sessions?
**A**: It reduces load balancing effectiveness (one server may get overloaded) and makes horizontal scaling harder (can't easily add/remove servers).
