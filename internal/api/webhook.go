package api

import (
	"edith/internal/telegram"
	"github.com/labstack/echo/v4"
)

func WebhookRoutes(e *echo.Echo) {
	webhook := e.Group("/webhook")
	webhook.POST("/bot", telegram.HandleUpdate)
}
