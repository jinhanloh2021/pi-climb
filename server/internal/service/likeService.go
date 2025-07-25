package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
)

type LikeService interface {
	CreateLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error)
	DeleteLike(c context.Context, userID uuid.UUID, postID uint) error
	GetLikes(c context.Context, userID uuid.UUID, postID uint) ([]models.Like, error)
	GetMyLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error)
}

type likeService struct {
	LikeRepo repository.LikeRepository
}

func NewLikeService(r repository.LikeRepository) LikeService {
	return &likeService{LikeRepo: r}
}

func (s *likeService) CreateLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error) {
	like, err := s.LikeRepo.CreateLike(c, userID, postID)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (s *likeService) DeleteLike(c context.Context, userID uuid.UUID, postID uint) error {
	return nil
}

func (s *likeService) GetLikes(c context.Context, userID uuid.UUID, postID uint) ([]models.Like, error) {
	return nil, nil
}

func (s *likeService) GetMyLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error) {
	return nil, nil
}
