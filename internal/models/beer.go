package models

import "time"

type Beer struct {
	ID         int64     `sql:"id"`
	Title      string    `sql:"title"`
	ABV        string    `sql:"abv"`
	AddingTime time.Time `sql:"adding_time"`
}
