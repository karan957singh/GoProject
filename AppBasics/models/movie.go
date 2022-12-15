package models

import "github.com/go-sql-driver/mysql"

type Movies struct {
	Id              int            `json:"id"`
	Title           string         `json:"title"`
	Genre_id        int            `json:"genre_id"`
	NumberInStock   int            `json:"numberInStock"`
	DailyRentalRate float32        `json:"dailyRentalRate"`
	Liked           int            `json:"liked"`
	Created_at      mysql.NullTime `json:"created_at"`
	Updated_at      mysql.NullTime `json:"updated_at"`
	Deleted_at      mysql.NullTime `json:"deleted_at"`
}
