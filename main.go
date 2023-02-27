//go:generate tools/scripts/mockgen.sh
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"go-lambda/conf"
	"go-lambda/container"
	"log"
)

type Event struct {
	Name string `json:"name"`
}

func HandleLambdaEvent(event Event) {
	if event.Name == "Monthly" {
		cfg := conf.NewConfig()
		dic := container.New(cfg)

		eventHandler := dic.NewEventHandler(dic.NewSlackClient())
		err := eventHandler.RemindForReports()
		if err != nil {
			log.Println("RemindForReports has return an error", err)
		}
	}
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
