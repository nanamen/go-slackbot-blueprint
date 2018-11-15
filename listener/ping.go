package listener

import "github.com/nlopes/slack"

func (s *SlackListener) ping() (attachment slack.Attachment) {
	attachment = slack.Attachment{
		Pretext: "pong",
		Color:   "#B733FF",
	}
	return attachment
}
