package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(c context.Context, postID uint, body *dto.CreateCommentRequest, userID uuid.UUID) (*models.Comment, error)
	DeleteComment(c context.Context, commentID uint, userID uuid.UUID) error
}

type commentRepository struct {
	*BaseRepository
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{BaseRepository: NewBaseRepository(db)}
}

func (r *commentRepository) CreateComment(c context.Context, postID uint, body *dto.CreateCommentRequest, userID uuid.UUID) (*models.Comment, error) {
	comment := models.Comment{
		Text:   *body.Text,
		UserID: userID,
		PostID: postID,
	}
	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		if err := tx.Create(&comment).Error; err != nil {
			return fmt.Errorf("failed to create comment for post %d", postID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) DeleteComment(c context.Context, commentID uint, userID uuid.UUID) error {
	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Comment{}, commentID).Error; err != nil {
			return fmt.Errorf("failed to delete comment %d", commentID)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
