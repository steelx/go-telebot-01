package main

import (
	"log"
	"os"

	"github.com/steelx/go-telebot-01/bothandlers"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicUrl = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicUrl},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	bot, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	//handlers
	bothandlers.InitButtons(bot)
	bothandlers.MapRoutes(bot)

	bot.Start()

}
