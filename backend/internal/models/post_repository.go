package models

// PostRepositoryInterface defines the contract for post repository operations
// This interface enables the Proxy pattern by allowing different implementations
type PostRepositoryInterface interface {
	Create(post *Post) error
	FindByID(id int64) (*Post, error)
	FindAll(status, contentType string, limit, offset int) ([]*Post, error)
	Update(post *Post) error
	Delete(id int64) error
}
