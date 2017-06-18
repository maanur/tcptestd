package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func init() {
	receivers = append(receivers, slackbot)
}

func slackbot() {
	logger := log.New(os.Stdout, "[slackbot] ", log.Lshortfile|log.LstdFlags)
	token := os.Getenv("ESCBOT_SLACK_TOKEN")
	if token == "" {
		log.Fatal("$ESCBOT_SLACK_TOKEN must be set")
	}
	api := slack.New(token)
	slack.SetLogger(logger)
	api.SetDebug(true)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for msg := range rtm.IncomingEvents {
		logger.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			logger.Println("Infos:", ev.Info)
			logger.Println("Connection counter:", ev.ConnectionCount)
			// Replace #general with your Channel ID
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "#general"))

		case *slack.MessageEvent:
			logger.Printf("Message: %v\n", ev)

		case *slack.PresenceChangeEvent:
			logger.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			logger.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			logger.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			logger.Printf("Invalid credentials")
			return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
