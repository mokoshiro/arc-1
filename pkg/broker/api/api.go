package api

import (
	"log"

	"github.com/Bo0km4n/arc/pkg/broker/msg"
	"github.com/Bo0km4n/arc/pkg/broker/socket"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type brokerAPI struct {
	logger *zap.Logger
}

type brokerHTTPAPI struct {
}

func NewGRPCTrackerAPI() *brokerAPI {
	return &brokerAPI{}
}

func NewHTTPTrackerAPI() *brokerHTTPAPI {
	return &brokerHTTPAPI{}
}

func (a *brokerHTTPAPI) Run() {
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
	e.POST(
		"/api/signaling", a.Signaling,
	)
	e.POST(
		"/api/signaling/peer/status", a.PingPeerStatus,
	)
	e.POST(
		"/api/room/notification", a.RoomNotification,
	)
	e.Run(":8000")
}

func (a *brokerHTTPAPI) Register(c *gin.Context) {
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

func (a *brokerHTTPAPI) LookupMembers(c *gin.Context) {
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

func (a *brokerHTTPAPI) PutMember(c *gin.Context) {
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

func (a *brokerHTTPAPI) Signaling(c *gin.Context) {
	m := &msg.SignalingRequest{}
	if err := c.BindJSON(m); err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}
	res, err := socket.Signaling(m)
	if err != nil {
		log.Println(err)
		c.AbortWithError(404, err)
		return
	}
	c.JSON(200, res)
}

func (a *brokerHTTPAPI) RoomNotification(c *gin.Context) {
	m := &msg.RoomNotificationRequest{}
	if err := c.BindJSON(m); err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}
	res, err := socket.RoomNotification(m)
	if err != nil {
		log.Println(err)
		c.AbortWithError(404, err)
		return
	}
	c.JSON(200, res)
}

func (a *brokerHTTPAPI) PingPeerStatus(c *gin.Context) {

}
