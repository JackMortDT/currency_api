package app

import (
	"os"

	"github.com/go-pg/pg/v10"
)

func startDatabase() *pg.DB {
	addr := os.Getenv("ADDR")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")

	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: database,
	})

	return db
}
