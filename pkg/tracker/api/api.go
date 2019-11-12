package api

import (
	"log"

	"github.com/Bo0km4n/arc/pkg/tracker/msg"
	"github.com/Bo0km4n/arc/pkg/tracker/socket"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type trackerAPI struct {
	logger *zap.Logger
}

type trackerHTTPAPI struct {
}

func NewGRPCTrackerAPI() *trackerAPI {
	return &trackerAPI{}
}

func NewHTTPTrackerAPI() *trackerHTTPAPI {
	return &trackerHTTPAPI{}
}

func (a *trackerHTTPAPI) Run() {
	e := gin.Default()
	e.POST(
		"/api/member", a.Register,
	)
	e.GET(
		"/api/member", a.LookupMembers,
	)
	e.PUT(
		"/api/member", a.PutMember,
	)
	e.Run(":8000")
}

func (a *trackerHTTPAPI) Register(c *gin.Context) {
	m := &msg.GreetMessage{}
	if err := c.BindJSON(m); err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}
	if err := socket.Greet(m); err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

func (a *trackerHTTPAPI) LookupMembers(c *gin.Context) {
	m := &msg.LookupRequest{}
	if err := c.BindJSON(m); err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}
	res, err := socket.Lookup(m)
	if err != nil {
		log.Println(err)
		c.AbortWithError(404, err)
		return
	}
	c.JSON(200, res)
}

func (a *trackerHTTPAPI) PutMember(c *gin.Context) {
	m := &msg.TrackingMessage{}
	if err := c.BindJSON(m); err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}
	if err := socket.Tracking(m); err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}
