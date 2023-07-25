package config

import (
	"github.com/BryanSF/swagger/infra/config"
	"github.com/BryanSF/swagger/infra/http"
	env "github.com/Netflix/go-env"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("config",
	fx.Provide(NewConfig),
	fx.Provide(func(cfg Config) config.Config { return cfg.Gcp }),
	fx.Provide(func(cfg Config) http.Config { return cfg.Http }),
)

type Config struct {
	Gcp    config.Config
	Http   http.Config
	Extras *env.EnvSet
}

func NewConfig(logger *zap.SugaredLogger) Config {
	var cfg Config = Config{}
	err := cfg.loadConfig()
	if err != nil {
		logger.Error(err)
	}

	return cfg
}

func (c *Config) loadConfig() error {
	es, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		return err
	}
	c.Extras = &es

	return nil
}
