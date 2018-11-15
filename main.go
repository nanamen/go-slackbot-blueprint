package main

import (
	"github.com/nanamen/go-slackbot-blueprint/listener"
	"github.com/nanamen/go-slackbot-blueprint/schedule"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/nlopes/slack"
)

type envConfig struct {
	WebhookURL string `envconfig:"WEBHOOK_URL" required:"true"`
	BotToken   string `envconfig:"BOT_TOKEN" required:"true"`
	BotID      string `envconfig:"BOT_ID" required:"true"`
	ChannelID  string `envconfig:"CHANNEL_ID" required:"true"`
}

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		return 1
	}

	log.Printf("[INFO] Strat Scheduled events")
	s := schedule.NewSchedule(env.WebhookURL, env.BotID, env.ChannelID)
	s.Start()

	// Listening slack event and response
	log.Printf("[INFO] Start slack event listening")
	client := slack.New(env.BotToken)
	slackListener := listener.NewSlackListener(client, env.BotID, env.ChannelID)
	return slackListener.ListenAndResponse()
}
