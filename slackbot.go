package main

import (
	"io"
	"log"
	"os"

	"strings"

	"github.com/nlopes/slack"
)

func init() {
	//receivers = append(receivers, slackbot)
}

func slackbot(output io.Writer) {
	logger := log.New(output, "[slackbot] ", log.Lshortfile|log.LstdFlags)
	token := os.Getenv("ESCBOT_SLACK_TOKEN")
	if token == "" {
		log.Println("$ESCBOT_SLACK_TOKEN must be set")
		return
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
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "#random"))

		case *slack.MessageEvent:
			logger.Printf("Message: %v\n", ev)
			for actName, action := range actions {
				if strings.Contains(ev.Msg.Text, actName) {
					act := action()
					func(act Action) {
						actGPS, actGPE := act.GetParam(ev.Msg.Text)
						if actGPE != nil {
							rtm.SendMessage(rtm.NewOutgoingMessage(actGPE.Error(), ev.Msg.Channel))
							return
						}
						pmp := slack.PostMessageParameters{}
						var att slack.Attachment
						att.Text = "Запустить " + act.Name() + " с параметром " + actGPS
						_, _, err := api.PostMessage(ev.Msg.Channel, "Start testtcpmail?", pmp)
						if err != nil {
							logger.Println(err)
						}
					}(act)
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
