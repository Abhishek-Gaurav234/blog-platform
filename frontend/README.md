# Blog Platform Frontend

A modern React frontend for the blog platform built with Vite.

## Features

- 📝 Create, read, update, and delete blog posts
- 🎨 Beautiful, responsive UI with gradient designs
- 🔍 Filter posts by status and type
- 📱 Mobile-friendly design
- ⚡ Fast development with Vite
- 🎯 Type badges and status indicators
- 🔄 Real-time updates

## Prerequisites

- Node.js (v18 or higher)
- npm or yarn
- Backend API running on `http://localhost:8080`

## Installation

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

## Running the Application

1. Start the development server:
```bash
npm run dev
```

2. Open your browser and navigate to:
```
http://localhost:3000
```

The frontend will automatically proxy API requests to `http://localhost:8080`.

## Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build

## Project Structure

```
frontend/
├── src/
│   ├── components/         # React components
│   │   ├── PostList.jsx   # List all posts
│   │   ├── PostDetail.jsx # View single post
│   │   ├── CreatePost.jsx # Create new post
│   │   ├── EditPost.jsx   # Edit existing post
│   │   └── PostForm.jsx   # Reusable form component
│   ├── services/
│   │   └── api.js         # API service layer
│   ├── App.jsx            # Main app component
│   ├── App.css            # App styles
│   ├── main.jsx           # Entry point
│   └── index.css          # Global styles
├── index.html             # HTML template
├── vite.config.js         # Vite configuration
└── package.json           # Dependencies
```

## API Integration

The frontend communicates with the Go backend API through the following endpoints:

- `GET /api/v1/posts` - Get all posts
- `GET /api/v1/posts/:id` - Get single post
- `POST /api/v1/posts` - Create new post
- `PUT /api/v1/posts/:id` - Update post
- `DELETE /api/v1/posts/:id` - Delete post

## Building for Production

1. Build the application:
```bash
npm run build
```

2. The production files will be in the `dist` directory.

3. Preview the production build:
```bash
npm run preview
```

## Environment Configuration

The API base URL is configured in `vite.config.js`. To change it, modify the proxy configuration:

```javascript
server: {
  proxy: {
    '/api': {
      target: 'http://your-api-url:port',
      changeOrigin: true,
    }
  }
}
```

## Technologies Used

- **React** - UI library
- **React Router** - Client-side routing
- **Axios** - HTTP client
- **Vite** - Build tool and dev server
- **CSS3** - Styling with modern features

## Features Showcase

### Post Types
- 📄 Article (Blue badge)
- 📚 Tutorial (Green badge)
- ⭐ Review (Red badge)

### Post Status
- ✅ Published (Green)
- 📝 Draft (Orange)
- 📦 Archived (Gray)

### Filters
- Filter by status (All, Published, Draft, Archived)
- Filter by type (All, Article, Tutorial, Review)

## Troubleshooting

**Problem: Cannot connect to backend**
- Ensure the Go backend is running on `http://localhost:8080`
- Check CORS settings in the backend

**Problem: npm command not found**
- Install Node.js from https://nodejs.org/

**Problem: Dependencies installation fails**
- Clear npm cache: `npm cache clean --force`
- Delete `node_modules` and `package-lock.json`
- Run `npm install` again
