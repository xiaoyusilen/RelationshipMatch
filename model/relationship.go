package model

type Relationship struct {
	UserId  string `pg:"user_id"`
	OtherId string `pg:"other_id"`
	State   string `pg:"status"`
	Type    string `pg:"type"`
}

type UserRelationship struct {
	OtherId string `pg:"other_id"`
	Status  string `pg:"status"`
	Type    string `pg:"type"`
}
