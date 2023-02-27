package handler

import (
	"go-lambda/client"
	"go-lambda/conf"
)

type EventHandler struct {
	Config      *conf.Config
	SlackClient client.SlackClientInt
}
