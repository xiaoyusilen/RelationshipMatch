package route

import (
	"fmt"

	"RelationshipMatch/repository"

	"github.com/gin-gonic/gin"
)

// GetUserRelationship
//
// GET /users/:user_id/relationships
//
// Response body
//	{
//		"result": [
//			{
//				"user_id": "222333444",
//				"state": "liked" ,
//				"type": "relationship"
//			},
//			{
//				"user_id": "333222444",
//				"state": "matched" ,
//				"type": "relationship"
//			},
//			{
//				"user_id": "444333222",
//				"state": "disliked" ,
//				"type": "relationship"
//			}
//		]
//	}

func (api *RestApi) GetUserRelationship(c *gin.Context) {
	user_id := c.Param("user_id")

	relationship, err := repository.GetUserRelationship(api.PG, user_id)

	if err != nil {
		result := fmt.Sprintf("Get user relationship error: %s.", err)
		c.JSON(200, gin.H{
			"result": result,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": relationship,
	})
	return

}
