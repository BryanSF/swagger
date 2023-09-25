package http

import (
	"github.com/BryanSF/swagger/infra/http/controller"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("http",
	FiberModule,
	fx.Provide(controller.NewCloundController),
	fx.Invoke(RegisterControllers),
)

func RegisterControllers(app *fiber.App, cloundController *controller.CloundController) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/swagger/*", swagger.HandlerDefault) // default

	v1.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "http://example.com/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	cloundController.RegisterRoutes(v1)
}
