package clound

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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
	}
}

func (c *Clound) GetObjectURL(bucketName, objectName string) (string, error) {
	ctx := context.Background()

	// Obtenha o objeto do bucket
	object := c.Client.Bucket(bucketName).Object(objectName)

	// Obtenha os atributos do objeto
	attrs, err := object.Attrs(ctx)
	if err != nil {
		return "", err
	}

	// Retorna a URL pública do objeto
	return attrs.MediaLink, nil
}

/*

func main() {
	ctx := context.Background()
	bucketName := "imagem_2"
	c := new(Clound)
	// Defina o nome do objeto (imagem)
	objectName := "aleatorio.jpeg"

	// Obtenha o objeto do bucket
	object := c.Client.Bucket(bucketName).Object(objectName)

	// Obtenha a URL pública do objeto
	attrs, err := object.Attrs(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Crie um aplicativo Fiber
	app := fiber.New()

	// Rota para exibir a imagem
	app.Get("/imagem", func(c *fiber.Ctx) error {
		// Crie uma requisição HTTP para obter o conteúdo da imagem
		resp, err := http.Get(attrs.MediaLink)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Copie o conteúdo da resposta para a resposta do Fiber
		c.Set("Content-Type", attrs.ContentType)
		_, err = io.Copy(c, resp.Body)
		return err
	})

	// Execute o servidor Fiber
	log.Fatal(app.Listen(":3000"))
}
*/

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
