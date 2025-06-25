package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindBySupabaseID(ctx context.Context, supabaseID uuid.UUID) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindBySupabaseID(ctx context.Context, supabaseID uuid.UUID) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("supabase_id = ?", supabaseID).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
