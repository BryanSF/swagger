package clound

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"github.com/BryanSF/swagger/domain/repository"
	"github.com/BryanSF/swagger/infra/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type Clound struct {
	Client *storage.Client
	Ctx    context.Context
	Cfg    config.Config
}

var Module = fx.Module("cloud",
	fx.Provide(NewClient),
	fx.Invoke(HookGoogleCloud),
	fx.Provide(func(r *Clound) repository.GoogleRepository { return r }),
)

func NewClient(c config.Config) *Clound {
	ctx := context.Background()

	gcp, mErr := json.Marshal(c)
	if mErr != nil {
		fmt.Println(mErr)
	}

	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(gcp))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	return &Clound{
		Client: client,
		Ctx:    ctx,
		Cfg:    c,
	}
}

func (c *Clound) GetObjectURL(bucketName, objectName string) (*string, error) {
	// Obtenha o objeto do bucket
	url, err := c.Client.Bucket(bucketName).SignedURL(objectName, &storage.SignedURLOptions{
		GoogleAccessID: c.Cfg.ClientEmail,
		PrivateKey:     []byte(c.Cfg.PrivateKey),
		Method:         "GET",
		Expires:        time.Now().Add(24 * time.Hour),
		Headers:        nil,
		// QueryParams:    nil,
	})

	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (c *Clound) UploadObject(ctx context.Context, file io.Reader, filename string) (string, error) {
	sw := c.Client.Bucket("").Object(filename).NewWriter(ctx)

	if _, err := io.Copy(sw, file); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}

	return sw.Attrs().Name, nil
}

func HookGoogleCloud(lc fx.Lifecycle, c *Clound, logger *zap.SugaredLogger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Google clound connection has been established successfully!")
			return nil
		},
		OnStop: func(context.Context) error {
			c.Client.Close()
			logger.Info("Google Clound connection Closed!")
			return nil
		},
	})
}
