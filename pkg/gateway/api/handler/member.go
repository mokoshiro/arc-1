package handler

import (
	"fmt"

	"github.com/Bo0km4n/arc/pkg/gateway/domain/model"
	"github.com/Bo0km4n/arc/pkg/gateway/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MemberHandler struct {
	logger   *zap.Logger
	MemberUC usecase.MemberUsecase
}

func MemberResource(e *gin.Engine, muc usecase.MemberUsecase, logger *zap.Logger) {
	h := MemberHandler{MemberUC: muc, logger: logger}
	g := e.Group("/api/member")
	{
		g.POST("/register", h.Register)
		g.GET("", h.GetMemberByRadius)
	}
}

// Register POST /api/register
func (rh *MemberHandler) Register(c *gin.Context) {
	req := &model.RegisterRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithError(400, fmt.Errorf("Invalid register request"))
		return
	}
	if err := rh.MemberUC.Register(req); err != nil {
		rh.logger.Error(err.Error())
		c.AbortWithError(503, fmt.Errorf("Failed register process"))
		return
	}
	c.JSON(200, &model.RegisterResponse{})
}

func (mh *MemberHandler) GetMemberByRadius(c *gin.Context) {
	req := &model.GetMemberByRadiusRequest{}
	if err := c.BindQuery(req); err != nil {
		c.AbortWithError(400, fmt.Errorf("Inavlid get member by radius request"))
		return
	}
	c.JSON(200, gin.H{"message": "ok"})
}
