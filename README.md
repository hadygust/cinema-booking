# Cinema Booking — Concurrency-Safe Seat Reservation

A **cinema ticket booking system** built in Go that solves the classic double-booking race condition — ensuring two users can never claim the same seat, even under high-concurrency load.

---

## The Problem This Solves

Without protection, two users booking the same seat at the same instant both succeed:

```
User A ──► read seat A1 → "free" ──► write booking ──► success ✓
User B ──► read seat A1 → "free" ──► write booking ──► success ✓ (BUG!)
```

Now two people show up for the same seat. This project demonstrates multiple strategies to prevent this in a high-contention environment, providing a good user experience when a seat is already taken.

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.25 |
| Concurrency Control | Redis distributed locking (`go-redis/v8`) |
| IDs | UUID (`google/uuid`) |
| Containerization | Docker Compose |

---

## Features

- **Distributed locking** via Redis to guarantee at-most-one winner per seat
- **Interactive seat map** UI — click to book, instant feedback on conflicts
- **High-contention safe** — handles concurrent booking attempts gracefully
- **Docker Compose** — spins up the full stack with one command

---

## Project Structure

```
cinema-booking/
├── cmd/                        # Application entry point
├── internal/
│   ├── adapter/          
│   ├── booking/                # 
│   └── service/                # Core booking business logic
│       ├── memory_store.go     # Simulate a storage system and booking logic
│       ├── concurrent_store.go # Simulate a storage system and booking logic with locking
│       └── redis_store.go      # Redis storage and booking logic with locking
├── docker-compose.yaml         # Redis + app setup
├── go.mod
└── go.sum
```

---

## Getting Started

### Prerequisites

- Go 1.25+
- Docker & Docker Compose

### 1. Clone the repository

```bash
git clone https://github.com/hadygust/cinema-booking.git
cd cinema-booking
```

### 2. Start Redis

```bash
docker compose up -d
```

### 3. Change store variable to the store system you want to test by calling its constructor

```Go
store := NewMemoryStore()                                   # Memory Store
store := NewConcurrentStore()                               # Concurrent Store
store := NewRedisStore(adapter.NewClient("localhost:6379")) # Redis Store
```

### 4. Run the test with race condition

```bash
go test .\internal\booking\... -v -race
```

---

## How the Locking Works

When a user selects a seat, the server:

1. Attempts to acquire a **Redis distributed lock** on that seat ID with a short TTL
2. If the lock is acquired → the booking proceeds and is committed
3. If the lock is already held → the user is immediately notified the seat is taken
4. The lock is released after the booking transaction completes

This ensures only one request can write a booking for any given seat at a time, regardless of server replicas or concurrent users.

---

## Author

**Hady Gustianto** — [LinkedIn](https://linkedin.com/in/hadygustianto) · [GitHub](https://github.com/hadygust)
