package email

import (
	"edith/internal/middleware"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

func Send(to string, subject string, body string) {
	config := GetConfig()
	m := gomail.NewMessage()
	m.SetHeader("From", config.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)

	err := d.DialAndSend(m)
	logger := middleware.GetLogger()
	if err != nil {
		logger.Error("Send Email Error", zap.Error(err))
	}
}
