//go:build integration
// +build integration

package integration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"go-lambda/client"
	"go-lambda/conf"
	goHandler "go-lambda/handler"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPostRemindSlackMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	config := conf.NewConfig()
	config.RemindChannel = "channel"
	config.SlackToken = "token"
	config.SlackUrl = "https://example.com/%s"
	data := map[string]any{
		"text":       "@here please login your time .... you must have 168 hours this month",
		"channel":    "channel",
		"link_names": 1,
	}

	url := fmt.Sprintf(config.SlackUrl, "chat.postMessage")
	bearer := fmt.Sprintf("Bearer %s", config.SlackToken)
	mdata, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(mdata))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-type", "application/json")

	responseBody := ioutil.NopCloser(bytes.NewReader([]byte(`{"name": "darvin"}`)))
	response := &http.Response{
		StatusCode: 200,
		Body:       responseBody,
	}

	httpClientMock := client.NewMockHttpClient(ctrl)
	httpClientMock.EXPECT().Do(gomock.Any()).DoAndReturn(
		func(request *http.Request) (compare bool, err error) {
			if request.URL != req.URL ||
				request.Header.Get("Authorization") != request.Header.Get("Authorization") {
				return false, errors.New("Request objects are not equal")
			}

			return true, nil
		}).Return(response, nil)

	handler := goHandler.EventHandler{
		Config: config,
		SlackClient: client.SlackClient{
			Config:     config,
			HttpClient: httpClientMock,
		},
	}

	handler.RemindForReports()
}
