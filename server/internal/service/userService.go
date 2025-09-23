package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/pi-climb/internal/dto"
	"github.com/jinhanloh2021/pi-climb/internal/models"
	"github.com/jinhanloh2021/pi-climb/internal/repository"
)

type UserService interface {
	GetUserByID(c context.Context, userID uuid.UUID, targetID uuid.UUID) (*models.User, error)
	GetUserByUsername(c context.Context, username string, userID uuid.UUID) (*models.User, error)
	UpdateUser(c context.Context, userID uuid.UUID, body *dto.UpdateUserRequest) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) GetUserByID(c context.Context, userID uuid.UUID, targetID uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.FindByUserID(c, userID, targetID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByUsername(c context.Context, username string, userID uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.FindByUsername(c, username, userID)
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
