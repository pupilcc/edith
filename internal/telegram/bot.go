package telegram

import (
	"edith/internal/middleware"
	"edith/internal/util"
	"go.uber.org/zap"
	telebot "gopkg.in/telebot.v3"
	"strconv"
)

var bot *telebot.Bot

type BotEnv struct {
	Token  string
	Secret string
	Domain string
	UserId int64
}

func GetBotEnv() *BotEnv {
	num, _ := strconv.ParseInt(util.GetEnv("BOT_USER_ID"), 10, 64)
	return &BotEnv{
		util.GetEnv("BOT_TOKEN"),
		util.GetEnv("BOT_SECRET"),
		util.GetEnv("BOT_DOMAIN"),
		num,
	}
}

func GetBot() *telebot.Bot {
	if bot != nil {
		return bot
	}

	var err error
	bot, err = telebot.NewBot(telebot.Settings{
		Token: GetBotEnv().Token,
	})

	logger := middleware.GetLogger()
	if err != nil {
		logger.Error("Failed to create bot", zap.Error(err))
		return nil
	}

	return bot
}

func GetWebhook() *telebot.Webhook {
	endpoint := &telebot.WebhookEndpoint{
		PublicURL: GetBotEnv().Domain + "/webhook/bot",
	}
	webhook := &telebot.Webhook{
		SecretToken: GetBotEnv().Secret,
		Endpoint:    endpoint,
	}
	return webhook
}
