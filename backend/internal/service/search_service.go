package service

import (
	"blog-platform/internal/models"
	"blog-platform/pkg/circuitbreaker"
	"fmt"
	"strings"
)

// SearchService handles search operations with circuit breaker
type SearchService struct {
	postRepo       *models.PostRepository
	circuitBreaker *circuitbreaker.CircuitBreaker
}

// NewSearchService creates a new search service with circuit breaker
func NewSearchService(postRepo *models.PostRepository) *SearchService {
	return &SearchService{
		postRepo:       postRepo,
		circuitBreaker: circuitbreaker.NewCircuitBreaker("search", 5, 30),
	}
}

// SearchPosts searches posts by title or content with circuit breaker protection
func (s *SearchService) SearchPosts(query string) ([]PostViewModel, error) {
	var results []PostViewModel

	// Execute search with circuit breaker
	result, err := s.circuitBreaker.Execute(func() (interface{}, error) {
		return s.performSearch(query)
	})

	// Circuit breaker is open - return cached/fallback results
	if err != nil && strings.Contains(err.Error(), "circuit breaker is open") {
		return s.getFallbackResults(query), nil
	}

	// Search execution failed but circuit is closed/half-open
	if err != nil {
		return nil, err
	}

	// Successful search
	if result != nil {
		results = result.([]PostViewModel)
	}

	return results, nil
}

// performSearch executes the actual search against the database
func (s *SearchService) performSearch(query string) ([]PostViewModel, error) {
	// Get all posts (in a real implementation, this would be a SQL LIKE query)
	posts, err := s.postRepo.FindAll("", "", 100, 0)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	// Filter posts by search query (simple implementation)
	var results []PostViewModel
	searchLower := strings.ToLower(query)
	
	for _, post := range posts {
		titleMatch := strings.Contains(strings.ToLower(post.Title), searchLower)
		contentMatch := strings.Contains(strings.ToLower(post.Content), searchLower)
		
		if titleMatch || contentMatch {
			results = append(results, PostViewModel{
				ID:        post.ID,
				Title:     post.Title,
				Content:   post.Content,
				Type:      post.Type,
				AuthorID:  post.AuthorID,
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
				Status:    post.Status,
			})
		}
	}

	return results, nil
}

// getFallbackResults returns cached or default results when circuit breaker is open
func (s *SearchService) getFallbackResults(query string) []PostViewModel {
	// In a real implementation, this could return cached results or popular posts
	// For now, return an empty slice
	return []PostViewModel{}
}

// GetCircuitBreakerState returns the current state of the circuit breaker
func (s *SearchService) GetCircuitBreakerState() string {
	state := s.circuitBreaker.GetState()
	switch state {
	case circuitbreaker.StateClosed:
		return "closed"
	case circuitbreaker.StateOpen:
		return "open"
	case circuitbreaker.StateHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}
