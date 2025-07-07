package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreateNewPost(c context.Context, userID uuid.UUID, body *dto.CreatePostRequest) (*models.Post, error)
}

type postRepository struct {
	*BaseRepository
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{BaseRepository: NewBaseRepository(db)}
}

func (r *postRepository) CreateNewPost(c context.Context, userID uuid.UUID, body *dto.CreatePostRequest) (*models.Post, error) {
	var post *models.Post
	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		newPost := models.Post{
			Caption:    body.Caption,
			HoldColour: body.HoldColour,
			Grade:      body.Grade,
			GymID:      body.GymID,
		}
		var user models.User
		if findUserErr := tx.Select("id").Where("supabase_id = ?", userID).First(&user).Error; findUserErr != nil {
			if errors.Is(findUserErr, gorm.ErrRecordNotFound) {
				return gorm.ErrRecordNotFound
			}
			return fmt.Errorf("failed to find user: %w", findUserErr)
		}
		newPost.UserID = user.ID // Set the GORM UserID

		if createErr := tx.Create(&newPost).Error; createErr != nil {
			return fmt.Errorf("failed to create post: %w", createErr)
		}

		// 2. Create associated Media records
		if len(body.Media) > 0 {
			mediaRecords := make([]models.Media, len(body.Media))
			for i, mediaDto := range body.Media {

				mediaRecords[i] = models.Media{
					URL:       mediaDto.URL,
					MediaType: models.MediaType(mediaDto.MediaType),
					Order:     mediaDto.Order,
					OwnerID:   newPost.ID, // Link to the newly created Post
					OwnerType: "posts",    // Polymorphic association value
				}
			}

			if createMediaErr := tx.Create(&mediaRecords).Error; createMediaErr != nil {
				return fmt.Errorf("failed to create media for post: %w", createMediaErr)
			}
		}

		if findErr := tx.Preload("Media").Preload("User").Preload("Gym").Where("id = ?", newPost.ID).First(&newPost).Error; findErr != nil {
			return fmt.Errorf("failed to retrieve created post with associations: %w", findErr)
		}
		post = &newPost
		return nil
	})

	if err != nil {
		return nil, err
	}
	return post, nil
}
