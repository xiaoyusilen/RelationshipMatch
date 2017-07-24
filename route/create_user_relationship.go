package route

import (
	"fmt"

	"RelationshipMatch/model"
	"RelationshipMatch/repository"

	"github.com/gin-gonic/gin"
)

// CreateUserRelationship
//
// PUT /users/:user_id/relationships/:other_user_id
//
// Request body
// {
//	"user_id": "21341231231",
//	"state": "liked" ,
//	"type": "relationship"
// }

type CreateUserRelationshipReq struct {
	UserId string `json:"user_id"`
	State  string `json:"state"`
	Type   string `json:"type"`
}

func (api *RestApi) CreateUserRelationship(c *gin.Context) {

	user_id := c.Param("user_id")

	other_user_id := c.Param("other_user_id")

	params := CreateUserRelationshipReq{}

	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(200, gin.H{
			"result": "Json format error.",
		})
		return
	}

	// State only allowed fields liked | disliked
	if params.State != "liked" && params.State != "disliked" {
		c.JSON(200, gin.H{
			"result": "State format error.",
		})
		return
	}

	// validate user_id and other_user_id
	is_user_exist := repository.IsUserExist(api.PG, user_id)
	if !is_user_exist {
		result := fmt.Sprintf("User id %s is not exist.", user_id)
		c.JSON(200, gin.H{
			"result": result,
		})
		return
	}

	is_other_user_id_exist := repository.IsUserExist(api.PG, other_user_id)
	if !is_other_user_id_exist {
		result := fmt.Sprintf("User id %s is not exist.", other_user_id)
		c.JSON(200, gin.H{
			"result": result,
		})
		return
	}

	relationship := &model.Relationship{
		UserId:  user_id,
		OtherId: other_user_id,
		State:   params.State,
		Type:    params.Type,
	}

	// add relationship to database
	res := repository.CreateUserRelationship(api.PG, relationship)

	if !res {
		c.JSON(200, gin.H{
			"result": "Create relationship failed.",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "Create relationship successed.",
	})
	return

}
