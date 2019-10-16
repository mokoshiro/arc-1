package router

import (
	"github.com/Bo0km4n/arc/pkg/gateway/api/handler"
	"github.com/Bo0km4n/arc/pkg/gateway/domain/repository"
	"github.com/Bo0km4n/arc/pkg/gateway/usecase"
	"github.com/gin-gonic/gin"
)

// New returns a router object binding some functions.
func New() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	{
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"Message": "pong",
			})
		})
	}
	{
		// Register
		metadataRepo := repository.NewMetadataRepository()
		trackerRepo := repository.NewTrackerRepository()
		ruc := usecase.NewRegisterUsecase(metadataRepo, trackerRepo)
		handler.RegisterResource(r, ruc)
	}
	return r
}
