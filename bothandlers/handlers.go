package bothandlers

import (
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
)

func MapRoutes(bot *tb.Bot) {
	// hello
	bot.Handle("/hello", func(m *tb.Message) {
		bot.Send(m.Sender, "Hi !")
		fmt.Printf("user entered: %v", m.Payload)
	})

	// pick_time
	bot.Handle("/picktime", func(m tb.Message) {
		bot.Send(
			m.Sender,
			"Day or night ?",
			&tb.ReplyMarkup{InlineKeyboard: BtnInlineKeys},
		)
	})
}
