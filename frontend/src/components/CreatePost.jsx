import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { postsAPI } from '../services/api'
import PostForm from './PostForm'

function CreatePost() {
  const navigate = useNavigate()
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  const initialData = {
    title: '',
    content: '',
    type: 'article',
    author_id: 1,
    status: 'draft'
  }

  const handleSubmit = async (postData) => {
    try {
      setLoading(true)
      setError(null)
      await postsAPI.createPost(postData)
      navigate('/')
    } catch (err) {
      setError('Failed to create post. Please try again.')
      console.error('Error creating post:', err)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="create-post-container">
      <h1>Create New Post</h1>
      {error && <div className="error-message">{error}</div>}
      <PostForm 
        initialData={initialData}
        onSubmit={handleSubmit}
        loading={loading}
        submitLabel="Create Post"
      />
    </div>
  )
}

export default CreatePost
