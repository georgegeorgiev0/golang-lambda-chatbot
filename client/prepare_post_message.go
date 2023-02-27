package client

import "encoding/json"

func (slackClient SlackClient) PreparePostMessagePayload(message, channel string) ([]byte, error) {
	data := map[string]any{
		"text":       message,
		"channel":    channel,
		"link_names": 1,
	}

	return json.Marshal(data)
}
