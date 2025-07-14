package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreateNewPost(c context.Context, userID uuid.UUID, body *dto.CreatePostRequest) (*models.Post, error)
	GetFollowingFeed(c context.Context, userID uuid.UUID, feedCursor *dto.FeedCursor) ([]models.Post, string, error)
	GetTrendingFeed(c context.Context, userID uuid.UUID, feedCursor *dto.FeedCursor) ([]models.Post, string, error)
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
			UserID:     userID,
			GymID:      body.GymID,
		}

		if createErr := tx.Create(&newPost).Error; createErr != nil {
			return fmt.Errorf("failed to create post: %w", createErr)
		}

		// 2. Create associated Media records
		if len(body.Media) > 0 {
			mediaRecords := make([]models.Media, len(body.Media))
			for i, mediaDto := range body.Media {
				mediaRecords[i] = models.Media{
					URL:           mediaDto.URL,
					StoragePath:   mediaDto.StoragePath,
					ThumbnailURL:  mediaDto.ThumbnailURL,
					CompressedURL: mediaDto.CompressedURL,

					Filename: mediaDto.Filename,
					FileSize: mediaDto.FileSize,
					MimeType: mediaDto.MimeType,
					Order:    mediaDto.Order,

					Width:    mediaDto.Width,
					Height:   mediaDto.Height,
					Duration: mediaDto.Duration,

					OwnerID:   newPost.ID, // Link to the newly created Post
					OwnerType: "posts",    // Polymorphic association value
					UserID:    userID,
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

func (r *postRepository) GetFollowingFeed(c context.Context, userID uuid.UUID, feedCursor *dto.FeedCursor) ([]models.Post, string, error) {
	var posts []models.Post
	var followingUserID []uuid.UUID
	var nextCursor string
	const defaultFeedLimit int = 20

	err := r.withRLSTransaction(c, userID, func(tx *gorm.DB) error {
		if findFollowingErr := tx.Model(&models.Follow{}).Select("to_user_id").Where("from_user_id = ?", userID).Find(&followingUserID).Error; findFollowingErr != nil {
			return fmt.Errorf("failed to find following users: %w", findFollowingErr)
		}
		if len(followingUserID) == 0 {
			return nil
		}
		query := tx.Preload("User").Preload("Media").Where("user_id IN (?)", followingUserID)
		cursorTimestampNano, cursorPostID := dto.ParsePostCursor(feedCursor.FollowingCursor)
		if cursorTimestampNano != -1 && cursorPostID != -1 {
			query = query.Where("(created_at < ?) OR (created_at = ? AND id < ?)",
				time.Unix(0, cursorTimestampNano), time.Unix(0, cursorTimestampNano), cursorPostID)
		}
		findPostsErr := query.Order("created_at DESC, id DESC").Limit(defaultFeedLimit).Find(&posts).Error
		if findPostsErr != nil {
			return findPostsErr
		}
		if len(posts) > 0 {
			lastPost := posts[len(posts)-1]
			nextCursor = dto.FormatPostCursor(&lastPost)
		}
		return nil
	})

	if err != nil {
		return nil, "", err
	}
	return posts, nextCursor, nil
}

func (r *postRepository) GetTrendingFeed(c context.Context, userID uuid.UUID, feedCursor *dto.FeedCursor) ([]models.Post, string, error) {
	// todo: Get posts of all users, sorted by likes or views
	return nil, "", nil
}
