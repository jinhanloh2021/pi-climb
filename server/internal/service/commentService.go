package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/pi-climb/internal/dto"
	"github.com/jinhanloh2021/pi-climb/internal/models"
	"github.com/jinhanloh2021/pi-climb/internal/repository"
)

type CommentService interface {
	GetComments(c context.Context, postID uint, userID uuid.UUID) ([]models.Comment, error)
	CreateComment(c context.Context, postID uint, body *dto.CreateCommentRequest, userID uuid.UUID) (*models.Comment, error)
	DeleteComment(c context.Context, commentID uint, userID uuid.UUID) error
}

type commentService struct {
	commentRepo repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) CommentService {
	return &commentService{commentRepo: r}
}

func (s *commentService) GetComments(c context.Context, postID uint, userID uuid.UUID) ([]models.Comment, error) {
	comments, err := s.commentRepo.GetComments(c, postID, userID)
	if err != nil {
		return nil, err
	}
	return comments, nil
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
