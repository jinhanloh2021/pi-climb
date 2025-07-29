package dto

type CreateCommentRequest struct {
	Text *string `json:"text" binding:"required"`
}
