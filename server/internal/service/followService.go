package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
)

type FollowService interface {
	CreateFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) (*models.Follow, error)
	DeleteFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) error
}

type followService struct {
	FollowRepo repository.FollowRepository
}

func NewFollowService(r repository.FollowRepository) FollowService {
	return &followService{FollowRepo: r}
}

func (s *followService) CreateFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) (*models.Follow, error) {
	follow, err := s.FollowRepo.CreateFollow(c, fromUserID, toUserID)
	if err != nil {
		return nil, err
	}
	return follow, nil
}
func (s *followService) DeleteFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) error {
	err := s.FollowRepo.DeleteFollow(c, fromUserID, toUserID)
	if err != nil {
		return err
	}
	return nil
}
