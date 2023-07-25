package controller

import (
	"fmt"
	"io"
	"net/http"

	"github.com/BryanSF/swagger/domain/service"
	"github.com/BryanSF/swagger/infra/http/dto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CloundController struct {
	logger   *zap.SugaredLogger
	CService *service.GoogleService
}

func NewCloundController(logger *zap.SugaredLogger, c *service.GoogleService) *CloundController {
	return &CloundController{logger: logger, CService: c}
}

func (c CloundController) RegisterRoutes(app fiber.Router) {
	clound := app.Group("/imgs")
	clound.Get("/get", c.GetObjectURL)
}

// GetObjectURL retrieves the URL of an object from the Google Cloud Storage.
//
// @Summary Get object URL from Google Cloud Storage
// @Description Retrieves the URL of an object from the Google Cloud Storage.
// @Tags Cloud Storage
// @Accept json
// @Produce json
// @Param bucket body string true "Bucket"
// @Param file body string true "File"
// @Success 200 {string} dto.Base "Success"
// @Failure 400 {object} dto.Base "Bad request"
// @Failure 500 {object} dto.Base "Internal Server Error"
// @Router /imgs/get [get]
func (c *CloundController) GetObjectURL(ctx *fiber.Ctx) error {
	response := dto.Base{
		Success: false,
		Message: "",
		Error:   "",
	}

	type payload struct {
		Bucket string `json:"bucket"`
		File   string `json:"file"`
	}

	var request payload

	if err := ctx.BodyParser(&request); err != nil {
		response.Message = "Não foi possível completar essa operação"
		response.Error = "Bad Request"
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	url, err := c.CService.GetObjectURL(request.Bucket, request.File)

	if err != nil {
		response.Message = "Não foi possível completar essa operação"
		response.Error = "Aconteceu alguma coisa"
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	// Recupere os dados do objeto no Google Cloud Storage
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		response.Message = "Não foi possível completar essa operação"
		response.Error = "Aconteceu alguma coisa"
		fmt.Println(resp.StatusCode)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	// Define o cabeçalho Content-Type para a imagem
	ctx.Set("Content-Type", resp.Header.Get("Content-Type"))

	// Retorna a imagem como resposta
	_, err = io.Copy(ctx, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
