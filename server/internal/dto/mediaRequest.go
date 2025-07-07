package dto

type CreateMediaDto struct {
	URL       string `json:"url" binding:"required,url"`
	MediaType string `json:"media_type" binding:"required"`
	Order     *int   `json:"order"`
}
