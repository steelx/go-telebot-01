package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicUrl = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
	)

	//fmt.Printf("port %v url %v token %v", port, publicUrl, token)

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

	bot.Handle("/hello", func(m *tb.Message) {
		_, _ = bot.Send(m.Sender, "Hi !")
	})

	bot.Start()

}
