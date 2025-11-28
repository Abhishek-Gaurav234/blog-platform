import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { postsAPI } from '../services/api'
import './PostList.css'

function PostList() {
  const [posts, setPosts] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)
  const [filter, setFilter] = useState({ status: '', type: '' })

  useEffect(() => {
    fetchPosts()
  }, [filter])

  const fetchPosts = async () => {
    try {
      setLoading(true)
      const params = {}
      if (filter.status) params.status = filter.status
      if (filter.type) params.type = filter.type
      
      const response = await postsAPI.getAllPosts(params)
      setPosts(response.data || [])
      setError(null)
    } catch (err) {
      setError('Failed to fetch posts. Make sure the backend server is running.')
      console.error('Error fetching posts:', err)
    } finally {
      setLoading(false)
    }
  }

  const handleDelete = async (id) => {
    if (!window.confirm('Are you sure you want to delete this post?')) {
      return
    }

    try {
      await postsAPI.deletePost(id)
      setPosts(posts.filter(post => post.id !== id))
    } catch (err) {
      alert('Failed to delete post')
      console.error('Error deleting post:', err)
    }
  }

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    })
  }

  const getTypeColor = (type) => {
    const colors = {
      article: '#3498db',
      tutorial: '#2ecc71',
      review: '#e74c3c'
    }
    return colors[type] || '#95a5a6'
  }

  const getStatusBadge = (status) => {
    const badges = {
      published: { color: '#27ae60', text: 'Published' },
      draft: { color: '#f39c12', text: 'Draft' },
      archived: { color: '#7f8c8d', text: 'Archived' }
    }
    return badges[status] || badges.draft
  }

  if (loading) {
    return <div className="loading">Loading posts...</div>
  }

  if (error) {
    return <div className="error">{error}</div>
  }

  return (
    <div className="post-list-container">
      <div className="header">
        <h1>Blog Posts</h1>
        <Link to="/create" className="btn btn-primary">
          + New Post
        </Link>
      </div>

      <div className="filters">
        <select 
          value={filter.status} 
          onChange={(e) => setFilter({...filter, status: e.target.value})}
          className="filter-select"
        >
          <option value="">All Status</option>
          <option value="published">Published</option>
          <option value="draft">Draft</option>
          <option value="archived">Archived</option>
        </select>

        <select 
          value={filter.type} 
          onChange={(e) => setFilter({...filter, type: e.target.value})}
          className="filter-select"
        >
          <option value="">All Types</option>
          <option value="article">Article</option>
          <option value="tutorial">Tutorial</option>
          <option value="review">Review</option>
        </select>
      </div>

      {posts.length === 0 ? (
        <div className="no-posts">
          <p>No posts found. Create your first post!</p>
          <Link to="/create" className="btn btn-primary">Create Post</Link>
        </div>
      ) : (
        <div className="posts-grid">
          {posts.map(post => (
            <div key={post.id} className="post-card">
              <div className="post-header">
                <span 
                  className="post-type" 
                  style={{ backgroundColor: getTypeColor(post.type) }}
                >
                  {post.type}
                </span>
                <span 
                  className="post-status"
                  style={{ backgroundColor: getStatusBadge(post.status).color }}
                >
                  {getStatusBadge(post.status).text}
                </span>
              </div>

              <Link to={`/posts/${post.id}`} className="post-link">
                <h2 className="post-title">{post.title}</h2>
              </Link>

              <p className="post-excerpt">
                {post.content.substring(0, 150)}
                {post.content.length > 150 ? '...' : ''}
              </p>

              <div className="post-meta">
                <span className="post-date">
                  📅 {formatDate(post.created_at)}
                </span>
                <span className="post-author">
                  👤 Author #{post.author_id}
                </span>
              </div>

              <div className="post-actions">
                <Link to={`/posts/${post.id}`} className="btn btn-view">
                  View
                </Link>
                <Link to={`/edit/${post.id}`} className="btn btn-edit">
                  Edit
                </Link>
                <button 
                  onClick={() => handleDelete(post.id)} 
                  className="btn btn-delete"
                >
                  Delete
                </button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export default PostList
