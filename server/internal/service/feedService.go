package service

import (
	"context"
	"slices"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
)

type FeedService interface {
	GetFeed(c context.Context, userID uuid.UUID, feedCursor *dto.FeedCursor) ([]models.Post, error)
}

type feedService struct {
	postRepo repository.PostRepository
}

func NewFeedService(r repository.PostRepository) FeedService {
	return &feedService{postRepo: r}
}

func (s *feedService) GetFeed(c context.Context, userID uuid.UUID, feedCursor *dto.FeedCursor) ([]models.Post, error) {
	followingFeed, err := s.postRepo.GetFollowingFeed(c, userID, feedCursor)
	if err != nil {
		return nil, err
	}
	trendingFeed, err := s.postRepo.GetTrendingFeed(c, userID, feedCursor)
	if err != nil {
		return nil, err
	}

	feed := slices.Concat(followingFeed, trendingFeed)

	return feed, nil
}
