package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
)

type CommentService interface {
	CreateComment(c context.Context, postID uint, body *dto.CreateCommentRequest, userID uuid.UUID) (*models.Comment, error)
	DeleteComment(c context.Context, commentID uint, userID uuid.UUID) error
}

type commentService struct {
	commentRepo repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) CommentService {
	return &commentService{commentRepo: r}
}

func (s *commentService) CreateComment(c context.Context, postID uint, body *dto.CreateCommentRequest, userID uuid.UUID) (*models.Comment, error) {
	comment, err := s.commentRepo.CreateComment(c, postID, body, userID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *commentService) DeleteComment(c context.Context, commentID uint, userID uuid.UUID) error {
	err := s.commentRepo.DeleteComment(c, commentID, userID)
	if err != nil {
		return err
	}
	return nil
}
