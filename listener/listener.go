package listener

import (
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
)

type SlackListener struct {
	client    *slack.Client
	botID     string
	botName   string
	channelID string
}

func NewSlackListener(client *slack.Client, botID string, channelID string) *SlackListener {
	return &SlackListener{
		client:    client,
		botID:     botID,
		channelID: channelID,
	}
}

func (s *SlackListener) ListenAndResponse() int {
	rtm := s.client.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				s.botName = ev.Info.User.Name

			case *slack.MessageEvent:
				if err := s.handleMessageEvent(ev); err != nil {
					log.Printf("[ERROR] Failed to handle message: %s", err)
				}

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1
			}
		}
	}
}

func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) error {
	if ev.Channel != s.channelID {
		log.Printf("%s %s", ev.Channel, ev.Msg.Text)
		return nil
	}

	if !strings.HasPrefix(ev.Msg.Text, fmt.Sprintf("<@%s> ", s.botID)) {
		return nil
	}

	m := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")[1:]
	if len(m) == 0 {
		return fmt.Errorf("invalid message")
	}

	var attachment slack.Attachment
	switch m[0] {
	case "help":
		attachment = s.help()
	case "ping":
		attachment = s.ping()
	default:
		return fmt.Errorf("invalid message")
	}

	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			attachment,
		},
	}

	if _, _, err := s.client.PostMessage(ev.Channel, "", params); err != nil {
		return fmt.Errorf("failed to post message: %s", err)
	}

	return nil
}
