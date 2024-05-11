package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eydeveloper/highload-social/internal/entity"
	"github.com/eydeveloper/highload-social/internal/repository"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type FeedService struct {
	repo        repository.Post
	redisClient *redis.Client
	amqpChannel *amqp.Channel
}

func NewFeedService(repo repository.Post, redisClient *redis.Client, amqpChannel *amqp.Channel) *FeedService {
	return &FeedService{
		repo:        repo,
		redisClient: redisClient,
		amqpChannel: amqpChannel,
	}
}

func (s *FeedService) AddPost(userId string, post entity.Post) error {
	ctx := context.Background()
	key := fmt.Sprintf("user:%s:feed", userId)
	_, err := s.redisClient.LPush(ctx, key, post.Id).Result()

	if err != nil {
		return fmt.Errorf("failed to add post to cached feed: %w", err)
	}

	_, err = s.redisClient.LTrim(ctx, key, 0, 999).Result()

	if err != nil {
		return fmt.Errorf("failed to trim cached feed: %w", err)
	}

	jsonPost, err := json.Marshal(post)

	if err != nil {
		return fmt.Errorf("failed to parse the post to json format: %w", err)
	}

	err = s.amqpChannel.PublishWithContext(
		context.Background(),
		"feed",
		userId,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonPost,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish to the posts exchange: %w", err)
	}

	return nil
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

func (s *FeedService) GetRealTime(userId string) (<-chan amqp.Delivery, error) {
	amqpQueue, err := s.amqpChannel.QueueDeclare(
		"feed",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	err = s.amqpChannel.QueueBind(
		amqpQueue.Name,
		userId,
		"feed",
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	deliveries, err := s.amqpChannel.Consume(
		amqpQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return deliveries, nil
}
