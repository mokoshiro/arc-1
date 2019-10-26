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
		g.PUT("", h.Update)
		g.DELETE("", h.Delete)
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
	res, err := mh.MemberUC.GetMemberByRadius(req)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	c.JSON(200, res)
}

func (mh *MemberHandler) Update(c *gin.Context) {
	req := &model.UpdateRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithError(400, fmt.Errorf("Invalid update request"))
		return
	}
	res, err := mh.MemberUC.Update(req)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	c.JSON(200, res)
}

func (mh *MemberHandler) Delete(c *gin.Context) {
	req := &model.DeleteRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithError(400, fmt.Errorf("Invalid delete request"))
		return
	}
	res, err := mh.MemberUC.Delete(req)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	c.JSON(200, res)
}
