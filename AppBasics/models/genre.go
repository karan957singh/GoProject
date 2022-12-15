package models

import "github.com/go-sql-driver/mysql"

type Genres struct {
	Name       string         `json:"name"` //name
	Id         int            `json:"id"`
	Created_at mysql.NullTime `json:"created_at"`
	Updated_at mysql.NullTime `json:"updated_at"`
	Deleted_at mysql.NullTime `json:"deleted_at"`
	// name string
	// id   int
}
