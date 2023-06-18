package bot

import (
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) getRouting() {
	b.telebot.Handle("/start", b.Info)
	b.telebot.Handle("/put", b.PutBeer)
	//b.telebot.Handle("/get", b.GetBeer)
	b.telebot.Handle("/watch", b.Watch)
	b.telebot.Handle(tele.OnText, b.TextHandler)
}
