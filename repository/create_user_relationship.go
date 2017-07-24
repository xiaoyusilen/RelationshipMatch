package repository

import (
	"RelationshipMatch/model"

	log "github.com/Sirupsen/logrus"
	"github.com/go-pg/pg"
)

const (
	GetRelationshipSQL    = `select status from relationship where user_id=? and other_id=?`
	InsertRelationshipSQL = `insert into relationship (user_id, other_id, status, type) values (?,?,?,?)`
	UpdateRelationshipSQL = `update relationship set status=? where user_id=? and other_id=?`
)

func CreateUserRelationship(pg *pg.DB, relationship *model.Relationship) (bool, error) {

	// if state is disliked, add to database and need no query
	if relationship.State == "disliked" {
		_, err := pg.Query(&relationship, InsertRelationshipSQL, relationship.UserId, relationship.OtherId, relationship.State, relationship.Type)
		if err != nil {
			log.Errorf("Insert relationship error: %s.", err)
			return false, err
		}
		return true, nil
	}

	// Get other_user_id's state to user_id
	var state_from string
	res, err := pg.Query(&state_from, GetRelationshipSQL, relationship.OtherId, relationship.UserId)
	if err != nil {
		log.Errorf("Get relationship error: %s.", err)
		return false, err
	}

	// if return rows > 0 means has relationship.
	if res.RowsReturned() > 0 {
		if state_from == "liked" {
			_, err = pg.Query(&relationship, InsertRelationshipSQL, relationship.UserId, relationship.OtherId, "matched", relationship.Type)
			if err != nil {
				log.Errorf("Insert relationship error: %s.", err)
				return false, err
			}
			_, err = pg.Query(&relationship, UpdateRelationshipSQL, "matched", relationship.OtherId, relationship.UserId)
			if err != nil {
				log.Errorf("Update relationship error: %s.", err)
				return false, err
			}
			return true, nil
		}
	}

	// have no relationship equals disliked.
	_, err = pg.Query(&relationship, InsertRelationshipSQL, relationship.UserId, relationship.OtherId, relationship.State, relationship.Type)
	if err != nil {
		log.Errorf("Insert relationship error: %s.", err)
		return false, err
	}
	return true, nil
}
