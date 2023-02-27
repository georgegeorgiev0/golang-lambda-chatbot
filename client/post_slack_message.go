package client

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// PostMessage send a message in Slack channel
func (slackClient SlackClient) PostMessage(url, bearer string, data []byte) error {
	// Create a new request using http
	req, err := slackClient.CreateRequest(url, bearer, data)
	if err != nil {
		log.Println("Error while creating new request.\n[ERROR] -", err)
		return err
	}

	// Send req using http Client
	resp, err := slackClient.HttpClient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return err
	}

	//Body should be closed
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return err
	}

	log.Println(string([]byte(body)))

	return nil
}

// CreateRequest create a request
func (slackClient SlackClient) CreateRequest(url, bearer string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Error while creating new request.\n[ERROR] -", err)
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-type", "application/json")

	return req, nil
}
