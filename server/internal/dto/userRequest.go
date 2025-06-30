package dto

import (
	"time"

	"github.com/google/uuid"
)

type UpdateDOBRequest struct {
	DateOfBirth  time.Time `json:"date_of_birth"`
	TargetUserID uuid.UUID `json:"target_user_id"`
}
