package service

import (
	"github.com/BryanSF/swagger/domain/repository"
)

type GoogleService struct {
	repo repository.GoogleRepository
} 

func NewGoogleService(r repository.GoogleRepository) *GoogleService {
	return &GoogleService{
		repo: r,
	}
}

func (s *GoogleService) GetObjectURL(bucketName, objectName string) (string, error) {
	b, err := s.repo.GetObjectURL(bucketName, objectName)
	if err != nil {
		return "", err
	} else {
		return b, nil
	}
}
