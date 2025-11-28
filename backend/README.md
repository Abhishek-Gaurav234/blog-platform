# Backend - Blog Platform API

Go-based REST API backend for the blog platform.

## Structure

```
backend/
├── cmd/api/              # Application entry point
│   └── main.go
├── internal/             # Private application code
│   ├── handler/         # HTTP request handlers
│   ├── models/          # Data models and database
│   └── service/         # Business logic services
├── pkg/                 # Public libraries
│   └── circuitbreaker/  # Circuit breaker pattern
├── go.mod               # Go dependencies
├── go.sum               # Dependency checksums
└── Dockerfile           # Container configuration
```

## Running Locally

```bash
cd backend
go run cmd/api/main.go
```

The server will start on http://localhost:8080

## Building

```bash
cd backend
go build -o blog-api cmd/api/main.go
```

## Docker

```bash
cd backend
docker build -t blog-api .
docker run -p 8080:8080 blog-api
```

## API Endpoints

- `GET /api/v1/posts` - List all posts
- `GET /api/v1/posts/:id` - Get single post
- `POST /api/v1/posts` - Create new post
- `PUT /api/v1/posts/:id` - Update post
- `DELETE /api/v1/posts/:id` - Delete post

## Environment Variables

- `DB_PATH` - SQLite database file path (default: `./blog.db`)
