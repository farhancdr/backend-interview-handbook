# API Gateway Patterns

## What is an API Gateway?
An **API Gateway** is a server that acts as a single entry point for all client requests to backend services. It handles:
1. **Routing**: Forward requests to the appropriate microservice.
2. **Authentication/Authorization**: Verify JWT tokens, API keys.
3. **Rate Limiting**: Prevent abuse.
4. **Request/Response Transformation**: Modify headers, body.
5. **Load Balancing**: Distribute traffic across service instances.
6. **Caching**: Cache responses to reduce backend load.
7. **Logging/Monitoring**: Centralized logging and metrics.

## Why API Gateway?

### Without API Gateway
```
Client → Service A (Auth)
Client → Service B (Users)
Client → Service C (Orders)
```

**Problems**:
- **Cross-Cutting Concerns**: Each service must implement auth, rate limiting, logging.
- **Chatty Clients**: Mobile apps make many requests (high latency).
- **Tight Coupling**: Clients know about all services.

### With API Gateway
```
Client → API Gateway → Service A/B/C
```

**Benefits**:
- **Single Entry Point**: Clients only know about the gateway.
- **Centralized Logic**: Auth, rate limiting, logging in one place.
- **Request Aggregation**: Gateway can combine multiple service calls into one response.

## API Gateway Patterns

### 1. Routing
Forward requests to the appropriate service based on the URL path.

**Example**:
```
GET /api/users → User Service
GET /api/orders → Order Service
GET /api/products → Product Service
```

### 2. Authentication & Authorization
Verify JWT tokens or API keys before forwarding requests.

**Example**:
```go
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if !isValidToken(token) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

### 3. Rate Limiting
Limit the number of requests per client (by IP, API key, user ID).

**Example**: Token Bucket algorithm (see `internal/system_design/rate_limiter.go`).

### 4. Request Aggregation (Backend for Frontend - BFF)
Combine multiple backend calls into a single response.

**Example**: Mobile app requests user profile + recent orders in one call.

**Without BFF**:
```
Client → GET /users/123 → User Service
Client → GET /orders?user=123 → Order Service
```

**With BFF**:
```
Client → GET /mobile/profile/123 → API Gateway
  → GET /users/123 → User Service
  → GET /orders?user=123 → Order Service
  → Combine responses → Client
```

### 5. Response Caching
Cache responses to reduce backend load.

**Example**: Cache `GET /products` for 5 minutes.

### 6. Circuit Breaker
If a backend service is down, stop sending requests to it (fail fast).

**Example**: After 5 consecutive failures, open the circuit (return error immediately). After 30 seconds, try again (half-open).

### 7. Request/Response Transformation
Modify headers, body, or format (e.g., XML → JSON).

**Example**: Add `X-User-ID` header based on JWT token.

## Backend for Frontend (BFF) Pattern

**Problem**: Different clients (web, mobile, IoT) need different data.

**Solution**: Create a separate API Gateway for each client type.

```
Web Client → Web BFF → Services
Mobile Client → Mobile BFF → Services
IoT Client → IoT BFF → Services
```

**Benefits**:
- **Optimized for each client**: Mobile BFF returns less data (smaller payloads).
- **Independent evolution**: Change mobile API without affecting web.

## Service Mesh vs API Gateway

| Feature | API Gateway | Service Mesh |
| :--- | :--- | :--- |
| **Scope** | **External** (client → services) | **Internal** (service → service) |
| **Deployment** | Centralized (single gateway) | Distributed (sidecar per service) |
| **Use Case** | Public API, authentication, rate limiting | Service discovery, load balancing, retries, circuit breaking |
| **Examples** | Kong, AWS API Gateway, Nginx | Istio, Linkerd, Consul Connect |

**Key Difference**: API Gateway handles **external traffic**. Service Mesh handles **internal service-to-service** communication.

## Go Implementation: Simple API Gateway

### Routing + Authentication
```go
package main

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "net/url"
    "strings"
)

func main() {
    http.HandleFunc("/", gatewayHandler)
    http.ListenAndServe(":8080", nil)
}

func gatewayHandler(w http.ResponseWriter, r *http.Request) {
    // Authentication
    token := r.Header.Get("Authorization")
    if !isValidToken(token) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    // Routing
    if strings.HasPrefix(r.URL.Path, "/api/users") {
        proxyTo(w, r, "http://localhost:8001") // User Service
    } else if strings.HasPrefix(r.URL.Path, "/api/orders") {
        proxyTo(w, r, "http://localhost:8002") // Order Service
    } else {
        http.Error(w, "Not Found", http.StatusNotFound)
    }
}

func proxyTo(w http.ResponseWriter, r *http.Request, targetURL string) {
    target, _ := url.Parse(targetURL)
    proxy := httputil.NewSingleHostReverseProxy(target)
    proxy.ServeHTTP(w, r)
}

func isValidToken(token string) bool {
    // Validate JWT token (simplified)
    return token == "Bearer valid-token"
}
```

### Rate Limiting Middleware
```go
import (
    "sync"
    "time"
)

type RateLimiter struct {
    requests map[string][]time.Time
    mu       sync.Mutex
    limit    int
    window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        requests: make(map[string][]time.Time),
        limit:    limit,
        window:   window,
    }
}

func (rl *RateLimiter) Allow(clientID string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    now := time.Now()
    cutoff := now.Add(-rl.window)
    
    // Remove old requests
    requests := rl.requests[clientID]
    validRequests := []time.Time{}
    for _, t := range requests {
        if t.After(cutoff) {
            validRequests = append(validRequests, t)
        }
    }
    
    if len(validRequests) >= rl.limit {
        return false
    }
    
    validRequests = append(validRequests, now)
    rl.requests[clientID] = validRequests
    return true
}

func rateLimitMiddleware(limiter *RateLimiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            clientID := r.RemoteAddr // Or use API key
            if !limiter.Allow(clientID) {
                http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
```

## Popular API Gateway Solutions

| Solution | Type | Features |
| :--- | :--- | :--- |
| **Kong** | Open-source | Plugins (auth, rate limiting, logging), Lua scripting |
| **AWS API Gateway** | Managed | Serverless, integrates with Lambda, DynamoDB |
| **Nginx** | Open-source | Reverse proxy, load balancing, caching |
| **Traefik** | Open-source | Auto-discovery (Kubernetes), Let's Encrypt |
| **Envoy** | Open-source | Service mesh (Istio uses Envoy as data plane) |

## Interview Questions

### Q: What's the difference between an API Gateway and a reverse proxy?
**A**: 
- **Reverse Proxy**: Forwards requests to backend servers (basic routing, load balancing).
- **API Gateway**: Adds application-level logic (auth, rate limiting, request aggregation, transformation).

### Q: What is the BFF pattern?
**A**: **Backend for Frontend** - Create separate API Gateways for different client types (web, mobile, IoT) to optimize responses for each.

### Q: How does an API Gateway improve security?
**A**: Centralizes authentication/authorization (verify JWT tokens, API keys), rate limiting (prevent DDoS), and hides internal service architecture from clients.

### Q: What's the difference between API Gateway and Service Mesh?
**A**: 
- **API Gateway**: Handles **external** traffic (client → services).
- **Service Mesh**: Handles **internal** traffic (service → service).
