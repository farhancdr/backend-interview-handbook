# DNS Deep Dive

## What is DNS?
The **Domain Name System** translates human-readable domain names (e.g., `google.com`) to IP addresses (e.g., `142.250.185.46`).

## DNS Resolution Process

### Recursive vs Iterative Queries

#### Recursive Query
The DNS resolver does all the work and returns the final answer to the client.

**Flow**:
1. Client asks Recursive Resolver: "What's the IP of `example.com`?"
2. Resolver asks Root Server: "Who handles `.com`?"
3. Root Server responds: "Ask the `.com` TLD server."
4. Resolver asks TLD Server: "Who handles `example.com`?"
5. TLD Server responds: "Ask `ns1.example.com`."
6. Resolver asks Authoritative Server: "What's the IP of `example.com`?"
7. Authoritative Server responds: `93.184.216.34`.
8. Resolver returns the IP to the client.

#### Iterative Query
The DNS resolver returns referrals, and the client does the work.

**Flow**: Same as above, but the client makes each query (rarely used in practice).

### Full DNS Lookup Example
```
Client → Recursive Resolver (e.g., 8.8.8.8)
  ↓
Root Server (.) → "Ask .com TLD"
  ↓
TLD Server (.com) → "Ask ns1.example.com"
  ↓
Authoritative Server (ns1.example.com) → "93.184.216.34"
  ↓
Recursive Resolver → Client
```

## DNS Record Types

| Type | Purpose | Example |
| :--- | :--- | :--- |
| **A** | IPv4 address | `example.com → 93.184.216.34` |
| **AAAA** | IPv6 address | `example.com → 2606:2800:220:1:248:1893:25c8:1946` |
| **CNAME** | Alias (Canonical Name) | `www.example.com → example.com` |
| **MX** | Mail server | `example.com → mail.example.com` (priority 10) |
| **TXT** | Arbitrary text (SPF, DKIM, verification) | `example.com → "v=spf1 include:_spf.google.com ~all"` |
| **NS** | Name server | `example.com → ns1.example.com` |
| **SRV** | Service location | `_http._tcp.example.com → server1.example.com:80` |
| **PTR** | Reverse DNS (IP → domain) | `34.216.184.93.in-addr.arpa → example.com` |

### CNAME vs A Record
- **A Record**: Direct mapping to an IP address.
- **CNAME**: Alias to another domain (which must have an A record).

**Example**:
```
www.example.com → CNAME → example.com → A → 93.184.216.34
```

**Limitation**: You cannot have a CNAME at the root domain (e.g., `example.com` cannot be a CNAME).

## DNS Caching

### Levels of Caching
1. **Browser Cache**: Browsers cache DNS results for a short time.
2. **OS Cache**: The operating system caches DNS results.
3. **Recursive Resolver Cache**: ISP or public DNS (8.8.8.8) caches results.

### Time to Live (TTL)
Each DNS record has a **TTL** (in seconds) that specifies how long it can be cached.

**Example**:
```
example.com.  300  IN  A  93.184.216.34
```
TTL = 300 seconds (5 minutes).

**Trade-off**:
- **Low TTL** (e.g., 60s): Fast propagation of changes, but more DNS queries (higher load).
- **High TTL** (e.g., 86400s = 1 day): Fewer DNS queries, but slow propagation of changes.

## DNS Load Balancing

### Round Robin DNS
Return multiple A records for the same domain. Clients randomly pick one.

**Example**:
```
example.com → 93.184.216.34
example.com → 93.184.216.35
example.com → 93.184.216.36
```

**Pros**: Simple.  
**Cons**: No health checks (if one server is down, clients may still try it). Caching reduces effectiveness.

### GeoDNS
Return different IP addresses based on the client's geographic location.

**Example**:
- US clients → `us-server.example.com` (IP: 1.2.3.4)
- EU clients → `eu-server.example.com` (IP: 5.6.7.8)

**Use Case**: CDNs (serve content from the nearest edge server).

## DNS Security

### DNS Spoofing (Cache Poisoning)
An attacker injects fake DNS records into a resolver's cache.

**Example**: Attacker makes `bank.com` resolve to a phishing site.

**Fix**: **DNSSEC** (DNS Security Extensions) - Cryptographically signs DNS records.

### DNS over HTTPS (DoH) / DNS over TLS (DoT)
Encrypt DNS queries to prevent eavesdropping.

**Traditional DNS**: Queries are sent in plaintext (ISP can see what sites you visit).  
**DoH/DoT**: Queries are encrypted.

## Go Context: DNS Lookup

### Basic DNS Lookup
```go
package main

import (
    "fmt"
    "net"
)

func main() {
    ips, err := net.LookupIP("google.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    for _, ip := range ips {
        fmt.Println("IP:", ip)
    }
}
```

### Custom DNS Resolver
```go
resolver := &net.Resolver{
    PreferGo: true,
    Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
        d := net.Dialer{Timeout: 5 * time.Second}
        return d.DialContext(ctx, network, "8.8.8.8:53") // Use Google DNS
    },
}

ips, err := resolver.LookupIP(context.Background(), "ip4", "example.com")
```

### Reverse DNS Lookup
```go
names, err := net.LookupAddr("8.8.8.8")
if err != nil {
    fmt.Println("Error:", err)
    return
}
for _, name := range names {
    fmt.Println("Name:", name) // dns.google
}
```

## Interview Questions

### Q: What happens if a DNS server is down?
**A**: The client tries the next DNS server in its list (usually configured with multiple DNS servers). If all fail, the client cannot resolve the domain.

### Q: Why is DNS caching important?
**A**: Reduces latency (no need to query DNS servers for every request) and reduces load on DNS servers.

### Q: How does GeoDNS improve performance?
**A**: By routing users to the nearest server (lower latency). CDNs use this to serve content from edge servers close to users.

### Q: What's the difference between CNAME and A record?
**A**: 
- **A Record**: Direct mapping to an IP address.
- **CNAME**: Alias to another domain (which must have an A record). CNAME adds an extra DNS lookup.
