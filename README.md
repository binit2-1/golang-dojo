## RESOURCES
- [Reddit User Notes](https://github.com/vbd/Fieldnotes/blob/main/golang.md)
- [Go by Example](https://gobyexample.com/)
- [Go Documentation](https://golang.org/doc/)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)


The Masterclass Curriculum

We will move from basic HTTP request lifecycles to distributed concurrent systems.
Project 1: The Middleware Chainer (Logging & Rate Limiter)

You need to understand how HTTP requests flow through a server before you can secure them.

    The Goal: Build a raw net/http API that applies a chain of middlewares to every request.

    What you learn: The http.Handler interface. How to write functions that wrap other functions (closures). You will build a Logger middleware (records response times) and a Rate Limiter middleware (prevents a single IP from making >100 requests a minute).

    Why it matters: Middlewares are the backbone of backend engineering. Every auth check, CORS header, and payload validator runs in middleware.

Project 2: Stateful Session Server (Authentication)

Authentication is proving who you are. We start with traditional, highly secure session-based auth.

    The Goal: A login API that creates a secure session stored in Redis, and sets a strict HttpOnly cookie on the client.

    What you learn: Password hashing (bcrypt), managing Redis connections in Go, setting secure HTTP cookies, and preventing CSRF (Cross-Site Request Forgery).

    Why it matters: Despite the hype around JWTs, session-based auth remains the gold standard for secure, first-party web applications.

Project 3: Stateless Identity Provider (JWT Authentication & Authorization)

Now we move to distributed auth. Authorization is proving what you have permission to do.

    The Goal: Build an OAuth2-style microservice. It issues signed JWTs (JSON Web Tokens) upon login. Other "dummy" services will verify this token without needing to talk to the database.

    What you learn: Cryptography in Go (HMAC vs. RSA keys), JWT Claims, and writing an Authorization middleware that checks if a user's role (e.g., admin vs user) allows them to access a specific route.

    Why it matters: This is how modern microservices and mobile APIs communicate. You will learn the critical flaw of JWTs (invalidation) and how to build a refresh-token rotation system to fix it.

Project 4: The Concurrent Webhook Dispatcher

Time to unleash Go's concurrency on the network.

    The Goal: An API that receives a single payload, and then must concurrently forward that payload (via HTTP POST) to 50 different registered URLs (webhooks) as fast as possible.

    What you learn: sync.WaitGroup, worker pools, and Channels. You will learn how to prevent 50 simultaneous network requests from crashing your server by limiting concurrent workers to a safe number (e.g., 10 at a time).

    Why it matters: This teaches you how to handle heavy I/O operations without blocking the main HTTP thread that is serving users.

Project 5: Real-Time Collaborative Canvas (WebSockets)

REST is stateless. WebSockets are stateful and persistent.

    The Goal: A server where multiple connected clients can draw on a shared digital grid in real-time.

    What you learn: Upgrading a standard net/http connection to a WebSocket protocol. Managing a thread-safe map (sync.RWMutex) of thousands of active client connections. Broadcasting messages efficiently.

    Why it matters: Real-time infrastructure is highly sought after. You will master memory management, as dropped connections that aren't cleaned up will cause massive memory leaks.

Project 6: The Production API (Clean Architecture & Postgres)

The capstone. We bring everything together into a structured, production-ready backend.

    The Goal: A complete RESTful CRUD API (like a blog or a simplified inventory system).

    What you learn: Integrating a real relational database (PostgreSQL). Writing raw SQL queries using the database/sql package (no ORMs!). Implementing database migrations, context cancellation (killing DB queries if the user closes their browser early), and graceful server shutdowns.

    Why it matters: This is the exact architecture you will build day-to-day as a Go backend engineer.