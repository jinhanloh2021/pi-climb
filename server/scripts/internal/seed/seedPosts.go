package seed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/scripts/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostSeedInput struct {
	UserID     string   `json:"user_id"`
	Caption    string   `json:"caption"`
	HoldColour string   `json:"hold_colour"`
	Grade      string   `json:"grade"`
	Media      []string `json:"media"`
}

func SeedPosts() {
	file, err := os.Open("./scripts/data/postsData.json")
	const bucket string = "media"
	if err != nil {
		log.Fatalf("Failed to open follows data file: %v", err)
	}
	defer file.Close()

	var posts []PostSeedInput
	if err := json.NewDecoder(file).Decode(&posts); err != nil {
		log.Fatalf("Failed to decode posts data: %v", err)
	}

	seedConfig := config.LoadSeedConfig()
	dsn := seedConfig.DbURLPostgresRole
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Using image path, get mimeType, filesize, width, height
	// Generate filename and upload to Supabase Storage with URL
	// Get storage key
	// Post to BE

	fmt.Printf("Seeding %d posts...\n", len(posts))
	for i, p := range posts {
		userID, err := GetUserIDByEmail(db, p.UserID)
		if err != nil {
			log.Printf("Skipping post: cannot find user %s: %v", p.UserID, err)
			continue
		}

		var postMedia []models.Media
		for j, m := range p.Media {
			path := "./scripts/data/img/" + m
			mimeType, fileSize, width, height, err := extractImageMetadata(path)
			if err != nil {
				log.Printf("Skipping media: Failed to extract metadata from %s: %v", m, err)
				continue
			}
			storageKey, err := uploadToSupabaseStorage(userID, path, bucket, m)
			if err != nil {
				log.Printf("Skipping media: Failed to upload %s: %v", m, err)
				continue
			}
			postMedia = append(postMedia, models.Media{
				StorageKey:   storageKey,
				Bucket:       bucket,
				OriginalName: m,
				FileSize:     fileSize,
				MimeType:     mimeType,
				Order:        &j,
				Width:        &width,
				Height:       &height,
				UserID:       userID,
			})
		}
		if len(p.Media) != len(postMedia) {
			log.Printf("Skipping post: Failed to upload media for post %d", i)
			continue
		}

		err = db.Transaction(func(tx *gorm.DB) error {
			post := models.Post{
				Caption:    &p.Caption,
				HoldColour: &p.HoldColour,
				Grade:      &p.Grade,
				UserID:     userID,
				Media:      postMedia, // auto creates Media with poly association
			}
			if err := tx.Create(&post).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Printf("Skipping post: Failed to create post for %s", p.UserID)
		} else {
			log.Printf("Seeded post for %s\n", p.UserID)
		}
	}
}

// Done by FE, then StorageKey sent to BE to create Media object
func uploadToSupabaseStorage(userID uuid.UUID, filePath, bucket, filename string) (string, error) {
	log.Printf("Uploading %s to bucket %s", filename, bucket)

	seedConfig := config.LoadSeedConfig()
	supabaseURL := seedConfig.SupabaseURL
	serviceRoleKey := seedConfig.ServiceRoleKey

	filenameParts := strings.Split(filename, ".")
	ext := filenameParts[len(filenameParts)-1]
	genFilename := fmt.Sprintf("%s.%s", uuid.New().String(), ext)

	storageKey := fmt.Sprintf("%s/%s", userID.String(), genFilename)
	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", supabaseURL, bucket, storageKey)

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("PUT", uploadURL, bytes.NewReader(fileBytes))
	if err != nil {
		return "", err
	}
	// will be logged-in user for FE
	req.Header.Set("Authorization", "Bearer "+serviceRoleKey)
	req.Header.Set("apikey", serviceRoleKey)
	req.Header.Set("Content-Type", "application/octet-stream")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("upload failed: %s", string(body))
	}

	return storageKey, nil
}

// for .jpg and .png
func extractImageMetadata(filePath string) (mimeType string, fileSize int64, width, height int, err error) {
	f, err := os.Open(filePath)
	f.Seek(0, io.SeekStart)
	if err != nil {
		return
	}
	defer f.Close()

	stat, _ := f.Stat()
	fileSize = stat.Size()

	img, format, err := image.DecodeConfig(f)
	if err != nil {
		return
	}
	width = img.Width
	height = img.Height

	mimeType = mime.TypeByExtension(filepath.Ext(filePath))
	if mimeType == "" {
		mimeType = "image/" + format
	}
	return
}
