package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindBySupabaseID(c context.Context, callerID uuid.UUID, supabaseID uuid.UUID) (*models.User, error)
	SetDOBBySupabaseID(c context.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error)
	FindByUsername(c context.Context, username string, userUUID uuid.UUID) (*models.User, error)
	UpdateUser(c context.Context, userID uuid.UUID, body *dto.UpdateUserRequest) (*models.User, error)
}

type userRepository struct {
	*BaseRepository
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{BaseRepository: NewBaseRepository(db)}
}

// TODO: Logic for private/public profiles, check if following
func (r *userRepository) FindBySupabaseID(c context.Context, callerID uuid.UUID, supabaseID uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.withRLSTransaction(c, callerID, func(tx *gorm.DB) error {
		findResult := tx.Where("supabase_id = ?", supabaseID).First(&user)
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
func (r *userRepository) FindByUsername(c context.Context, username string, userUUID uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.withRLSTransaction(c, userUUID, func(tx *gorm.DB) error {
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

	// no updates
	if len(updates) == 0 {
		var currentUser models.User
		err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
			if findErr := tx.Where("supabase_id = ?", userID).First(&currentUser).Error; findErr != nil {
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
		updateResult := tx.Model(&models.User{}).Where("supabase_id = ?", userID).Updates(updates)
		if updateResult.Error != nil {
			return updateResult.Error
		}
		if updateResult.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return tx.Where("supabase_id = ?", userID).First(&user).Error
	})

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) SetDOBBySupabaseID(c context.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error) {
	var user models.User
	err := r.withRLSTransaction(c, callerID, func(tx *gorm.DB) error {
		// Find the user to update within this RLS-enabled transaction.
		findResult := tx.Where("supabase_id = ?", targetID).First(&user)
		if findResult.Error != nil {
			if errors.Is(findResult.Error, gorm.ErrRecordNotFound) {
				return gorm.ErrRecordNotFound
			}
			return findResult.Error
		}

		// Update the DOB
		updateResult := tx.Model(&user).Update("date_of_birth", DOB)
		if updateResult.Error != nil {
			return updateResult.Error
		}

		// If 0 rows affected, it means the update didn't occur for reasons
		// 1. The record didn't exist (already handled by `First()` above)
		// 2. The RLS policy prevented the update
		// 3. No actual change was made, DOB same
		if updateResult.RowsAffected == 0 {
			return fmt.Errorf("update failed for %s, possibly due to RLS policy or no changes applied", targetID)
		}

		return nil // No error, return nil error
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}
