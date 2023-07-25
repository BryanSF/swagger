package http

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var FiberModule = fx.Module(
	"fiber",
	fx.Provide(NewFiber),
	fx.Invoke(HookFiber),
)

func NewFiber() *fiber.App {
	config := Config{
		DisableStartupMessage: true,
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: config.DisableStartupMessage,
	})
	app.Use(cors.New())

	return app
}

func HookFiber(lc fx.Lifecycle, app *fiber.App, l *zap.SugaredLogger, c Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := app.Listen(fmt.Sprintf(":%d", c.Port)); err != nil {
					l.Error(err)
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			return app.Shutdown()
		},
	})
}
