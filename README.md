# 📝 Blog Platform - Full Stack Application

A modern full-stack blog platform with **React frontend** and **Go backend**, implementing multiple design patterns and best practices.

> **📁 Organized Structure:** Backend code is in `backend/` folder, frontend in `frontend/` folder for clean separation.

## ✨ Features

### Frontend (React)
- 🎨 Modern, responsive UI with gradient designs
- 📱 Mobile-friendly interface
- 🔍 Real-time filtering by status and type
- ⚡ Fast navigation with React Router
- 🎯 Form validation
- 🌈 Color-coded post types and status badges

### Backend (Go)
- **CQRS Pattern**: Separation of read and write operations
- **Factory Pattern**: Creation of different content types (articles, tutorials, reviews)
- **Observer Pattern**: Event-driven architecture for post notifications
- **Singleton Pattern**: Database connection management
- **Circuit Breaker Pattern**: Resilience for external service calls

## 🚀 Technology Stack

### Frontend
- **React 18** - UI library
- **React Router 6** - Client-side routing
- **Axios** - HTTP client
- **Vite** - Build tool and dev server
- **CSS3** - Modern styling

### Backend
- **Go 1.19+** - Programming language
- **Gin** - Web framework
- **SQLite** - Lightweight database
- **Design Patterns** - CQRS, Factory, Observer, Singleton, Circuit Breaker

## 📁 Project Structure

```
blog-platform/
├── backend/              # Go backend
│   ├── cmd/api/         # Application entry point
│   ├── internal/        # Private application code
│   │   ├── handler/    # HTTP handlers
│   │   ├── models/     # Data models & database
│   │   └── service/    # Business logic
│   ├── pkg/            # Public libraries
│   ├── go.mod          # Go dependencies
│   ├── go.sum
│   └── Dockerfile
├── frontend/            # React frontend
│   ├── src/
│   │   ├── components/ # React components
│   │   ├── services/   # API client
│   │   ├── App.jsx
│   │   └── main.jsx
│   ├── package.json
│   └── vite.config.js
├── docker-compose.yml   # Orchestration
└── README.md
```

## 🚀 Quick Start

### Option 1: Docker (Recommended - Easiest)

```bash
# Start both frontend and backend
docker-compose up --build

# Or use the batch file
start-docker.bat
```

Then open:
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080

### Option 2: Local Development

**Prerequisites:**
- Node.js 18+ and npm
- Go 1.19+
- GCC (for SQLite on Windows)

**Using PowerShell script:**
```powershell
.\start-local.ps1
```

**Or manually:**

Terminal 1 (Backend):
```bash
go run cmd/api/main.go
```

Terminal 2 (Frontend):
```bash
cd frontend
npm install
npm run dev
```

## 📚 Documentation

- **[FULLSTACK_GUIDE.md](FULLSTACK_GUIDE.md)** - Complete setup and usage guide
- **[MIGRATION_NOTES.md](MIGRATION_NOTES.md)** - PostgreSQL to SQLite migration details
- **[frontend/README.md](frontend/README.md)** - Frontend-specific documentation

## API Endpoints

### Posts

- `POST /api/v1/posts` - Create a new post
- `GET /api/v1/posts` - List all posts
- `GET /api/v1/posts/:id` - Get a specific post
- `PUT /api/v1/posts/:id` - Update a post
- `DELETE /api/v1/posts/:id` - Delete a post
- `GET /api/v1/posts/search` - Search posts (with circuit breaker)

## 🎨 Screenshots & Features

### Post Management
- ✅ Create, read, update, and delete posts
- 📊 Filter by status (Published, Draft, Archived)
- 🏷️ Filter by type (Article, Tutorial, Review)
- 📱 Responsive grid layout

### Post Types
- 📄 **Article** (Blue) - General blog articles
- 📚 **Tutorial** (Green) - Step-by-step guides
- ⭐ **Review** (Red) - Product or service reviews

### Post Status
- ✅ **Published** (Green) - Live posts
- 📝 **Draft** (Orange) - Work in progress
- 📦 **Archived** (Gray) - Old posts

## ⚙️ Configuration

### Backend Environment Variables
- `DB_PATH` - SQLite database file path (default: `./blog.db`)

### Frontend Environment Variables
- `VITE_API_URL` - Backend API URL (default: proxied to `http://localhost:8080`)

## Design Patterns Implementation

### 1. CQRS Pattern

- **Command Service**: Handles write operations (Create, Update, Delete)
- **Query Service**: Handles read operations (Get, List)

### 2. Factory Pattern

- **Content Factory**: Creates different content types (Article, Tutorial, Review)
- Each content type has specific validation and fields

### 3. Observer Pattern

- **Post Service**: Manages observers for post events
- **Observers**: SearchIndex, Notification, Analytics
- Events: post_created, post_updated, post_deleted

### 4. Singleton Pattern

- **Database**: Single instance database connection

### 5. Circuit Breaker Pattern

- **Circuit Breaker**: Protects external service calls
- **Search Service**: Wrapped with circuit breaker for resilience

## Testing

```bash
go test ./...
```

## 🐳 Docker Support

The application includes full Docker support with:
- Backend container with Go and SQLite
- Frontend container with Node.js and Vite
- Persistent volume for SQLite database
- Network configuration for inter-service communication

```bash
# Start in detached mode
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## 📁 Enhanced Project Structure

```
blog-platform/
├── frontend/              # React frontend
│   ├── src/
│   │   ├── components/   # React components
│   │   │   ├── PostList.jsx
│   │   │   ├── PostDetail.jsx
│   │   │   ├── CreatePost.jsx
│   │   │   ├── EditPost.jsx
│   │   │   └── PostForm.jsx
│   │   ├── services/     # API services
│   │   │   └── api.js
│   │   ├── App.jsx
│   │   └── main.jsx
│   ├── package.json
│   └── vite.config.js
├── cmd/api/              # Backend entry point
│   └── main.go
├── internal/             # Backend logic
│   ├── handler/         # HTTP handlers
│   ├── models/          # Database models
│   └── service/         # Business services
├── pkg/                 # Shared packages
│   └── circuitbreaker/
├── docker-compose.yml   # Docker orchestration
├── Dockerfile          # Backend container
└── README.md
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

MIT License
