package handler

import (
	"fmt"
	"log"
)

const postMessage = "@here please login your time .... you must have 168 hours this month"

func (h EventHandler) RemindForReports() error {
	url := fmt.Sprintf(h.Config.SlackUrl, "chat.postMessage")

	data, err := h.SlackClient.PreparePostMessagePayload(postMessage, h.Config.RemindChannel)
	if err != nil {
		log.Println("Error when PreparePostMessagePayload was executed", err)
		return err
	}

	bearer := fmt.Sprintf("Bearer %s", h.Config.SlackToken)
	if err != nil {
		log.Println("Error when sprintf fill bearer token", err)
		return err
	}

	err = h.SlackClient.PostMessage(url, bearer, data)
	if err != nil {
		log.Println("Error when PostMessage was executed", err)
		return err
	}

	return nil
}
