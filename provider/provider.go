package provider

import "sync"

type Database struct {
	dsn string
}

func (db *Database) Use() {
	println("数据库被使用了")
}

var once sync.Once
var db *Database

func NewDB() *Database {
	once.Do(func() {
		db = &Database{}
	})
	return db
}
