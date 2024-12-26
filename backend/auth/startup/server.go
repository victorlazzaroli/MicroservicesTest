package startup

import (
	"github.com/gin-gonic/gin"
	"github.com/victorlazzaroli/microservicesTest/auth/config"
)

func NewServer(env *config.Env) *gin.Engine {
	router := gin.Default()

	return router
}
