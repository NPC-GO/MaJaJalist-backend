package database

import "github.com/go-pg/pg/v9"

func New(options *pg.Options) *pg.DB {
	migration() //第一次連接先migration確保所有table都存在
	return pg.Connect(options)
}
