package service

import (
	"context"
	"fmt"

	"github.com/eydeveloper/highload-social/internal/entity"
	"github.com/eydeveloper/highload-social/internal/repository"
	"github.com/redis/go-redis/v9"
)

type FeedService struct {
	repo        repository.Post
	redisClient *redis.Client
}

func NewFeedService(repo repository.Post, redisClient *redis.Client) *FeedService {
	return &FeedService{
		repo:        repo,
		redisClient: redisClient,
	}
}

func (s *FeedService) Get(userId string) ([]entity.Post, error) {
	ctx := context.Background()
	key := fmt.Sprintf("user:%s:feed", userId)
	postsIds, err := s.redisClient.LRange(ctx, key, 0, 999).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to get cached feed: %w", err)
	}

	if len(postsIds) == 0 {
		return []entity.Post{}, nil
	}

	feed, err := s.repo.GetByIds(postsIds)

	if err != nil {
		return nil, fmt.Errorf("failed to get feed: %w", err)
	}

	return feed, nil
}

func (s *FeedService) AddPost(userId string, postId string) error {
	ctx := context.Background()
	key := fmt.Sprintf("user:%s:feed", userId)
	_, err := s.redisClient.LPush(ctx, key, postId).Result()

	if err != nil {
		return fmt.Errorf("failed to add post to cached feed: %w", err)
	}

	_, err = s.redisClient.LTrim(ctx, key, 0, 999).Result()

	if err != nil {
		return fmt.Errorf("failed to trim cached feed: %w", err)
	}

	return nil
}
