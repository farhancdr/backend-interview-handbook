# Real-Time Communication: WebSockets vs SSE vs Long Polling

## The Problem
HTTP is **request-response** (client asks, server responds). But what if the server needs to **push** data to the client (e.g., chat messages, live notifications)?

## Solutions

### 1. Short Polling (Naive Approach)
Client repeatedly sends requests to the server (e.g., every 5 seconds).

**Pros**: Simple.  
**Cons**: High latency (up to 5 seconds). Wastes bandwidth (many empty responses).

### 2. Long Polling
Client sends a request. Server holds the connection open until new data is available (or timeout).

**Flow**:
1. Client: `GET /messages`
2. Server: (waits for new message)
3. Server: (new message arrives) → Response: `{"msg": "Hello"}`
4. Client: Immediately sends another `GET /messages`

**Pros**: Lower latency than short polling.  
**Cons**: Still uses HTTP overhead (headers, connection setup). Server must hold many open connections.

### 3. Server-Sent Events (SSE)
Server pushes data to the client over a **single long-lived HTTP connection**.

**Protocol**: `Content-Type: text/event-stream`

**Example**:
```
data: {"msg": "Hello"}

data: {"msg": "World"}
```

**Pros**: Simple (built on HTTP). Automatic reconnection. Supported by browsers (`EventSource` API).  
**Cons**: **Unidirectional** (server → client only). Limited to text (no binary). Max 6 connections per domain (browser limit).

### 4. WebSockets
**Full-duplex** (bidirectional) communication over a single TCP connection.

**Handshake** (HTTP Upgrade):
```
Client → Server:
GET /chat HTTP/1.1
Upgrade: websocket
Connection: Upgrade

Server → Client:
HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
```

After the handshake, the connection switches to the WebSocket protocol (no more HTTP overhead).

**Pros**: **Bidirectional**. Low latency. Supports binary data. No HTTP overhead after handshake.  
**Cons**: More complex than SSE. Requires WebSocket server support. Harder to scale (stateful connections).

## Comparison Table

| Feature | Long Polling | SSE | WebSockets |
| :--- | :--- | :--- | :--- |
| **Direction** | Client → Server (request-response) | Server → Client (unidirectional) | **Bidirectional** |
| **Protocol** | HTTP | HTTP (text/event-stream) | WebSocket (after HTTP upgrade) |
| **Latency** | Medium (depends on poll interval) | Low | **Very Low** |
| **Overhead** | High (HTTP headers per request) | Low (single connection) | **Very Low** (no HTTP headers) |
| **Binary Support** | Yes | No (text only) | Yes |
| **Browser Support** | Universal | Modern browsers (`EventSource`) | Modern browsers (`WebSocket`) |
| **Use Case** | Legacy systems | Live notifications, stock tickers | **Chat, gaming, real-time collaboration** |

## When to Use Each

### Long Polling
- **Legacy systems** where WebSockets/SSE are not supported.
- **Infrequent updates** (e.g., check for new email every minute).

### Server-Sent Events (SSE)
- **Unidirectional** data flow (server → client).
- **Live updates**: Stock prices, news feeds, notifications.
- **Simplicity**: Easier to implement than WebSockets.

### WebSockets
- **Bidirectional** communication (e.g., chat, multiplayer games).
- **Low latency** is critical.
- **High-frequency updates** (e.g., real-time collaboration tools like Google Docs).

## Go Implementation

### WebSocket Server (using `gorilla/websocket`)
```go
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Upgrade error:", err)
        return
    }
    defer conn.Close()

    for {
        msgType, msg, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Read error:", err)
            break
        }
        fmt.Printf("Received: %s\n", msg)
        
        // Echo back
        if err := conn.WriteMessage(msgType, msg); err != nil {
            fmt.Println("Write error:", err)
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", wsHandler)
    http.ListenAndServe(":8080", nil)
}
```

### SSE Server (Standard Library)
```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")

    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
        return
    }

    for i := 0; i < 10; i++ {
        fmt.Fprintf(w, "data: Message %d\n\n", i)
        flusher.Flush()
        time.Sleep(1 * time.Second)
    }
}

func main() {
    http.HandleFunc("/events", sseHandler)
    http.ListenAndServe(":8080", nil)
}
```

### Long Polling Server
```go
func longPollHandler(w http.ResponseWriter, r *http.Request) {
    // Wait for new data (or timeout after 30 seconds)
    select {
    case msg := <-messageChan:
        json.NewEncoder(w).Encode(msg)
    case <-time.After(30 * time.Second):
        w.WriteHeader(http.StatusNoContent) // No new data
    }
}
```

## Scaling Considerations

### WebSockets
**Challenge**: Stateful connections (each client holds a connection to a specific server).

**Solutions**:
1. **Sticky Sessions**: Route clients to the same server (load balancer uses IP hash or cookies).
2. **Message Broker** (Redis Pub/Sub, Kafka): Servers subscribe to a shared message bus. When a message arrives, all servers push it to their connected clients.

### SSE
Same challenges as WebSockets (stateful connections).

### Long Polling
Easier to scale (stateless). Each request can go to any server.

## Interview Questions

### Q: When would you use WebSockets over SSE?
**A**: When you need **bidirectional** communication (e.g., chat, gaming). SSE is unidirectional (server → client only).

### Q: How do you scale WebSocket servers?
**A**: Use **sticky sessions** (route clients to the same server) or a **message broker** (Redis Pub/Sub) to synchronize messages across servers.

### Q: What's the main drawback of long polling?
**A**: High overhead (HTTP headers per request) and higher latency compared to WebSockets/SSE.
