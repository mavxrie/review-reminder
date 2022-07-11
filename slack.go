package main

import (
	"log"

	"github.com/slack-go/slack"
)

// SlackWrite sends a message to slack.
func SlackWrite(slack_token, target, msg string) error {
	api := slack.New(slack_token)

	_, _, _, err := api.SendMessage(
		target,
		slack.MsgOptionText(msg, false),
	)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
