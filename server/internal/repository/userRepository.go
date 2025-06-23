package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

// Defines the interface for user data operations
type UserRepository interface {
	GetUserBySupabaseID(ctx context.Context, supabaseID uuid.UUID) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	// Other CRUD operations for User model
}

type userRepository struct {
	db *gorm.DB
}

// Creates a new UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Retrieves a user by their Supabase ID
func (r *userRepository) GetUserBySupabaseID(ctx context.Context, supabaseID uuid.UUID) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("supabase_id = ?", supabaseID).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // User not found
		}
		return nil, result.Error
	}
	return &user, nil // User found
}

// Creates a new user record
func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
