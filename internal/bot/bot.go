package bot

import (
	tele "gopkg.in/telebot.v3"
	"telebot_BeerRefrigerator/internal/models"
	"telebot_BeerRefrigerator/internal/refrigerator"
	"time"
)

type Bot struct {
	telebot    *tele.Bot
	cfg        *BotConfig
	repo       *refrigerator.Repo
	botCtx     string
	cachedBeer *models.Beer
}

func New() (bot *Bot, err error) {
	bot = &Bot{cachedBeer: &models.Beer{}}
	bot.cfg, err = getBotConfig()

	bot.repo, err = refrigerator.NewRepo()

	pref := tele.Settings{
		Token:  bot.cfg.token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot.telebot, err = tele.NewBot(pref)
	if err != nil {
		return
	}

	bot.getRouting()

	return
}

func (b *Bot) Start() {
	b.telebot.Start()
}
