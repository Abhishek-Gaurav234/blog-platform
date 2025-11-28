package service

import (
	"time"

	"blog-platform/internal/models"
)

type GetPostQuery struct {
	ID int64
}

type ListPostsQuery struct {
	Status string
	Type   string
	Limit  int
	Offset int
}

type PostViewModel struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Type      string    `json:"type"`
	AuthorID  int64     `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}

type QueryService struct {
	postRepo models.PostRepositoryInterface
}

func NewQueryService(postRepo models.PostRepositoryInterface) *QueryService {
	return &QueryService{postRepo: postRepo}
}

func (s *QueryService) GetPost(query GetPostQuery) (*PostViewModel, error) {
	post, err := s.postRepo.FindByID(query.ID)
	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, nil
	}

	return &PostViewModel{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Type:      post.Type,
		AuthorID:  post.AuthorID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Status:    post.Status,
	}, nil
}

func (s *QueryService) ListPosts(query ListPostsQuery) ([]PostViewModel, error) {
	posts, err := s.postRepo.FindAll(query.Status, query.Type, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	viewModels := make([]PostViewModel, len(posts))
	for i, post := range posts {
		viewModels[i] = PostViewModel{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			Type:      post.Type,
			AuthorID:  post.AuthorID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Status:    post.Status,
		}
	}

	return viewModels, nil
}
