package repository

import (
	"RelationshipMatch/model"

	log "github.com/Sirupsen/logrus"
	"github.com/go-pg/pg"
)

const (
	GetUserRelationshipSQL = `select other_id, status, type from relationship where user_id=?`
)

func GetUserRelationship(pg *pg.DB, user_id string) (relationship []model.UserRelationship, err error) {

	_, err = pg.Query(&relationship, GetUserRelationshipSQL, user_id)

	if err != nil {
		log.Errorf("Get users error: %s.", err)
		return relationship, err
	}

	return relationship, nil
}
