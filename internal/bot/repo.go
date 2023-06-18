package bot

import tele "gopkg.in/telebot.v3"

func (b *Bot) Info(c tele.Context) error {
	return c.Send("/put - добавить пивасик\n/watch - посмотреть содержимое холодильника")
}

func (b *Bot) PutBeer(c tele.Context) error {
	b.botCtx = "gettingTitle"
	return c.Send("Введите название")
}

//func (b *Bot) GetBeer(c tele.Context) error {
//	return b.repo.GetBeer(c)
//}

func (b *Bot) Watch(c tele.Context) error {
	b.botCtx = ""
	return b.repo.Watch(c)
}

func (b *Bot) TextHandler(c tele.Context) error {
	text := c.Text()
	switch b.botCtx {
	case "gettingTitle":
		b.cachedBeer.Title = text
		b.botCtx = "gettingABV"
		return c.Send("Введите кол-во градусов")
	case "gettingABV":
		b.cachedBeer.ABV = text
		b.botCtx = ""
		return b.repo.PutBeer(c, b.cachedBeer)
	default:
	}

	return nil
}
