package dto

type CreateMediaDto struct {
	URL           string `json:"url" binding:"required,url"`
	StoragePath   string `json:"storage_path"`
	ThumbnailURL  string `json:"thumbnail_url"`
	CompressedURL string `json:"compressed_url"`

	Filename string `json:"file_name" binding:"required"`
	FileSize int64  `json:"file_size" binding:"required"`
	MimeType string `json:"mime_type" binding:"required"`
	Order    *int   `json:"order"`

	Width    *int `json:"width"`
	Height   *int `json:"height"`
	Duration *int `json:"duration"`
}
