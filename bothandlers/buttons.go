package bothandlers

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

//buttons
var BtnMoon = tb.InlineButton{
	Unique: "Moon",
	Text:   "Moon 🌹",
}

var BtnSun = tb.InlineButton{
	Unique: "Sun",
	Text:   "Sun 🌞",
}

//define keyboard for buttons
var BtnInlineKeys = [][]tb.InlineButton{
	{BtnMoon, BtnSun},
}

func InitButtons(bot *tb.Bot) {
	// buttons
	bot.Handle(&BtnMoon, func(c *tb.Callback) {
		_ = bot.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})

		_, _ = bot.Send(c.Sender, "Moon says hi!")
	})
	bot.Handle(&BtnSun, func(c *tb.Callback) {
		_ = bot.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})

		_, _ = bot.Send(c.Sender, "Sun says 'hi!' dad ")
	})
}
