package main

import (
	"edith/internal/api"
	"edith/internal/middleware"
	"edith/internal/telegram"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	e := echo.New()

	// Route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "UP",
		})
	})

	// Routes
	api.WebhookRoutes(e)

	// Logger
	e.Use(middleware.RequestLogger())
	logger := middleware.GetLogger()

	// Telegram Bot
	b := telegram.GetBot()
	err := b.SetWebhook(telegram.GetWebhook())
	if err != nil {
		logger.Error("Error setting webhook: ", zap.Error(err))
	}

	// Start the service
	e.Logger.Fatal(e.Start(":1455"))
}
