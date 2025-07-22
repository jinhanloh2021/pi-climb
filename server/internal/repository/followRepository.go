package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

type FollowRepository interface {
	CreateFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) (*models.Follow, error)
	DeleteFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) error
}

type followRepository struct {
	*BaseRepository
}

func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepository{BaseRepository: NewBaseRepository(db)}
}

func (r *followRepository) CreateFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) (*models.Follow, error) {
	follow := models.Follow{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
	}
	err := r.withRLSTransaction(c, fromUserID, func(tx *gorm.DB) error {
		if err := tx.Create(&follow).Error; err != nil {
			return fmt.Errorf("failed to create follow: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &follow, nil
}

func (r *followRepository) DeleteFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) error {
	follow := models.Follow{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
	}
	err := r.withRLSTransaction(c, fromUserID, func(tx *gorm.DB) error {
		if err := tx.Delete(&follow).Error; err != nil {
			return fmt.Errorf("failed to delete follow: %w", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
