package main

import (
	"io"
	"log"
	"os"

	"strings"

	"github.com/nlopes/slack"
)

func init() {
	receivers = append(receivers, slackbot)
}

func slackbot(output io.Writer) {
	logger := log.New(output, "[slackbot] ", log.Lshortfile|log.LstdFlags)
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
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "@grushin_m"))

		case *slack.MessageEvent:
			logger.Printf("Message: %v\n", ev)
			if strings.Contains(ev.Msg.Text, "testtcpmail") {
				pmp := slack.PostMessageParameters{}
				_, _, err := api.PostMessage(ev.Channel, "Let's start testtcpmail", pmp)
				if err != nil {
					logger.Println(err)
				}
			}

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
