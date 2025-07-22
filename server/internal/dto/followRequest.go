package dto

type FollowRequest struct {
	UserID string `json:"user_id" binding:"required"`
}
