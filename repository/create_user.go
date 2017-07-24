package repository

import (
	"RelationshipMatch/model"

	log "github.com/Sirupsen/logrus"
	"github.com/go-pg/pg"
)

const (
	CreateUserSQL  = `insert into users (id,name,type) values (?,?,?)`
	UserIsExistSQL = `select * from users where id=?`
)

func IsUserExist(pg *pg.DB, user_id string) bool {
	var u model.User

	res, err := pg.Query(&u, UserIsExistSQL, user_id)

	if err != nil {
		log.Errorf("Find user error: %s.", err)
		// If find user error, we think this user is existed.
		return true
	}

	if res.RowsReturned() > 0 {
		return true
	}

	return false
}

func CreateUser(pg *pg.DB, user *model.User) error {

	var u model.User

	_, err := pg.Query(&u, CreateUserSQL, user.Id, user.Name, user.Type)

	if err != nil {
		log.Errorf("Create users error: %s.", err)
		return err
	}

	return nil
}
