package dto

import (
	"time"

	"github.com/google/uuid"
)

type UpdateDOBRequest struct {
	DateOfBirth  time.Time `json:"date_of_birth"`
	TargetUserID uuid.UUID `json:"target_user_id"`
}

type UpdateUserRequest struct {
	Username    *string    `json:"username"`
	Bio         *string    `json:"bio"`
	IsPublic    *bool      `json:"is_public"`
	DateOfBirth *time.Time `json:"date_of_birth"`
}
