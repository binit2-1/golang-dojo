# golang-dojo

Welcome to golang-dojo — a messy, fun, hands-on playground to get good at Go. This repo is for learning, experimenting, and breaking stuff. Not for flexing on résumés.

Why this repo exists
- Learn idiomatic Go by building tiny servers, microservices, and small tools.
- Break things intentionally: make mistakes, write tests, fix ’em, and repeat.

Quick start
- Install Go (1.20+ recommended).
- From the repo root, run tests across everything:

```bash
go test ./...
```

- Run a service (example):

```bash
cd ticketblitz && go run ./cmd/api
```

Hands-on learning path (what's inside)
- `middleware-chaining` — raw `net/http` middleware: logging, rate-limiting, and chaining patterns.
- `auth-project` & `microservices-jwt` — stateful (session + Redis) and stateless (JWT) auth patterns.
- `webhook-dispatcher` — worker pools, concurrency, and retry/backoff strategies.
- `realtime-canvas` — WebSockets and real-time broadcasting.
- `ticketblitz`, `ecom`, `sysmon` — larger examples tying DB, metrics, and clean architecture together.

How to use this repo (learn by doing)
- Pick a folder that excites you.
- Read the small README or `cmd/...` entrypoint in that folder.
- Run it, add `fmt.Println` and break the behavior to learn what each part does.
- Prefer writing a failing test first (Learn Go With Tests style).

Style & etiquette
- This repo favors clarity and education over production polish.
- Expect global state, quick hacks, and comment-rich experiments.
- If you want production-grade examples, ask — we can lift patterns into a clean module.

Contribute
- Add a short note to the project you explored: what you learned, what broke, and how you fixed it.
- Open a PR for small fixes, better README sections, or toy exercises.

Resources
- Go by Example: https://gobyexample.com/
- Learn Go With Tests: https://quii.gitbook.io/learn-go-with-tests/

Have fun. Break things. Send questions or PRs.

— The golang-dojo crew