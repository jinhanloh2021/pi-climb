package dto

type CreateMediaDto struct {
	StorageKey string `json:"storage_key" binding:"required,url"`
	Bucket     string `json:"bucket" binding:"required"`

	OriginalName string `json:"file_name" binding:"required"`
	FileSize     int64  `json:"file_size" binding:"required"`

	MimeType string `json:"mime_type"`
	Order    *int   `json:"order"`

	Width    *int `json:"width"`
	Height   *int `json:"height"`
	Duration *int `json:"duration"`
}
