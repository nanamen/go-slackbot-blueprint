package schedule

import (
	"github.com/nlopes/slack"
	"time"
)

func (s *Schedule) doSomething() {
	// 土日はやらない
	doNotWork := map[time.Weekday]int{
		time.Saturday: 0,
		time.Sunday:   0,
	}
	if _, ok := doNotWork[time.Now().Weekday()]; ok {
		return
	}

	// do something
	msg := &slack.WebhookMessage{}
	field := slack.AttachmentField{
		Title: "foo",
		Value: "bar",
	}

	attachment := slack.Attachment{
		Pretext: "reslut",
		Fields:  []slack.AttachmentField{field},
		Color:   "#B733FF",
	}

	msg.Attachments = []slack.Attachment{
		attachment,
	}

	slack.PostWebhook(s.webhookUrl, msg)
}
