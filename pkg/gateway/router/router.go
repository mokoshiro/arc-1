package router

import "github.com/gin-gonic/gin"

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
	return r
}
