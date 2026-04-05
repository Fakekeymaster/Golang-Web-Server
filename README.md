# Go API Gateway

A high-performance API Gateway built in Go that supports request forwarding, authentication, rate limiting, load balancing, and fault tolerance.

## Features

- Request forwarding (GET/POST support)
- Middleware architecture
- JWT-style authentication (Bearer token)
- Rate limiting per client
- Load balancing (round-robin)
- Health checks for backend services
- Retry mechanism for failed requests
- Timeout handling
- Latency logging

## Architecture

Client → API Gateway → Backend Servers → Response

## Tech Stack

- Go (Golang)
- net/http
- Goroutines
- Mutex (thread safety)

## How to Run

```bash
go run main.go