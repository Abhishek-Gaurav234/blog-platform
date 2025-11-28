package proxy

import (
	"sync"
	"time"

	"blog-platform/internal/models"
)

// CacheEntry represents a cached post with expiration
type CacheEntry struct {
	Post      *models.Post
	ExpiresAt time.Time
}

// CacheStatistics tracks cache performance metrics
type CacheStatistics struct {
	Hits         int64
	Misses       int64
	Evictions    int64
	CurrentSize  int
	MaxSize      int
	HitRate      float64
}

// PostRepositoryCachingProxy implements the Proxy design pattern
// It wraps the real PostRepository and adds a caching layer
// Use case: Reduces database queries for frequently accessed posts
type PostRepositoryCachingProxy struct {
	realRepository models.PostRepositoryInterface
	cache          map[int64]*CacheEntry
	cacheMutex     sync.RWMutex
	maxCacheSize   int
	cacheTTL       time.Duration
	
	// Statistics
	hits      int64
	misses    int64
	evictions int64
	statsMutex sync.RWMutex
}

// NewPostRepositoryCachingProxy creates a new caching proxy
// Parameters:
//   - realRepo: The actual repository implementation to wrap
//   - maxSize: Maximum number of posts to cache (LRU eviction)
//   - ttl: Time-to-live for cached entries
func NewPostRepositoryCachingProxy(realRepo models.PostRepositoryInterface, maxSize int, ttl time.Duration) *PostRepositoryCachingProxy {
	return &PostRepositoryCachingProxy{
		realRepository: realRepo,
		cache:          make(map[int64]*CacheEntry),
		maxCacheSize:   maxSize,
		cacheTTL:       ttl,
	}
}

// FindByID implements transparent caching for post retrieval
// First checks cache, falls back to database on miss
func (p *PostRepositoryCachingProxy) FindByID(id int64) (*models.Post, error) {
	// Try to get from cache first
	p.cacheMutex.RLock()
	entry, exists := p.cache[id]
	p.cacheMutex.RUnlock()

	if exists && time.Now().Before(entry.ExpiresAt) {
		// Cache hit!
		p.recordHit()
		return entry.Post, nil
	}

	// Cache miss - fetch from real repository
	p.recordMiss()
	post, err := p.realRepository.FindByID(id)
	if err != nil || post == nil {
		return post, err
	}

	// Store in cache
	p.addToCache(id, post)
	return post, nil
}

// Create passes through to real repository and invalidates cache
func (p *PostRepositoryCachingProxy) Create(post *models.Post) error {
	err := p.realRepository.Create(post)
	if err == nil {
		// Add newly created post to cache
		p.addToCache(post.ID, post)
	}
	return err
}

// Update passes through and invalidates cache entry
func (p *PostRepositoryCachingProxy) Update(post *models.Post) error {
	err := p.realRepository.Update(post)
	if err == nil {
		// Invalidate cache for this post
		p.invalidateCache(post.ID)
	}
	return err
}

// Delete passes through and invalidates cache entry
func (p *PostRepositoryCachingProxy) Delete(id int64) error {
	err := p.realRepository.Delete(id)
	if err == nil {
		// Remove from cache
		p.invalidateCache(id)
	}
	return err
}

// FindAll passes through to real repository
// Note: List operations are typically not cached due to varying filters
func (p *PostRepositoryCachingProxy) FindAll(status, contentType string, limit, offset int) ([]*models.Post, error) {
	return p.realRepository.FindAll(status, contentType, limit, offset)
}

// addToCache adds or updates a cache entry with LRU eviction
func (p *PostRepositoryCachingProxy) addToCache(id int64, post *models.Post) {
	p.cacheMutex.Lock()
	defer p.cacheMutex.Unlock()

	// Check if we need to evict (simple LRU: remove oldest)
	if len(p.cache) >= p.maxCacheSize {
		// Find oldest entry
		var oldestID int64
		oldestTime := time.Now()
		for cacheID, entry := range p.cache {
			if entry.ExpiresAt.Before(oldestTime) {
				oldestTime = entry.ExpiresAt
				oldestID = cacheID
			}
		}
		delete(p.cache, oldestID)
		p.recordEviction()
	}

	// Add new entry
	p.cache[id] = &CacheEntry{
		Post:      post,
		ExpiresAt: time.Now().Add(p.cacheTTL),
	}
}

// invalidateCache removes an entry from cache
func (p *PostRepositoryCachingProxy) invalidateCache(id int64) {
	p.cacheMutex.Lock()
	defer p.cacheMutex.Unlock()
	delete(p.cache, id)
}

// ClearCache removes all cached entries
func (p *PostRepositoryCachingProxy) ClearCache() {
	p.cacheMutex.Lock()
	defer p.cacheMutex.Unlock()
	p.cache = make(map[int64]*CacheEntry)
}

// GetStatistics returns cache performance metrics
func (p *PostRepositoryCachingProxy) GetStatistics() CacheStatistics {
	p.statsMutex.RLock()
	defer p.statsMutex.RUnlock()
	
	p.cacheMutex.RLock()
	currentSize := len(p.cache)
	p.cacheMutex.RUnlock()

	total := p.hits + p.misses
	hitRate := 0.0
	if total > 0 {
		hitRate = float64(p.hits) / float64(total) * 100
	}

	return CacheStatistics{
		Hits:        p.hits,
		Misses:      p.misses,
		Evictions:   p.evictions,
		CurrentSize: currentSize,
		MaxSize:     p.maxCacheSize,
		HitRate:     hitRate,
	}
}

// recordHit increments cache hit counter
func (p *PostRepositoryCachingProxy) recordHit() {
	p.statsMutex.Lock()
	defer p.statsMutex.Unlock()
	p.hits++
}

// recordMiss increments cache miss counter
func (p *PostRepositoryCachingProxy) recordMiss() {
	p.statsMutex.Lock()
	defer p.statsMutex.Unlock()
	p.misses++
}

// recordEviction increments eviction counter
func (p *PostRepositoryCachingProxy) recordEviction() {
	p.statsMutex.Lock()
	defer p.statsMutex.Unlock()
	p.evictions++
}
