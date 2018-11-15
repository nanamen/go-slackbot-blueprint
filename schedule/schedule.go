package schedule

import (
	"github.com/robfig/cron"
)

type Schedule struct {
	cron       *cron.Cron
	webhookUrl string
	botID      string
	channelID  string
}

func NewSchedule(webhookUrl string, botID string, channelID string) *Schedule {
	s := &Schedule{
		botID:      botID,
		channelID:  channelID,
		webhookUrl: webhookUrl,
	}

	c := cron.New()
	c.AddFunc("0 * * * * *", s.doSomething)

	s.cron = c
	return s
}

func (s *Schedule) Start() {
	s.cron.Start()
}

func (s *Schedule) Stop() {
	s.cron.Stop()
}
