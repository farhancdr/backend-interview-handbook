# 🌐 Networking Fundamentals

> **Master the protocols that connect distributed systems**

Networking knowledge is essential for backend engineers. This section covers everything from TCP/IP to modern protocols like gRPC and HTTP/3.

---

## 📖 Topics

### Protocol Foundations

| Topic | Description | Key Interview Questions |
|:------|:------------|:------------------------|
| **[TCP, UDP & HTTP](tcp_udp_http.md)** | Transport layer protocols and HTTP basics | "Explain the TCP 3-way handshake" |
| **[TLS/SSL](tls_ssl.md)** | Encryption, certificates, handshake process | "How does HTTPS work?" |
| **[DNS Deep Dive](dns_deep_dive.md)** | DNS resolution, caching, load balancing | "Explain the DNS resolution process" |

### Modern Protocols

| Topic | Description | Key Interview Questions |
|:------|:------------|:------------------------|
| **[Modern Protocols](modern_protocols.md)** | HTTP/2, HTTP/3, gRPC, GraphQL | "What are the benefits of HTTP/2 over HTTP/1.1?" |
| **[Real-time Communication](realtime_communication.md)** | WebSockets, Server-Sent Events, long polling | "When would you use WebSockets vs SSE?" |

### Infrastructure & Architecture

| Topic | Description | Key Interview Questions |
|:------|:------------|:------------------------|
| **[Load Balancing](load_balancing.md)** | L4 vs L7, algorithms, health checks | "Explain the difference between Layer 4 and Layer 7 load balancing" |
| **[CDN & Edge Computing](cdn_edge.md)** | Content delivery, edge caching, PoPs | "How does a CDN improve performance?" |
| **[API Gateway](api_gateway.md)** | Routing, rate limiting, authentication | "What problems does an API Gateway solve?" |

---

## 🎯 Interview Preparation Guide

### Must-Know Concepts

1. **OSI Model** - Understand the 7 layers (focus on L4 and L7)
2. **TCP vs UDP** - Know when to use each protocol
3. **HTTP Methods** - GET, POST, PUT, DELETE, PATCH semantics
4. **Status Codes** - 2xx, 3xx, 4xx, 5xx meanings
5. **TLS Handshake** - Understand the certificate exchange process
6. **Load Balancing Algorithms** - Round-robin, least connections, consistent hashing

### Common Interview Patterns

**System Design Questions:**
- "Design a chat application" → Requires WebSockets knowledge
- "Design a video streaming service" → Requires CDN and HTTP understanding
- "Design an API rate limiter" → Requires understanding of API Gateway patterns

**Technical Deep Dives:**
- "What happens when you type google.com in your browser?"
- "Explain how HTTPS prevents man-in-the-middle attacks"
- "How would you debug a slow API response?"

---

## 🔗 Related Topics

- **[System Design Patterns](../../internal/system_design/)** - See rate limiter and pub-sub implementations
- **[Concurrency Patterns](../../internal/concurrency/)** - Understand concurrent network programming in Go
- **[OS Fundamentals](../os/)** - Learn about system calls used in networking

---

## 📚 Study Tips

1. **Visualize the flow** - Draw packet flows and handshake diagrams
2. **Use Wireshark** - Capture and analyze real network traffic
3. **Practice with curl** - Experiment with HTTP headers and methods
4. **Understand trade-offs** - TCP reliability vs UDP speed, HTTP/1.1 vs HTTP/2

### Hands-on Practice

```bash
# Inspect HTTP headers
curl -v https://google.com

# Test DNS resolution
dig google.com

# Check TLS certificate
openssl s_client -connect google.com:443
```

---

## 🌟 Pro Tips

- **Latency matters** - Always consider network latency in system design
- **Idempotency** - Understand why it's critical for HTTP APIs
- **Connection pooling** - Know how to reuse TCP connections efficiently
- **Timeouts** - Always set appropriate timeouts for network calls

---

[← Back to Fundamentals](../) | [↑ Back to Main](../../README.md)
