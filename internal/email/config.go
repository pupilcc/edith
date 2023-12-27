package email

import (
	"edith/internal/util"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

func GetConfig() *Config {
	port, _ := strconv.Atoi(util.GetEnv("EMAIL_PORT"))
	return &Config{
		util.GetEnv("EMAIL_HOST"),
		port,
		util.GetEnv("EMAIL_USERNAME"),
		util.GetEnv("EMAIL_PASSWORD"),
	}
}
