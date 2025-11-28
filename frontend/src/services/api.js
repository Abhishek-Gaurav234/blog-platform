import axios from 'axios';

const API_BASE_URL = '/api/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Posts API
export const postsAPI = {
  // Get all posts
  getAllPosts: (params = {}) => {
    return api.get('/posts', { params });
  },

  // Get a single post by ID
  getPost: (id) => {
    return api.get(`/posts/${id}`);
  },

  // Create a new post
  createPost: (postData) => {
    return api.post('/posts', postData);
  },

  // Update a post
  updatePost: (id, postData) => {
    return api.put(`/posts/${id}`, postData);
  },

  // Delete a post
  deletePost: (id) => {
    return api.delete(`/posts/${id}`);
  },

  // Search posts
  searchPosts: (query) => {
    return api.get('/posts/search', { params: { q: query } });
  },
};

export default api;
