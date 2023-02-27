package client

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"go-lambda/conf"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSlackClient_PostMessageWithHttpClientError(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientHttp := NewMockHttpClient(ctrl)
	config := conf.NewConfig()
	config.RemindChannel = "channel"
	config.SlackToken = "token"
	config.SlackUrl = "https://example.com/"

	clientHttp.EXPECT().Do(gomock.Any()).Return(nil, errors.New("oops"))
	slackClient := SlackClient{Config: config, HttpClient: clientHttp}
	err := slackClient.PostMessage("https://example.com/", "token", []byte{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSlackClient_PostMessageWithRequestError(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientHttp := NewMockHttpClient(ctrl)
	config := conf.NewConfig()
	config.RemindChannel = "channel"
	config.SlackToken = "token"
	config.SlackUrl = "wrong%shttp"

	clientHttp.EXPECT().Do(gomock.Any()).Times(0)
	slackClient := SlackClient{Config: config, HttpClient: clientHttp}
	err := slackClient.PostMessage("wrong%shttp", "token", []byte{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSlackClient_PostMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientHttp := NewMockHttpClient(ctrl)
	config := conf.NewConfig()
	config.RemindChannel = "channel"
	config.SlackToken = "token"
	config.SlackUrl = "https://example.com/"

	responseBody := ioutil.NopCloser(bytes.NewReader([]byte(`{"name": "darvin"}`)))
	response := &http.Response{
		StatusCode: 200,
		Body:       responseBody,
	}

	clientHttp.EXPECT().Do(gomock.Any()).Return(response, nil)

	slackClient := SlackClient{Config: config, HttpClient: clientHttp}
	err := slackClient.PostMessage("https://example.com/", "token", []byte{})
	if err != nil {
		t.Fatal("Unexpected error ", err)
	}
}
