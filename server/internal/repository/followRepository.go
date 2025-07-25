package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

type FollowRepository interface {
	CreateFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) (*models.Follow, error)
	DeleteFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) error
	GetFollowers(c context.Context, userID uuid.UUID, targetUserID uuid.UUID) ([]models.Follow, error)
	GetFollowing(c context.Context, userID uuid.UUID, targetUserID uuid.UUID) ([]models.Follow, error)
}

type followRepository struct {
	*BaseRepository
}

func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepository{BaseRepository: NewBaseRepository(db)}
}

var ErrAlreadyFollowing = fmt.Errorf("user is already following the target")

func (r *followRepository) CreateFollow(c context.Context, fromUserID uuid.UUID, toUserID uuid.UUID) (*models.Follow, error) {
	follow := models.Follow{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
	}
	err := r.withRLSTransaction(c, fromUserID, func(tx *gorm.DB) error {
		// check for soft deleted row, restore if soft deleted
		err := tx.Unscoped().Where("from_user_id = ? AND to_user_id = ?", fromUserID, toUserID).First(&follow).Error
		if err == nil {
			// soft deleted row found
			if follow.DeletedAt.Valid {
				follow.DeletedAt = gorm.DeletedAt{}
				return tx.Unscoped().Save(&follow).Error
			}
			// active row found
			return ErrAlreadyFollowing
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// create new row
		if err = tx.Create(&follow).Error; err != nil {
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

func (r *followRepository) GetFollowers(c context.Context, userID uuid.UUID, targetUserID uuid.UUID) ([]models.Follow, error) {
	var followers []models.Follow
	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		if err := tx.Select("from_user_id").Preload("FromUser").Where("to_user_id = ?", targetUserID).Find(&followers).Error; err != nil {
			return fmt.Errorf("failed to find followers of %s: %w", targetUserID, err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func (r *followRepository) GetFollowing(c context.Context, userID uuid.UUID, targetUserID uuid.UUID) ([]models.Follow, error) {
	var following []models.Follow
	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		if err := tx.Select("to_user_id").Preload("ToUser").Where("from_user_id = ?", targetUserID).Find(&following).Error; err != nil {
			return fmt.Errorf("failed to find following of %s: %w", targetUserID, err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return following, nil
}
