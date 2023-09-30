package service

import (
	"context"
	"io"

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

func (s *GoogleService) GetObjectURL(bucketName, objectName string) (*string, error) {
	b, err := s.repo.GetObjectURL(bucketName, objectName)
	if err != nil {
		return nil, err
	} else {
		return b, nil
	}
}

func (s *GoogleService) UploadObject(ctx context.Context, file io.Reader, filename string) (string, error) {
	u, err := s.repo.UploadObject(ctx, file, filename)
	if err != nil {
		return "", err
	} else {
		return u, nil
	}
}
