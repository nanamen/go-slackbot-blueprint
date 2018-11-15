package listener

import "github.com/nlopes/slack"

var (
	commands = map[string]string{
		"help": "Displays all of the help commands.",
		"ping": "Reply with pong.",
	}
)

func (s *SlackListener) help() (attachment slack.Attachment) {
	fields := make([]slack.AttachmentField, 0)

	for k, v := range commands {
		fields = append(fields, slack.AttachmentField{
			Title: "@" + s.botName + " " + k,
			Value: v,
		})
	}

	attachment = slack.Attachment{
		Pretext: s.botName + "Command List",
		Color:   "#B733FF",
		Fields:  fields,
	}
	return attachment
}
