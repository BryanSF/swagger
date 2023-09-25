package controller

type Config struct {
	Bucket string `env:"BUCKET" required:"true"`
}
