package route

import (
	"RelationshipMatch/repository"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// GetUsers
//
// GET /users
//
// Response body
// "result": [
//	{
//		"id": "21341231231",
//		"name": "Bob" ,
//		"type": "user"
//	},
//	{
//		"id": "31231242322",
//		"name": "Samantha" ,
//		"type": "user"
//	}
// ]

func (api *RestApi) GetUsers(c *gin.Context) {

	// TODO: add validation
	users := repository.GetUsers(api.PG)

	c.JSON(200, gin.H{
		"result": users,
	})
	log.Debugf("Get Users successed.")

	return
}
