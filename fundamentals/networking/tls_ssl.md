# TLS/SSL Handshake

## What is TLS?
**Transport Layer Security (TLS)** encrypts data in transit to prevent eavesdropping and tampering. It's the successor to **SSL (Secure Sockets Layer)**.

**HTTPS** = HTTP + TLS

## Goals of TLS
1. **Confidentiality**: Data is encrypted (only sender and receiver can read it).
2. **Integrity**: Data cannot be tampered with (detected via MAC).
3. **Authentication**: Verify the server's identity (via certificates).

## Symmetric vs Asymmetric Encryption

| Type | Description | Speed | Use Case |
| :--- | :--- | :--- | :--- |
| **Symmetric** | Same key for encryption and decryption. | **Fast** | Bulk data encryption (AES). |
| **Asymmetric** | Public key encrypts, private key decrypts. | Slow | Key exchange, digital signatures (RSA, ECDSA). |

**TLS Strategy**: Use **asymmetric encryption** to exchange a **symmetric key**, then use **symmetric encryption** for the actual data.

## TLS 1.2 Handshake

### Full Handshake (2 Round Trips)

#### Round Trip 1: Client Hello + Server Hello
1. **Client Hello**: Client sends:
   - TLS version (1.2)
   - Supported cipher suites (e.g., `TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256`)
   - Random bytes (client random)

2. **Server Hello**: Server sends:
   - Chosen cipher suite
   - Random bytes (server random)
   - **Certificate** (contains server's public key)
   - **Server Key Exchange** (for ECDHE: server's ephemeral public key)
   - **Server Hello Done**

#### Round Trip 2: Key Exchange + Finished
3. **Client Key Exchange**: Client sends:
   - **Pre-Master Secret** (encrypted with server's public key)
   - **Change Cipher Spec** (switch to encrypted communication)
   - **Finished** (encrypted with the new session key)

4. **Server Finished**: Server sends:
   - **Change Cipher Spec**
   - **Finished** (encrypted)

**Session Key Derivation**:
```
Session Key = PRF(Pre-Master Secret, Client Random, Server Random)
```

**Total**: 2 round trips (~100-200ms latency).

## TLS 1.3 Handshake (Faster!)

### 1-RTT Handshake
TLS 1.3 reduces the handshake to **1 round trip**.

#### Round Trip 1: Client Hello + Server Hello + Application Data
1. **Client Hello**: Client sends:
   - Supported cipher suites
   - **Key Share** (client's ephemeral public key for ECDHE)

2. **Server Hello**: Server sends:
   - Chosen cipher suite
   - **Key Share** (server's ephemeral public key)
   - **Certificate** (encrypted with the derived key)
   - **Finished** (encrypted)

3. **Client Finished**: Client sends:
   - **Finished** (encrypted)

**Total**: 1 round trip (~50-100ms latency).

### 0-RTT (Zero Round Trip Time)
For **resumed sessions**, the client can send encrypted application data in the **first message** (no handshake delay).

**Caveat**: Vulnerable to **replay attacks** (attacker can resend the 0-RTT data).

## Certificate Validation

### Certificate Chain
A certificate is signed by a **Certificate Authority (CA)**. The CA's certificate is signed by a **Root CA**.

**Example**:
```
Root CA (trusted by browser)
  ↓ signs
Intermediate CA
  ↓ signs
example.com (server certificate)
```

### Validation Steps
1. **Check Signature**: Verify the certificate is signed by a trusted CA.
2. **Check Expiration**: Ensure the certificate is not expired.
3. **Check Domain**: Ensure the certificate's Common Name (CN) or Subject Alternative Name (SAN) matches the domain.
4. **Check Revocation**: Check if the certificate has been revoked (via CRL or OCSP).

## Perfect Forward Secrecy (PFS)

**Problem**: If the server's private key is compromised, an attacker can decrypt **all past traffic** (if they recorded it).

**Solution**: Use **ephemeral keys** (temporary keys generated per session). Even if the server's private key is compromised, past sessions remain secure.

**Cipher Suites with PFS**:
- `TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256` (ECDHE = Ephemeral Diffie-Hellman)

**Without PFS**:
- `TLS_RSA_WITH_AES_128_CBC_SHA` (RSA key exchange, no ephemeral keys)

## Common Attacks

### 1. Man-in-the-Middle (MITM)
Attacker intercepts traffic and impersonates the server.

**Fix**: Certificate validation (ensure the certificate is signed by a trusted CA).

### 2. Downgrade Attack
Attacker forces the client and server to use a weaker cipher suite.

**Fix**: TLS 1.3 removes support for weak ciphers.

### 3. Replay Attack (0-RTT in TLS 1.3)
Attacker resends a 0-RTT request (e.g., "transfer $100").

**Fix**: Use 0-RTT only for **idempotent** requests (e.g., GET, not POST).

## Go Context: TLS in Go

### HTTPS Server
```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, HTTPS!")
}

func main() {
    http.HandleFunc("/", handler)
    // Requires server.crt and server.key
    http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
}
```

### Custom TLS Config
```go
tlsConfig := &tls.Config{
    MinVersion: tls.VersionTLS13, // Enforce TLS 1.3
    CipherSuites: []uint16{
        tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
    },
}

server := &http.Server{
    Addr:      ":443",
    TLSConfig: tlsConfig,
}
server.ListenAndServeTLS("server.crt", "server.key")
```

### Client: Skip Certificate Verification (INSECURE!)
```go
tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
```

**Warning**: Only use `InsecureSkipVerify` for testing!

## Interview Questions

### Q: What's the difference between TLS 1.2 and TLS 1.3?
**A**: 
- **TLS 1.3**: Faster (1-RTT handshake vs 2-RTT), more secure (removes weak ciphers), supports 0-RTT.
- **TLS 1.2**: Slower, supports more legacy ciphers.

### Q: What is Perfect Forward Secrecy?
**A**: Using ephemeral keys (temporary keys per session) so that compromising the server's private key doesn't compromise past sessions.

### Q: How does a browser verify a certificate?
**A**: 
1. Check the certificate is signed by a trusted CA.
2. Check the certificate is not expired.
3. Check the domain matches the certificate's CN/SAN.
4. Check the certificate is not revoked (CRL/OCSP).
