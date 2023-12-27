package things

import (
	"edith/internal/email"
	"edith/internal/util"
	"strings"
)

func AddTask(message string) {
	task := strings.Split(message, "/")
	to := util.GetEnv("THINGS_EMAIL")
	if len(task) != 2 {
		subject := task[0]
		email.Send(to, subject, subject)
	} else {
		subject := task[0]
		body := task[1]
		email.Send(to, subject, body)
	}
}
