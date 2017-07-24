package service

import (
	"fmt"
	"testing"
)

const (
	InsertUserSQL = `insert into users (id,name,type) values (?,?,?)`
	GetUsersSQL   = `select * from users`
)

type User struct {
	Id   string `pg:"id"`
	Name string `pg:"name"`
	Type string `pg:"type"`
}

func TestConnectionPG(t *testing.T) {
	db := PGConnection()

	user := &User{
		Id:   "0x001",
		Name: "xiaoyusilen",
		Type: "user",
	}

	var u User

	_, err := db.Query(&u, InsertUserSQL, user.Id, user.Name, user.Type)

	if err != nil {
		fmt.Println(err)
	}

	//var u []User
	//_, err := db.Query(&u, GetUsersSQL)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(u)

}
