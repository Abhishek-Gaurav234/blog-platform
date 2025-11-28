import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import './PostForm.css'

function PostForm({ initialData, onSubmit, loading, submitLabel }) {
  const navigate = useNavigate()
  const [formData, setFormData] = useState(initialData)

  const handleChange = (e) => {
    const { name, value } = e.target
    setFormData(prev => ({
      ...prev,
      [name]: name === 'author_id' ? parseInt(value) : value
    }))
  }

  const handleSubmit = (e) => {
    e.preventDefault()
    onSubmit(formData)
  }

  return (
    <form onSubmit={handleSubmit} className="post-form">
      <div className="form-group">
        <label htmlFor="title">Title *</label>
        <input
          type="text"
          id="title"
          name="title"
          value={formData.title}
          onChange={handleChange}
          required
          placeholder="Enter post title"
          className="form-input"
        />
      </div>

      <div className="form-row">
        <div className="form-group">
          <label htmlFor="type">Type *</label>
          <select
            id="type"
            name="type"
            value={formData.type}
            onChange={handleChange}
            required
            className="form-select"
          >
            <option value="article">Article</option>
            <option value="tutorial">Tutorial</option>
            <option value="review">Review</option>
          </select>
        </div>

        <div className="form-group">
          <label htmlFor="status">Status *</label>
          <select
            id="status"
            name="status"
            value={formData.status}
            onChange={handleChange}
            required
            className="form-select"
          >
            <option value="draft">Draft</option>
            <option value="published">Published</option>
            <option value="archived">Archived</option>
          </select>
        </div>

        <div className="form-group">
          <label htmlFor="author_id">Author ID *</label>
          <input
            type="number"
            id="author_id"
            name="author_id"
            value={formData.author_id}
            onChange={handleChange}
            required
            min="1"
            className="form-input"
          />
        </div>
      </div>

      <div className="form-group">
        <label htmlFor="content">Content *</label>
        <textarea
          id="content"
          name="content"
          value={formData.content}
          onChange={handleChange}
          required
          placeholder="Write your post content here..."
          rows="15"
          className="form-textarea"
        />
      </div>

      <div className="form-actions">
        <button 
          type="button" 
          onClick={() => navigate(-1)} 
          className="btn btn-secondary"
          disabled={loading}
        >
          Cancel
        </button>
        <button 
          type="submit" 
          className="btn btn-primary"
          disabled={loading}
        >
          {loading ? 'Saving...' : submitLabel}
        </button>
      </div>
    </form>
  )
}

export default PostForm
