package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
)

type UserService interface {
	GetUserByUUID(c *gin.Context, supabaseID uuid.UUID) (*models.User, error)
	GetUserByUsername(c *gin.Context, username string, userUUID uuid.UUID) (*models.User, error)
	SetUserDOB(c *gin.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) GetUserByUUID(c *gin.Context, supabaseID uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.FindBySupabaseID(c, supabaseID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) SetUserDOB(c *gin.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error) {
	user, err := s.userRepo.SetDOBBySupabaseID(c, targetID, callerID, DOB)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByUsername(c *gin.Context, username string, userUUID uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.FindByUsername(c, username, userUUID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}
