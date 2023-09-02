package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	gogpt "github.com/tapp-ai/go-openai"
	"go.uber.org/zap"
)

type App struct {
	Log        *zap.Logger
	GptClient  *gogpt.Client
	HttpClient *http.Client
}

func (a *App) ServeRoot(c *fiber.Ctx) error {
	indexHtml := `pong from tapp coreapi`
	return c.Status(200).Send([]byte(indexHtml))
}
