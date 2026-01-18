# Networking Fundamentals

## TCP vs UDP

| Feature | TCP (Transmission Control Protocol) | UDP (User Datagram Protocol) |
| :--- | :--- | :--- |
| **Reliability** | **Reliable**. Guarantees delivery and order. | **Unreliable**. Best effort. Packets may drop or arrive out of order. |
| **Connection** | Connection-oriented (3-way handshake). | Connectionless (Fire and Forget). |
| **Overhead** | High (ACKs, Retries, Flow Control). | Low (Header is only 8 bytes). |
| **Use Case** | Web (HTTP), Email, File Transfer. | Video Streaming, Gaming, Voice Calls. |

### TCP Handshake (3-Way)
1.  **SYN**: Client says "Let's connect, my sequence is X".
2.  **SYN-ACK**: Server says "OK, I see X. My sequence is Y".
3.  **ACK**: Client says "OK, I see Y. Connected."

## HTTP Evolution

### HTTP/1.1
-   **Text-based**.
-   **Head-of-Line Blocking**: Sequential requests on one connection. If request 1 is slow, request 2 waits.
-   **Keep-Alive**: Reuses TCP connection.

### HTTP/2
-   **Binary protocol** (more efficient parsing).
-   **Multiplexing**: Multiple streams on one connection. No more blocking!
-   **Header Compression** (HPACK).
-   **Server Push**.

### HTTP/3 (QUIC)
-   **UDP-based**: Replaces TCP with QUIC.
-   **Why?**: TCP has Head-of-Line blocking at the packet level (if one packet is lost, OS holds back all subsequent ones). QUIC solves this.
-   **Built-in TLS 1.3**.

## What happens when you type `google.com`?
1.  **DNS Lookup**: Browser checks cache -> OS cache -> Router -> DNS Server (Recursive lookup) to get IP.
2.  **TCP Handshake**: SYN -> SYN/ACK -> ACK with google.com IP.
3.  **TLS Handshake**: Exchange keys to encrypt traffic (HTTPS).
4.  **HTTP Request**: Browser sends `GET /`.
5.  **HTTP Response**: Server sends HTML.
6.  **Rendering**: Browser parses HTML, fetches CSS/JS, paints DOM.
