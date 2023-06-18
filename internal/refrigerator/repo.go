package refrigerator

import (
	tele "gopkg.in/telebot.v3"
	"strconv"
	"telebot_BeerRefrigerator/internal/database"
	"telebot_BeerRefrigerator/internal/models"
	"time"
)

type Repo struct {
	db *database.DB
}

func NewRepo() (*Repo, error) {
	dbConfig := database.GetDefaultDBConfig()

	db, err := database.NewDatabaseConnection(dbConfig)
	if err != nil {
		return &Repo{}, err
	}

	return &Repo{db: db}, nil
}

func (r *Repo) PutBeer(c tele.Context, beer *models.Beer) error {
	beer.AddingTime = time.Now()

	id, err := r.db.InsertBeer(beer)
	if err != nil {
		return err
	}

	return c.Send("your beer id: " + strconv.Itoa(id))
}

//func (r *Repo) GetBeer(c tele.Context) error {
//
//	return c.Send("put!")
//}

func (r *Repo) Watch(c tele.Context) error {
	beers, err := r.db.GetAll()
	if err != nil {
		return err
	}

	beerList := ""

	for i, beer := range beers {
		beerList += "beer #" + strconv.Itoa(i+1) + ":\ntitle: " + beer.Title + "\nabv: " + beer.ABV + "\nwas added: " + beer.AddingTime.Format(time.RFC3339) + "\n\n"
	}

	return c.Send(beerList)
}
