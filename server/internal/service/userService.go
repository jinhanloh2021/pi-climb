package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
)

type UserService interface {
	GetUserByUUID(c context.Context, callerID uuid.UUID, supabaseID uuid.UUID) (*models.User, error)
	GetUserByUsername(c context.Context, username string, userUUID uuid.UUID) (*models.User, error)
	UpdateUser(c context.Context, userID uuid.UUID, body *dto.UpdateUserRequest) (*models.User, error)
	SetUserDOB(c context.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) GetUserByUUID(c context.Context, callerID uuid.UUID, supabaseID uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.FindBySupabaseID(c, callerID, supabaseID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) SetUserDOB(c context.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error) {
	user, err := s.userRepo.SetDOBBySupabaseID(c, targetID, callerID, DOB)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByUsername(c context.Context, username string, userUUID uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.FindByUsername(c, username, userUUID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(c context.Context, userID uuid.UUID, body *dto.UpdateUserRequest) (*models.User, error) {

	user, err := s.userRepo.UpdateUser(c, userID, body)
	if err != nil {
		return nil, err
	}
	return user, nil
}
