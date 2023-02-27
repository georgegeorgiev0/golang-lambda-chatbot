package container

import (
	"go-lambda/client"
	"go-lambda/conf"
	"go-lambda/handler"
	"net/http"
)

type Container struct {
	cfg          *conf.Config
	SlackClient  client.SlackClient
	EventHandler handler.EventHandler
}

func New(cfg *conf.Config) *Container {
	return &Container{
		cfg: cfg,
	}
}

func (container Container) NewSlackClient() client.SlackClient {
	container.SlackClient = client.SlackClient{
		Config:     container.cfg,
		HttpClient: &http.Client{},
	}

	return container.SlackClient
}

func (container Container) NewEventHandler(slackClient client.SlackClient) *handler.EventHandler {
	container.EventHandler = handler.EventHandler{
		Config:      container.cfg,
		SlackClient: slackClient,
	}

	return &container.EventHandler
}
