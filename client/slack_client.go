package client

import (
	"go-lambda/conf"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SlackClientInt interface {
	PostMessage(url, bearer string, data []byte) error
	CreateRequest(url, bearer string, data []byte) (*http.Request, error)
	PreparePostMessagePayload(message, channel string) ([]byte, error)
}

type SlackClient struct {
	Config     *conf.Config
	HttpClient HttpClient
}
