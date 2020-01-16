package services

import (
	"deku/models"
	"deku/repositories"
)

type PostService interface {
	GetAll() []models.Post
	GetByID(id int64) (models.Post, bool)
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postTmpService{
		repo: repo,
	}
}

type postTmpService struct {
	repo repositories.PostRepository
}

func (s *postTmpService) GetAll() []models.Post {
	return s.repo.SelectMany(func(_ models.Post) bool {
		return true
	}, -1)
}

func (s *postTmpService) GetByID(id int64) (models.Post, bool) {
	return s.repo.Select(func(m models.Post) bool {
		return m.ID == id
	})
}
