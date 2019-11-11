package bothandlers

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

func MapRoutes(bot *tb.Bot)  {
	bot.Handle("/hello", func(m *tb.Message) {
		_, _ = bot.Send(m.Sender, "Hi !")
		fmt.Println("user entered: "+ m.Payload)
	})
}
