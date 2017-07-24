package model

type User struct {
	Id   string `pg:"id"`
	Name string `pg:"name"`
	Type string `pg:"type"`
}
