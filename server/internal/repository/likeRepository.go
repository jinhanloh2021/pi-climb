package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

type LikeRepository interface {
	CreateLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error)
	DeleteLike(c context.Context, userID uuid.UUID, postID uint) error
	GetLikes(c context.Context, userID uuid.UUID, postID uint) ([]models.Like, error)
	GetMyLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error)
}

type likeRepository struct {
	*BaseRepository
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepository{BaseRepository: NewBaseRepository(db)}
}

var ErrAlreadyLiked = fmt.Errorf("user already like this post")

func (r *likeRepository) CreateLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error) {
	like := models.Like{
		UserID: userID,
		PostID: postID,
	}
	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		err := tx.Unscoped().First(&like, "user_id = ? and post_id = ?", userID, postID).Error
		if err == nil {
			if like.DeletedAt.Valid {
				like.DeletedAt = gorm.DeletedAt{}
				return tx.Unscoped().Save(&like).Error
			}
			return ErrAlreadyLiked
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err = tx.Create(&like).Error; err != nil {
			return fmt.Errorf("failed to create follow")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &like, nil
}

func (r *likeRepository) DeleteLike(c context.Context, userID uuid.UUID, postID uint) error {
	return nil
}

func (r *likeRepository) GetLikes(c context.Context, userID uuid.UUID, postID uint) ([]models.Like, error) {
	return nil, nil
}

func (r *likeRepository) GetMyLike(c context.Context, userID uuid.UUID, postID uint) (*models.Like, error) {
	return nil, nil
}
