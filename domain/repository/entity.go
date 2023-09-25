package repository

type GoogleRepository interface {
	GetObjectURL(string, string) (*string, error)
	UploadObject(string, string, string) error
}
