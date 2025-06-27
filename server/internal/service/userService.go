package service

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
)

type UserService interface {
	// TODO: Change to use *gin.Context
	GetUserByUUID(ctx context.Context, supabaseID uuid.UUID) (*models.User, error)
	GetUserByUsername(c *gin.Context, username string) (*models.User, error)
	SetUserDOB(ctx context.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) GetUserByUUID(ctx context.Context, supabaseID uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.FindBySupabaseID(ctx, supabaseID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) SetUserDOB(ctx context.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error) {
	user, err := s.userRepo.SetDOBBySupabaseID(ctx, targetID, callerID, DOB)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByUsername(c *gin.Context, username string) (*models.User, error) {
	user, err := s.userRepo.FindByUsername(c, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}
