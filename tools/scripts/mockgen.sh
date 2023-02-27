#!/bin/sh

go run github.com/golang/mock/mockgen -source=./client/slack_client.go -package=client -destination=./client/slack_client_mocks.go
