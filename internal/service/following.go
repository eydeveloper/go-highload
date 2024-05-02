package service

import "github.com/eydeveloper/highload-social/internal/repository"

type FollowingService struct {
	repo repository.Following
}

func NewFollowingService(repo repository.Following) *FollowingService {
	return &FollowingService{repo: repo}
}

func (s *FollowingService) Follow(followeeId string, followerId string) error {
	return s.repo.Follow(followeeId, followerId)
}

func (s *FollowingService) Unfollow(followeeId string, followerId string) error {
	return s.repo.Unfollow(followeeId, followerId)
}

func (s *FollowingService) GetFollowers(followeeId string) ([]string, error) {
	return s.repo.GetFollowers(followeeId)
}
