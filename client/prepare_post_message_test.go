package client

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"go-lambda/conf"
	"testing"
)

func TestSlackClient_PreparePostMessagePayload(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientHttp := NewMockHttpClient(ctrl)
	config := conf.NewConfig()
	data := map[string]any{
		"text":       "example",
		"channel":    "channel",
		"link_names": 1,
	}
	payload, _ := json.Marshal(data)

	slackClient := SlackClient{Config: config, HttpClient: clientHttp}
	preparedPayload, _ := slackClient.PreparePostMessagePayload("test message", "XHKSFD")
	res := bytes.Compare(preparedPayload, payload)
	if res == 0 {
		t.Fatal("expected prepared payload to be equal")
	}
}
