package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/pi-climb/internal/models"
	"github.com/jinhanloh2021/pi-climb/internal/repository"
)

type LikeService interface {
	CreateLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error)
	DeleteLike(c context.Context, userID uuid.UUID, postID uint) error
	GetPostLikes(c context.Context, userID uuid.UUID, postID uint) ([]models.Like, error)
	GetMyPostLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error)
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
	err := s.LikeRepo.DeleteLike(c, userID, postID)
	if err != nil {
		return err
	}
	return nil
}

func (s *likeService) GetPostLikes(c context.Context, userID uuid.UUID, postID uint) ([]models.Like, error) {
	likes, err := s.LikeRepo.GetPostLikes(c, userID, postID)
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func (s *likeService) GetMyPostLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error) {
	like, err := s.LikeRepo.GetMyPostLike(c, userID, postID)
	if err != nil {
		return nil, err
	}
	return like, nil
}
