package service

import (
	"RelationshipMatch/config"

	"github.com/go-pg/pg"
)

func PGConnection() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     config.DefaultPGAddr,
		User:     config.DefaultPGUser,
		Password: "",
		Database: config.DefaultPGDB,
	})
	return db
}
