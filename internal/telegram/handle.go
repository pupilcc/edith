package telegram

import (
	"edith/internal/middleware"
	"edith/internal/things"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
	"strings"
)

func verify(c echo.Context) (*telebot.Update, error) {
	secretToken := c.Request().Header.Get("X-Telegram-Bot-Api-Secret-Token")
	configSecretToken := GetBotEnv().Secret

	var update telebot.Update
	var err error
	if err = c.Bind(&update); err != nil {
		return nil, err
	}

	if secretToken != configSecretToken {
		err = telebot.ErrBadUserID
		return nil, err
	}

	chatId := update.Message.Chat.ID
	if chatId != GetBotEnv().UserId {
		err = telebot.ErrBadUserID
		return nil, err
	}

	return &update, nil
}

func HandleUpdate(c echo.Context) error {
	logger := middleware.GetLogger()
	update, err := verify(c)
	if err != nil {
		logger.Error("verify error", zap.Error(err))
		return err
	}

	b := GetBot()
	b.ProcessUpdate(*update)
	text := update.Message.Text

	isCommand := update.Message.Entities != nil
	if isCommand {
		parts := strings.Split(text, " ")
		command := parts[0]
		message := parts[1]
		if command == "/task" {
			logger.Info("/task", zap.String("message", message))
			things.AddTask(message)
		}
	}

	return nil
}
