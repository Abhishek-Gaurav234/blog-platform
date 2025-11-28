import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { postsAPI } from '../services/api'
import PostForm from './PostForm'

function EditPost() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [post, setPost] = useState(null)
  const [loading, setLoading] = useState(false)
  const [fetchLoading, setFetchLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    fetchPost()
  }, [id])

  const fetchPost = async () => {
    try {
      setFetchLoading(true)
      const response = await postsAPI.getPost(id)
      setPost(response.data)
      setError(null)
    } catch (err) {
      setError('Failed to fetch post')
      console.error('Error fetching post:', err)
    } finally {
      setFetchLoading(false)
    }
  }

  const handleSubmit = async (postData) => {
    try {
      setLoading(true)
      setError(null)
      await postsAPI.updatePost(id, postData)
      navigate(`/posts/${id}`)
    } catch (err) {
      setError('Failed to update post. Please try again.')
      console.error('Error updating post:', err)
    } finally {
      setLoading(false)
    }
  }

  if (fetchLoading) {
    return <div className="loading">Loading post...</div>
  }

  if (error && !post) {
    return (
      <div className="error-container">
        <div className="error">{error}</div>
      </div>
    )
  }

  return (
    <div className="edit-post-container">
      <h1>Edit Post</h1>
      {error && <div className="error-message">{error}</div>}
      {post && (
        <PostForm 
          initialData={post}
          onSubmit={handleSubmit}
          loading={loading}
          submitLabel="Update Post"
        />
      )}
    </div>
  )
}

export default EditPost
