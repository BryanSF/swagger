package module

import (
	"github.com/BryanSF/swagger/domain/service"
	"go.uber.org/fx"
)

var Service = fx.Module("services",
	fx.Provide(service.NewGoogleService),
)