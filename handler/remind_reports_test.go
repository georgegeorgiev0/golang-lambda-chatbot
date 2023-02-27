package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"go-lambda/client"
	"go-lambda/conf"
	"testing"
)

func TestEventHandler_RemindForReports(t *testing.T) {
	ctrl := gomock.NewController(t)
	slackClientMock := client.NewMockSlackClientInt(ctrl)
	data := map[string]any{
		"text":       "example",
		"channel":    "channel",
		"link_names": 1,
	}
	config := conf.NewConfig()
	config.RemindChannel = "channel"
	config.SlackToken = "token"
	config.SlackUrl = "https://example.com/%s"

	payload, _ := json.Marshal(data)
	slackClientMock.EXPECT().PostMessage(fmt.Sprintf(config.SlackUrl, "chat.postMessage"), "Bearer token", payload)
	slackClientMock.EXPECT().PreparePostMessagePayload("@here please login your time .... you must have 168 hours this month", config.RemindChannel).Return(payload, nil)

	handler := EventHandler{
		SlackClient: slackClientMock,
		Config:      config,
	}

	err := handler.RemindForReports()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
}

func TestEventHandler_RemindForReportsWithErrorFromPostMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	slackClientMock := client.NewMockSlackClientInt(ctrl)
	config := conf.NewConfig()
	config.RemindChannel = "channel"
	config.SlackToken = "token"
	config.SlackUrl = "https://example.com/%s"
	data := map[string]any{
		"text":       "example",
		"channel":    "channel",
		"link_names": 1,
	}
	payload, _ := json.Marshal(data)
	slackClientMock.EXPECT().PostMessage(fmt.Sprintf(config.SlackUrl, "chat.postMessage"), "Bearer token", payload).Return(errors.New("ooops"))
	slackClientMock.EXPECT().PreparePostMessagePayload("@here please login your time .... you must have 168 hours this month", config.RemindChannel).Return(payload, nil)

	handler := EventHandler{
		SlackClient: slackClientMock,
		Config:      config,
	}

	err := handler.RemindForReports()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEventHandler_RemindForReportsWithErrorFromPreparePostMessagePayload(t *testing.T) {
	ctrl := gomock.NewController(t)
	slackClientMock := client.NewMockSlackClientInt(ctrl)
	config := conf.NewConfig()
	config.RemindChannel = "channel"
	config.SlackToken = "token"
	config.SlackUrl = "https://example.com/%s"

	slackClientMock.EXPECT().PreparePostMessagePayload("@here please login your time .... you must have 168 hours this month", config.RemindChannel).Return(nil, errors.New("ooopps"))

	handler := EventHandler{
		SlackClient: slackClientMock,
		Config:      config,
	}

	err := handler.RemindForReports()
	if err == nil {
		t.Fatal("expected error")
	}
}
