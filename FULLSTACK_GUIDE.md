# Blog Platform - Full Stack Setup Guide

Complete guide to run the blog platform with React frontend and Go backend.

## 🚀 Quick Start with Docker

The easiest way to run both frontend and backend:

```bash
# Navigate to project root
cd blog-platform

# Start everything with Docker Compose
docker-compose up --build
```

Then open:
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080

## 📋 Prerequisites

### For Docker (Recommended)
- Docker Desktop installed
- Docker Compose installed

### For Local Development
- **Backend**: Go 1.19+ and GCC (for SQLite)
- **Frontend**: Node.js 18+ and npm

## 🎯 Running Locally (Without Docker)

### Step 1: Start the Backend

```powershell
# Navigate to backend directory
cd blog-platform/backend

# Set environment variable for database
$env:DB_PATH = "./blog.db"

# Run the backend
go run cmd/api/main.go
```

The API will be available at http://localhost:8080

### Step 2: Start the Frontend

Open a new terminal:

```powershell
# Navigate to frontend directory
cd blog-platform/frontend

# Install dependencies (first time only)
npm install

# Start the development server
npm run dev
```

The frontend will be available at http://localhost:3000

## 🎨 Using the Application

### Create a Post
1. Click "Create Post" in the navigation
2. Fill in the title, content, type, and status
3. Click "Create Post"

### View Posts
- Home page shows all posts in a card grid
- Click on any post card to view full details

### Filter Posts
- Use the dropdown filters to show only specific post types or statuses

### Edit a Post
- Click "Edit" button on a post card or detail page
- Modify the fields
- Click "Update Post"

### Delete a Post
- Click "Delete" button
- Confirm deletion in the popup

## 🏗️ Project Structure

```
blog-platform/
├── backend/             # Go backend
│   ├── cmd/api/        # Application entry point
│   ├── internal/       # Backend business logic
│   │   ├── handler/   # HTTP handlers
│   │   ├── models/    # Database models
│   │   └── service/   # Business services
│   ├── pkg/           # Public libraries
│   ├── go.mod
│   └── Dockerfile
├── frontend/           # React frontend
│   ├── src/
│   │   ├── components/ # React components
│   │   ├── services/   # API services
│   │   ├── App.jsx     # Main app
│   │   └── main.jsx    # Entry point
│   ├── package.json
│   └── vite.config.js
└── docker-compose.yml  # Docker orchestration
```

## 🔧 API Endpoints

All endpoints are prefixed with `/api/v1`:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/posts` | Get all posts (with filters) |
| GET | `/posts/:id` | Get single post |
| POST | `/posts` | Create new post |
| PUT | `/posts/:id` | Update post |
| DELETE | `/posts/:id` | Delete post |

### Query Parameters for GET /posts
- `status`: Filter by status (published, draft, archived)
- `type`: Filter by type (article, tutorial, review)
- `limit`: Number of posts to return (default: 10)
- `offset`: Pagination offset (default: 0)

## 🐳 Docker Configuration

The `docker-compose.yml` now includes:
- **Backend**: Go API server with SQLite
- **Frontend**: React development server
- **Volume**: SQLite database persistence

## 📱 Features

### Frontend Features
- ✨ Modern, responsive UI with gradient design
- 🎨 Color-coded post types and status badges
- 🔍 Real-time filtering
- 📱 Mobile-friendly design
- ⚡ Fast navigation with React Router
- 🎯 Form validation
- 🔄 Automatic API integration

### Backend Features
- 🗄️ SQLite database with auto-schema creation
- 🎨 Design patterns (Singleton, Factory, Observer, Command/Query)
- 🔒 CORS enabled for frontend
- 🚀 RESTful API with Gin framework
- 📊 Post management (CRUD operations)

## 🔥 Development Tips

### Hot Reload
- **Frontend**: Vite automatically reloads on file changes
- **Backend**: Use tools like `air` or `nodemon` for Go hot reload

### Debugging
- **Frontend**: Use React DevTools browser extension
- **Backend**: Add breakpoints in VS Code or use `fmt.Println()`

### Database Reset
To reset the SQLite database:
```powershell
# Stop the application
# Delete the database file
Remove-Item blog.db
# Restart the application (tables will be recreated)
```

## 🚨 Troubleshooting

### Backend won't start
- **Issue**: CGO_ENABLED error
- **Solution**: Install GCC or use Docker

### Frontend can't connect to backend
- **Issue**: CORS errors
- **Solution**: Ensure backend has CORS middleware (already added)

### Port already in use
- **Issue**: Port 8080 or 3000 already in use
- **Solution**: Change ports in code or stop conflicting services

### npm not found
- **Issue**: Node.js not installed
- **Solution**: Install from https://nodejs.org/

## 🎓 Next Steps

1. **Add Authentication**: Implement user login/signup
2. **Add Comments**: Allow comments on posts
3. **Add Images**: Upload and display images
4. **Add Search**: Implement full-text search
5. **Add Categories**: Organize posts by categories
6. **Deploy**: Deploy to production (Vercel for frontend, Cloud Run for backend)

## 📚 Technologies Used

### Frontend
- React 18
- React Router 6
- Axios
- Vite
- CSS3

### Backend
- Go 1.19
- Gin Web Framework
- SQLite3
- Design Patterns

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is for educational purposes.

---

**Enjoy building with the Blog Platform! 🚀**
