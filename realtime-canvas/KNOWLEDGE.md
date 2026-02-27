The Definitive WebSocket Architecture Guide (Go + Next.js)
==========================================================

This guide serves as a universal reference for building, scaling, and maintaining real-time applications using Go (gorilla/websocket) and React/Next.js.

Table of Contents
-----------------

1.  [Core Architecture (The Correct Way)](https://www.google.com/search?q=#1-core-architecture-the-correct-way)
    
2.  [React / Next.js Integration](https://www.google.com/search?q=#2-react--nextjs-integration)
    
3.  [Use Case Blueprints](https://www.google.com/search?q=#3-use-case-blueprints)
    
    *   Chat Applications
        
    *   Collaborative Canvases
        
    *   Live Leaderboards
        
4.  [Production Setup & Scaling Checklist](https://www.google.com/search?q=#4-production-setup--scaling-checklist)
    
5.  [References & Links](https://www.google.com/search?q=#5-references--links)
    

1\. Core Architecture (The Correct Way)
---------------------------------------

Do not reinvent the wheel. The undisputed industry standard for managing WebSockets in Go relies on the **Hub & Pump Architecture** using Channels, avoiding Mutex locks entirely.

### The Backend Components:

*   **The Hub:** A single background goroutine containing a map of all clients. It listens to Register, Unregister, and Broadcast channels using a select statement.
    
*   **The Read Pump:** A goroutine per client that constantly reads from the TCP socket and pushes data to the Hub's Broadcast channel.
    
*   **The Write Pump:** A goroutine per client that listens to a personal Send channel and writes data back to the client's TCP socket.
    

### Backend Best Practices:

*   **Heartbeats (Ping/Pong):** _Crucial._ Load balancers (like AWS ELB or Nginx) will drop idle TCP connections after 60 seconds. Your Write Pump must send a PingMessage every 50 seconds. Your Read Pump must set a SetReadDeadline and extend it every time it receives a PongMessage. If a pong is missed, the server assumes the client disconnected and cleans up the memory.
    
*   **Buffered Channels:** The Send channel for each client must be buffered (e.g., make(chan \[\]byte, 256)). If a client's internet is slow and the channel fills up, the Hub must immediately disconnect them rather than blocking the entire broadcast loop.
    
*   **JSON vs. Binary:** Use websocket.TextMessage for JSON (Chats, Leaderboards). Use websocket.BinaryMessage for heavy binary payloads (File streaming, complex Canvas data) to save bandwidth.
    

2\. React / Next.js Integration
-------------------------------

In Next.js, WebSockets **must** exist exclusively on the client side. You cannot instantiate a WebSocket inside a Server Component or during Server-Side Rendering (SSR).

### The Golden Custom Hook (useWebSocket.ts)

Always wrap your WebSocket logic in a custom hook. This ensures proper cleanup when a component unmounts (preventing the dreaded React Strict Mode double-connection bug) and handles auto-reconnection.

```ts
"use client";
import { useEffect, useRef, useState } from "react";


interface WSMessage {
  type: string;
  payload: any;
}


export function useWebSocket(url: string) {
  const [messages, setMessages] = useState<WSMessage[]>([]);
  const [isConnected, setIsConnected] = useState(false);
  const wsRef = useRef<WebSocket | null>(null);


  useEffect(() => {
    // 1. Initialize Connection
    const ws = new WebSocket(url);
    wsRef.current = ws;


    ws.onopen = () => setIsConnected(true);
    
    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setMessages((prev) => [...prev, data]);
    };


    ws.onclose = () => setIsConnected(false);


    // 2. The Cleanup Function (CRITICAL)
    return () => {
      if (wsRef.current) {
        wsRef.current.close();
      }
    };
  }, [url]);


  // Expose a typed send function
  const sendMessage = (msg: WSMessage) => {
    if (wsRef.current && wsRef.current.readyState === WebSocket.OPEN) {
      wsRef.current.send(JSON.stringify(msg));
    }
  };


  return { messages, isConnected, sendMessage };
}



```

3\. Use Case Blueprints
-----------------------

Depending on what you are building, the data flow changes drastically.

### A. Chat Applications

*   **The Problem:** Users need to see history when they join, plus live updates.
    
*   **The Solution:** 1. On load, Next.js makes a standard REST GET /v1/chats request to fetch the last 50 messages from the database.2. Simultaneously, open the WebSocket connection.3. When a user sends a message, send it via WebSocket. The Go server saves it to Postgres _first_, then broadcasts the saved JSON to the Hub.
    

### B. Collaborative Canvases

*   **The Problem:** Mouse movements generate hundreds of events per second. Broadcasting every single pixel will crash the server and the browser.
    
*   **The Solution (Throttling):** 1. The frontend must use requestAnimationFrame or lodash throttle (e.g., 30ms) before calling ws.send().2. Send **Vectors/Deltas**, not the whole image. Send { x0, y0, x1, y1, color } representing a single stroke, rather than the base64 image data.
    

### C. Live Leaderboards

*   **The Problem:** In a high-speed game, scores update 1,000 times a second. Broadcasting 1,000 updates to 10,000 users = 10,000,000 messages/sec.
    
*   **The Solution (Server Ticks):** 1. Do not broadcast on every score change.2. Instead, the Go server maintains the leaderboard in memory.3. A separate goroutine runs a time.Ticker every 1 second (1000ms).4. Every tick, the server takes a snapshot of the top 10 players and broadcasts that array to the clients once.
    

4\. Production Setup & Scaling Checklist
----------------------------------------

Before deploying a WebSocket server to production, verify this checklist:

*   \[ \] **WSS (TLS):** Always use wss:// in production. Browsers block ws:// connections on https:// websites (Mixed Content Security Policy).
    
*   \[ \] **CheckOrigin:** In your websocket.Upgrader, explicitly verify the r.Header.Get("Origin") matches your frontend domain to prevent Cross-Site WebSocket Hijacking.
    
*   \[ \] **Ulimit (File Descriptors):** In Linux, every WebSocket connection consumes one File Descriptor. The default OS limit is often 1,024. If 1,025 users connect, your Go server will panic. You must configure your server OS limit (ulimit -n 65535).
    
*   location /ws { proxy\_pass http://localhost:8080; proxy\_http\_version 1.1; proxy\_set\_header Upgrade $http\_upgrade; proxy\_set\_header Connection "Upgrade"; proxy\_read\_timeout 3600s; # Keep alive for 1 hour}
    
*   \[ \] **Scaling Beyond 1 Server:** A single Go server can handle ~100,000 concurrent connections. If you need 1,000,000, you must deploy multiple Go servers and use **Redis Pub/Sub** to act as a bridge so Server A can broadcast to users connected to Server B.
    

5\. References & Links
----------------------

*   **Gorilla WebSocket Official Docs (The Standard):** [https://pkg.go.dev/github.com/gorilla/websocket](https://www.google.com/search?q=https://pkg.go.dev/github.com/gorilla/websocket)
    
*   **Gorilla Chat Example (Study the code here):** [https://github.com/gorilla/websocket/tree/master/examples/chat](https://www.google.com/search?q=https://github.com/gorilla/websocket/tree/master/examples/chat)
    
*   **Scaling WebSockets in Go:** [Centrifugo Blog on Architecture](https://centrifugal.dev/blog/2020/11/12/scaling-websocket)
    
*   **Next.js Custom Server WebSocket Limitations:** Note that Vercel Serverless Functions _do not_ support WebSockets. You must host your Go server on a VPS (AWS EC2, DigitalOcean, Railway) or use a managed service.