package repository

import (
	"RelationshipMatch/model"

	log "github.com/Sirupsen/logrus"
	"github.com/go-pg/pg"
)

const (
	GetUsersSQL = `select * from users`
)

func GetUsers(pg *pg.DB) (users []model.User, err error) {

	_, err = pg.Query(&users, GetUsersSQL)

	if err != nil {
		log.Errorf("Get users error: %s.", err)
		return users, err
	}

	return users, nil
}
