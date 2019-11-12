package bothandlers

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

//buttons
var BtnMoon = tb.InlineButton{
	Unique: "moon",
	Text:   "Moon 🌹",
}

var BtnSun = tb.InlineButton{
	Unique: "sun",
	Text:   "Sun 🌞",
}

//define keyboard for buttons
var BtnInlineKeys = [][]tb.InlineButton{
	[]tb.InlineButton{BtnMoon, BtnSun},
}

func InitButtons(bot *tb.Bot) {
	bot.Handle(&BtnMoon, func(c *tb.Callback) {
		_ = bot.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})

		bot.Send(c.Sender, "Moon says hi!")
	})
	bot.Handle(&BtnSun, func(c *tb.Callback) {
		_ = bot.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})

		bot.Send(c.Sender, "Sun says 'hi!' dad ")
	})
}
