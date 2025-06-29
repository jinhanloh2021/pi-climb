package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindBySupabaseID(c *gin.Context, supabaseID uuid.UUID) (*models.User, error)
	SetDOBBySupabaseID(c *gin.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error)
	FindByUsername(ctx *gin.Context, username string, userUUID uuid.UUID) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindBySupabaseID(c *gin.Context, supabaseID uuid.UUID) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(c).Where("supabase_id = ?", supabaseID).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

// TODO: Mask PII in response, except if owner
func (r *userRepository) FindByUsername(c *gin.Context, username string, userUUID uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(fmt.Sprintf("SET app.current_user_id = '%s'", userUUID.String())).Error; err != nil {
			return fmt.Errorf("failed to set RLS context for transaction: %w", err)
		}
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

func (r *userRepository) SetDOBBySupabaseID(c *gin.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(fmt.Sprintf("SET app.current_user_id = '%s'", callerID.String())).Error; err != nil {
			return errors.New("failed to set RLS context for transaction")
		}

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
			return errors.New("update failed, possibly due to RLS policy or no changes applied")
		}

		return nil // No error, return nil error
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}
