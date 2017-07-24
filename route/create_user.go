package route

import (
	"fmt"

	"RelationshipMatch/model"
	"RelationshipMatch/repository"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// CreateUser
//
// POST /users
//
// Request body
//{
//	"id": "11231244213",
//	"name": "Alice" ,
//	"type": "user"
//}

type CreateUserReq struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (api *RestApi) CreateUser(c *gin.Context) {

	params := CreateUserReq{}

	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(200, gin.H{
			"result": "Json format error.",
		})
		return
	}

	// TODO: add validation
	user := &model.User{
		Id:   params.Id,
		Name: params.Name,
		Type: params.Type,
	}

	// Validate user is existed or not
	isExist, err := repository.IsUserExist(api.PG, user.Id)

	if err != nil {
		result := fmt.Sprintf("Validate user error: %s.", err)
		c.JSON(200, gin.H{
			"result": result,
		})
		return
	}

	if isExist {
		c.JSON(200, gin.H{
			"result": "This user is existed.",
		})
		return
	}

	err = repository.CreateUser(api.PG, user)

	if err != nil {
		result := fmt.Sprintf("Create user error: %s.", err)
		c.JSON(200, gin.H{
			"result": result,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "Create user successed.",
	})
	log.Debugf("Create user successed.")

	return
}
