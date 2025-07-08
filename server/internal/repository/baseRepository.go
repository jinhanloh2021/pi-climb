package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Contains common dependencies and methods for all repositories.
type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (b *BaseRepository) withRLSTransaction(ctx context.Context, userID uuid.UUID, fn func(tx *gorm.DB) error) error {
	return b.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Safe from SQL injection
		if err := tx.Exec(fmt.Sprintf("SET app.current_user_id = '%s'", userID.String())).Error; err != nil {
			return fmt.Errorf("Error setting app.current_user_id")
		}
		return fn(tx)
	})
}
