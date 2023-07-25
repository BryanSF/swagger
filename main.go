package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/BryanSF/swagger/domain/module"
	config "github.com/BryanSF/swagger/infra/dotenv"
	clound "github.com/BryanSF/swagger/infra/gcp"
	"github.com/BryanSF/swagger/infra/http"
	"github.com/BryanSF/swagger/infra/logger"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

// @title Test Image Endpoint's
// @version 0.0.1
// @description Return Image Bucket.
// @termsOfService None
// @contact.name None
// @contact.email suport@none.me
// @license.name Idp: v0.0.1
// @license.url none.me
// @host https://none.run.app
// @BasePath /
func main() {
	if os.Getenv("ENV") != "production" {
		LoadConfig()
	}

	fx.New(
		config.Module,
		clound.Module,
		logger.Module,
		http.Module,
		module.Service,
	).Run()
}

func LoadConfig() {
	_, b, _, _ := runtime.Caller(0)

	basepath := filepath.Dir(b)

	err := godotenv.Load(fmt.Sprintf("%v/.env", basepath))
	if err != nil {
		panic(err)
	}
}
