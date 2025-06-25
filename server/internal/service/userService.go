package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
)

type UserService interface {
	GetUserByUUID(ctx context.Context, supabaseID uuid.UUID) (*models.User, error)
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

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}
