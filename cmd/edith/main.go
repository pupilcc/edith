package main

import (
	"edith/internal/api"
	"edith/internal/middleware"
	"edith/internal/telegram"
	"github.com/labstack/echo/v4"
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

	// Telegram Bot
	b := telegram.GetBot()
	_ = b.SetWebhook(telegram.GetWebhook())

	// Routes
	api.WebhookRoutes(e)

	// Logger
	e.Use(middleware.RequestLogger())

	// Start the service
	e.Logger.Fatal(e.Start(":1455"))
}
