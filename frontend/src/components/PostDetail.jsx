import { useState, useEffect } from 'react'
import { useParams, useNavigate, Link } from 'react-router-dom'
import { postsAPI } from '../services/api'
import './PostDetail.css'

function PostDetail() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [post, setPost] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    fetchPost()
  }, [id])

  const fetchPost = async () => {
    try {
      setLoading(true)
      const response = await postsAPI.getPost(id)
      setPost(response.data)
      setError(null)
    } catch (err) {
      setError('Failed to fetch post')
      console.error('Error fetching post:', err)
    } finally {
      setLoading(false)
    }
  }

  const handleDelete = async () => {
    if (!window.confirm('Are you sure you want to delete this post?')) {
      return
    }

    try {
      await postsAPI.deletePost(id)
      navigate('/')
    } catch (err) {
      alert('Failed to delete post')
      console.error('Error deleting post:', err)
    }
  }

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    })
  }

  if (loading) {
    return <div className="loading">Loading post...</div>
  }

  if (error || !post) {
    return (
      <div className="error-container">
        <div className="error">{error || 'Post not found'}</div>
        <Link to="/" className="btn btn-primary">Back to Posts</Link>
      </div>
    )
  }

  return (
    <div className="post-detail-container">
      <div className="post-detail-header">
        <Link to="/" className="back-link">← Back to Posts</Link>
        
        <div className="post-detail-actions">
          <Link to={`/edit/${post.id}`} className="btn btn-edit">
            Edit Post
          </Link>
          <button onClick={handleDelete} className="btn btn-delete">
            Delete Post
          </button>
        </div>
      </div>

      <article className="post-detail">
        <div className="post-detail-meta">
          <span className={`badge badge-${post.type}`}>{post.type}</span>
          <span className={`badge badge-${post.status}`}>{post.status}</span>
        </div>

        <h1 className="post-detail-title">{post.title}</h1>

        <div className="post-detail-info">
          <div className="info-item">
            <span className="info-icon">👤</span>
            <span>Author #{post.author_id}</span>
          </div>
          <div className="info-item">
            <span className="info-icon">📅</span>
            <span>Created: {formatDate(post.created_at)}</span>
          </div>
          <div className="info-item">
            <span className="info-icon">✏️</span>
            <span>Updated: {formatDate(post.updated_at)}</span>
          </div>
        </div>

        <div className="post-detail-content">
          {post.content.split('\n').map((paragraph, index) => (
            <p key={index}>{paragraph}</p>
          ))}
        </div>
      </article>
    </div>
  )
}

export default PostDetail
