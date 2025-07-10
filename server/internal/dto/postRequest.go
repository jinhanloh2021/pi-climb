package dto

type CreatePostRequest struct {
	Caption    *string          `json:"caption" binding:"max=512"`
	HoldColour *string          `json:"hold_colour" binding:"max=64"`
	Grade      *string          `json:"grade" binding:"max=64"`
	Media      []CreateMediaDto `json:"media" binding:"required,dive"` // nested
	GymID      *uint            `json:"gym_id"`
}
