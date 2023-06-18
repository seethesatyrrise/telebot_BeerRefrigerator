package database

import (
	"context"
	"github.com/jackc/pgx"
	"telebot_BeerRefrigerator/internal/models"
	"time"
)

type DB struct {
	Handler *pgx.Conn
}

func NewDatabaseConnection(c *DBConfig) (*DB, error) {
	db, err := pgx.Connect(pgx.ConnConfig{
		Host:     c.PGAddress,
		User:     c.PGUser,
		Password: c.PGPassword,
		Database: c.PGDatabase,
	})
	if err != nil {
		return nil, err
	}

	return &DB{Handler: db}, nil
}

func (db *DB) InsertBeer(beer *models.Beer) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var id int
	row := db.Handler.QueryRowEx(ctx,
		`INSERT INTO beer_bot (title, abv, adding_time)
				VALUES ($1, $2, $3)
				RETURNING id;`,
		nil, beer.Title, beer.ABV, beer.AddingTime)

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *DB) GetAll() ([]models.Beer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	beers := make([]models.Beer, 0)

	rows, err := db.Handler.QueryEx(ctx,
		`select * from beer_bot;`, nil,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		beer := new(models.Beer)
		if err := rows.Scan(&beer.ID, &beer.Title, &beer.ABV, &beer.AddingTime); err != nil {
			return nil, err
		}
		beers = append(beers, *beer)
	}

	return beers, nil
}
