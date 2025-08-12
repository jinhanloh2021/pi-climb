package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	GetComments(c context.Context, postID uint, userID uuid.UUID) ([]models.Comment, error)
	CreateComment(c context.Context, postID uint, body *dto.CreateCommentRequest, userID uuid.UUID) (*models.Comment, error)
	DeleteComment(c context.Context, commentID uint, userID uuid.UUID) error
}

type commentRepository struct {
	*BaseRepository
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{BaseRepository: NewBaseRepository(db)}
}

var ErrCommentNotFound = errors.New("comment not found or not accessible")

func (r *commentRepository) GetComments(c context.Context, postID uint, userID uuid.UUID) ([]models.Comment, error) {
	comments := []models.Comment{}
	// todo: handle infinite scrolling, cursor based pagination for comments
	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		if err := tx.Preload("User").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
			return fmt.Errorf("failed to find comments for post %d", postID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return comments, nil
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
		res := tx.Delete(&models.Comment{}, commentID)
		if res.Error != nil {
			return fmt.Errorf("failed to delete comment %d", commentID)
		}
		if res.RowsAffected == 0 {
			return ErrCommentNotFound
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
