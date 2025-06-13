package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
	"gorm.io/gorm"
)

// UserService defines the interface for user business logic
type UserService interface {
	GetOrCreateUserBySupabaseID(ctx context.Context, supabaseID uuid.UUID, email string) (*models.User, error)
	// Other business logic related to users
}

type userService struct {
	userRepo repository.UserRepository
}

// Creates a new UserService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

// Retrieves a user by SupabaseID or creates a new profile for them
func (s *userService) GetOrCreateUserBySupabaseID(ctx context.Context, supabaseID uuid.UUID, email string) (*models.User, error) {
	user, err := s.userRepo.GetUserBySupabaseID(ctx, supabaseID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		// User profile does not exist, create a new one
		newUser := &models.User{
			SupabaseID: supabaseID,
			Email:      email,
			Username:   "user_" + supabaseID.String()[0:8], // auto-generated username
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		if err := s.userRepo.CreateUser(ctx, newUser); err != nil {
			return nil, err
		}
		return newUser, nil
	}
	return user, nil
}
