package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
)

type PostService interface {
	CreatePost(c context.Context, userID uuid.UUID, body *dto.CreatePostRequest) (*models.Post, error)
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(r repository.PostRepository) PostService {
	return &postService{postRepo: r}
}

func (s *postService) CreatePost(c context.Context, userID uuid.UUID, body *dto.CreatePostRequest) (*models.Post, error) {
	post, err := s.postRepo.CreateNewPost(c, userID, body)
	if err != nil {
		return nil, err
	}
	return post, nil
}
