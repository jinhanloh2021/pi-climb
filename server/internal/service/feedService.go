package service

import (
	"context"
	"slices"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/pi-climb/internal/dto"
	"github.com/jinhanloh2021/pi-climb/internal/models"
	"github.com/jinhanloh2021/pi-climb/internal/repository"
)

type FeedService interface {
	GetFeed(c context.Context, userID uuid.UUID, feedCursor *dto.FeedCursor, limit int) ([]models.Post, *dto.FeedCursor, error)
}

type feedService struct {
	postRepo repository.PostRepository
}

func NewFeedService(r repository.PostRepository) FeedService {
	return &feedService{postRepo: r}
}

func (s *feedService) GetFeed(c context.Context, userID uuid.UUID, feedCursor *dto.FeedCursor, limit int) ([]models.Post, *dto.FeedCursor, error) {
	followingFeed, nextFollowingCursor, err := s.postRepo.GetFollowingFeed(c, userID, feedCursor, limit)
	if err != nil {
		return nil, nil, err
	}
	trendingFeed, nextTrendingCursor, err := s.postRepo.GetTrendingFeed(c, userID, feedCursor, limit)
	if err != nil {
		return nil, nil, err
	}

	feed := slices.Concat(followingFeed, trendingFeed)

	return feed, &dto.FeedCursor{FollowingCursor: nextFollowingCursor, TrendingCursor: nextTrendingCursor}, nil
}
