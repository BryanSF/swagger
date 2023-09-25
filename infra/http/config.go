package http

type Config struct {
	Port                  int `env:"HTTP_PORT" required:"true"`
	DisableStartupMessage bool
}
