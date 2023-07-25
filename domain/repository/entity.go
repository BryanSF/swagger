package repository

type GoogleRepository interface {
	GetObjectURL(string, string) (string, error)
}
