package service

import (
	"errors"

	"blog-platform/internal/models"
)

type CreatePostCommand struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Type     string `json:"type"`
	AuthorID int64  `json:"author_id"`
}

type UpdatePostCommand struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
	Status  string `json:"status"`
}

type DeletePostCommand struct {
	ID int64 `json:"id"`
}

type CommandService struct {
	postRepo *models.PostRepository
}

func NewCommandService(postRepo *models.PostRepository) *CommandService {
	return &CommandService{postRepo: postRepo}
}

func (s *CommandService) CreatePost(cmd CreatePostCommand) (*models.Post, error) {
	if cmd.Title == "" || cmd.Content == "" {
		return nil, errors.New("title and content are required")
	}

	post := &models.Post{
		Title:    cmd.Title,
		Content:  cmd.Content,
		Type:     cmd.Type,
		AuthorID: cmd.AuthorID,
		Status:   "draft",
	}

	err := s.postRepo.Create(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *CommandService) UpdatePost(cmd UpdatePostCommand) (*models.Post, error) {
	post, err := s.postRepo.FindByID(cmd.ID)
	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, errors.New("post not found")
	}

	if cmd.Title != "" {
		post.Title = cmd.Title
	}

	if cmd.Content != "" {
		post.Content = cmd.Content
	}

	if cmd.Type != "" {
		post.Type = cmd.Type
	}

	if cmd.Status != "" {
		post.Status = cmd.Status
	}

	err = s.postRepo.Update(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *CommandService) DeletePost(cmd DeletePostCommand) error {
	post, err := s.postRepo.FindByID(cmd.ID)
	if err != nil {
		return err
	}

	if post == nil {
		return errors.New("post not found")
	}

	return s.postRepo.Delete(cmd.ID)
}
