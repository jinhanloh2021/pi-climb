package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
	"gorm.io/gorm"
)

type FollowService interface {
	CreateFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) (*models.Follow, error)
	DeleteFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) error
	GetFollowers(c context.Context, userID uuid.UUID, targetUserID uuid.UUID) ([]models.Follow, error)
	GetFollowing(c context.Context, userID uuid.UUID, targetUserID uuid.UUID) ([]models.Follow, error)
	GetFollowRelationship(c context.Context, userID uuid.UUID, targetUserID uuid.UUID) (*models.Follow, *models.Follow, error)
}

type followService struct {
	followRepo repository.FollowRepository
}

func NewFollowService(r repository.FollowRepository) FollowService {
	return &followService{followRepo: r}
}

func (s *followService) CreateFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) (*models.Follow, error) {
	follow, err := s.followRepo.CreateFollow(c, fromUserID, toUserID)
	if err != nil {
		return nil, err
	}
	return follow, nil
}
func (s *followService) DeleteFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) error {
	err := s.followRepo.DeleteFollow(c, fromUserID, toUserID)
	if err != nil {
		return err
	}
	return nil
}

func (s *followService) GetFollowers(c context.Context, userID uuid.UUID, targetUserID uuid.UUID) ([]models.Follow, error) {
	followers, err := s.followRepo.GetFollowers(c, userID, targetUserID)
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func (s *followService) GetFollowing(c context.Context, userID uuid.UUID, targetUserID uuid.UUID) ([]models.Follow, error) {
	following, err := s.followRepo.GetFollowing(c, userID, targetUserID)
	if err != nil {
		return nil, err
	}
	return following, nil
}

func (s *followService) GetFollowRelationship(c context.Context, userID uuid.UUID, targetUserID uuid.UUID) (*models.Follow, *models.Follow, error) {
	var userToTarget *models.Follow
	var targetToUser *models.Follow
	var combinedErr error

	userToTarget, err := s.followRepo.GetFollowEdge(c, userID, userID, targetUserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		combinedErr = errors.Join(combinedErr, err)
	}

	targetToUser, err = s.followRepo.GetFollowEdge(c, userID, targetUserID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		combinedErr = errors.Join(combinedErr, err)
	}

	return userToTarget, targetToUser, combinedErr
}
