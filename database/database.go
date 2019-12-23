package database

import "github.com/go-pg/pg/v9"

func New(options *pg.Options) *pg.DB {
	return pg.Connect(options)
}
