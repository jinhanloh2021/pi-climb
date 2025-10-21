package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/pi-climb/internal/dto"
	"github.com/jinhanloh2021/pi-climb/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUserID(c context.Context, userID uuid.UUID, targetID uuid.UUID) (*models.User, error)
	FindByUsername(c context.Context, username string, userID uuid.UUID) (*models.User, error)
	UpdateUser(c context.Context, userID uuid.UUID, body *dto.UpdateUserRequest) (*models.User, error)
}

type userRepository struct {
	*BaseRepository
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{BaseRepository: NewBaseRepository(db)}
}

// TODO: Logic for private/public profiles, check if following
func (r *userRepository) FindByUserID(c context.Context, userID uuid.UUID, targetID uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		findResult := tx.Where("id = ?", targetID).First(&user)
		if findResult.Error != nil {
			if errors.Is(findResult.Error, gorm.ErrRecordNotFound) {
				return gorm.ErrRecordNotFound
			}
			return fmt.Errorf("failed to find user: %w", findResult.Error)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// TODO: Mask PII in response, except if owner
func (r *userRepository) FindByUsername(c context.Context, username string, userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		findResult := tx.Where("username = ?", username).First(&user)
		if findResult.Error != nil {
			if errors.Is(findResult.Error, gorm.ErrRecordNotFound) {
				return gorm.ErrRecordNotFound
			}
			return fmt.Errorf("failed to find user: %w", findResult.Error)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(c context.Context, userID uuid.UUID, body *dto.UpdateUserRequest) (*models.User, error) {
	updates := make(map[string]any)
	if body.Username != nil {
		updates["username"] = *body.Username
	}
	if body.Bio != nil {
		updates["bio"] = *body.Bio
	}
	if body.IsPublic != nil {
		updates["is_public"] = *body.IsPublic
	}
	if body.DateOfBirth != nil {
		updates["date_of_birth"] = *body.DateOfBirth
	}

	// no updates, return unchanged user
	if len(updates) == 0 {
		var currentUser models.User
		err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
			if findErr := tx.Where("id = ?", userID).First(&currentUser).Error; findErr != nil {
				if errors.Is(findErr, gorm.ErrRecordNotFound) {
					return gorm.ErrRecordNotFound
				}
				return findErr
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		return &currentUser, nil
	}

	var user models.User
	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		updateResult := tx.Model(&models.User{}).Where("id = ?", userID).Updates(updates)
		if updateResult.Error != nil {
			return updateResult.Error
		}
		if updateResult.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return tx.Where("id = ?", userID).First(&user).Error
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}
