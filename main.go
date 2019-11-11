package main

import (
	"fmt"
	"github.com/steelx/go-telebot-01/bothandlers"
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

	//buttons
	btnMoon := tb.InlineButton{
		Unique: "Moon",
		Text: "Moon 🌹",
	}
	btnSun := tb.InlineButton{
		Unique:          "Sun",
		Text:            "Sun 🌞",
	}

	bot.Handle(&btnMoon, func(c *tb.Callback) {
		_ = bot.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})

		_, _ = bot.Send(c.Sender, "Moon says hi!")
	})
	bot.Handle(&btnSun, func(c *tb.Callback) {
		_ = bot.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})

		_, _ = bot.Send(c.Sender, "Sun says 'hi!' dad ")
	})


	//handlers

	bothandlers.MapRoutes(bot)

	bot.Start()

}
