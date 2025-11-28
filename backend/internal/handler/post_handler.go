package handler

import (
	"blog-platform/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	commandService *service.CommandService
	queryService   *service.QueryService
	contentFactory *service.ContentFactory
	postService    *service.PostService
	searchService  *service.SearchService
}

func NewPostHandler(
	cmdService *service.CommandService,
	queryService *service.QueryService,
	contentFactory *service.ContentFactory,
	postService *service.PostService,
	searchService *service.SearchService,
) *PostHandler {
	return &PostHandler{
		commandService: cmdService,
		queryService:   queryService,
		contentFactory: contentFactory,
		postService:    postService,
		searchService:  searchService,
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var createCmd service.CreatePostCommand
	if err := c.ShouldBindJSON(&createCmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the content factory to create appropriate content type
	content, err := h.contentFactory.CreateContent(createCmd.Type, map[string]interface{}{
		"title":     createCmd.Title,
		"content":   createCmd.Content,
		"author_id": createCmd.AuthorID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the content
	if err := content.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the post
	post, err := h.commandService.CreatePost(createCmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Notify observers about the new post
	h.postService.Notify(service.PostEvent{
		EventType: "post_created",
		PostID:    post.ID,
		Data:      post,
	})

	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetPost(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := h.queryService.GetPost(service.GetPostQuery{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if post == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) ListPosts(c *gin.Context) {
	query := service.ListPostsQuery{
		Status: c.Query("status"),
		Type:   c.Query("type"),
		Limit:  10, // Default limit
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 && limit <= 100 {
			query.Limit = limit
		}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil && offset >= 0 {
			query.Offset = offset
		}
	}

	posts, err := h.queryService.ListPosts(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var updateCmd service.UpdatePostCommand
	if err := c.ShouldBindJSON(&updateCmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateCmd.ID = id

	// Update the post
	post, err := h.commandService.UpdatePost(updateCmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Notify observers about the updated post
	h.postService.Notify(service.PostEvent{
		EventType: "post_updated",
		PostID:    post.ID,
		Data:      post,
	})

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	deleteCmd := service.DeletePostCommand{ID: id}

	// Delete the post
	if err := h.commandService.DeletePost(deleteCmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Notify observers about the deleted post
	h.postService.Notify(service.PostEvent{
		EventType: "post_deleted",
		PostID:    id,
		Data:      nil,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func (h *PostHandler) SearchPosts(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	// Search posts using circuit breaker protected service
	results, err := h.searchService.SearchPosts(query)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error":          "Search service temporarily unavailable",
			"circuit_breaker": h.searchService.GetCircuitBreakerState(),
		})
		return
	}

	// Return results even if empty (circuit breaker may have returned fallback)
	c.JSON(http.StatusOK, gin.H{
		"query":   query,
		"count":   len(results),
		"results": results,
		"circuit_breaker": h.searchService.GetCircuitBreakerState(),
	})
}
