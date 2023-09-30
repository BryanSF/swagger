package repository

import (
	"context"
	"io"
)

type GoogleRepository interface {
	GetObjectURL(string, string) (*string, error)
	UploadObject(context.Context, io.Reader, string) (string, error)
}
