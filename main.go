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

	//buttons
	btnMoon := tb.InlineButton{
		Unique: "Moon",
		Text:   "Moon 🌹",
	}

	btnSun := tb.InlineButton{
		Unique: "Sun",
		Text:   "Sun 🌞",
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

	//define keyboard for buttons
	inlineKeys := [][]tb.InlineButton{
		[]tb.InlineButton{btnMoon, btnSun},
	}
	bot.Handle("/pick_time", func(m tb.Message) {
		_, _ = bot.Send(m.Sender, "Day or night ?", &tb.ReplyMarkup{InlineKeyboard: inlineKeys})
	})

	//handlers

	bothandlers.MapRoutes(bot)

	bot.Start()

}
