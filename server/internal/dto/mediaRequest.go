package dto

type CreateMediaDto struct {
	StorageKey string `json:"storage_key" binding:"required"`
	Bucket     string `json:"bucket" binding:"required"`

	OriginalName string `json:"original_name" binding:"required"`
	FileSize     int64  `json:"file_size" binding:"required"`

	MimeType string `json:"mime_type"`
	Order    *uint  `json:"order"`

	Width    *uint `json:"width"`
	Height   *uint `json:"height"`
	Duration *uint `json:"duration"`
}
