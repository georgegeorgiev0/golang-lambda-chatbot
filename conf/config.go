package conf

import (
	"os"
)

type args struct {
	SlackUrl   string
	SlackToken string
}

type Config struct {
	SlackUrl      string
	SlackToken    string
	RemindChannel string
}

func NewConfig() *Config {
	return &Config{
		SlackUrl:      os.Getenv("SLACK_URL") + "%s",
		SlackToken:    os.Getenv("SLACK_TOKEN"),
		RemindChannel: os.Getenv("REMIND_CHANNEL"),
	}
}
