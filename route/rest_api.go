package route

import (
	"RelationshipMatch/common/service"
	"RelationshipMatch/config"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

type RestApi struct {
	Router *gin.Engine

	// Configurations
	Config *config.Config

	// PG
	PG *pg.DB
}

func HandleRest(cfg *config.Config) *RestApi {

	r := &RestApi{

		Router: gin.Default(),

		Config: cfg,

		PG: service.PGConnection(),
	}

	r.Router.GET("/users", r.GetUsers)
	r.Router.POST("/users", r.CreateUser)
	r.Router.GET("/users/:user_id/relationships", r.GetUserRelationship)
	r.Router.PUT("/users/:user_id/relationships/:other_user_id", r.CreateUserRelationship)
	return r
}
