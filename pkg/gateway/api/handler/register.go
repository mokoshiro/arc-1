package handler

import (
	"fmt"

	"github.com/Bo0km4n/arc/pkg/gateway/domain/model"
	"github.com/Bo0km4n/arc/pkg/gateway/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterHandler struct {
	logger     *zap.Logger
	RegisterUC usecase.RegisterUsecase
}

func RegisterResource(e *gin.Engine, ruc usecase.RegisterUsecase, logger *zap.Logger) {
	h := RegisterHandler{RegisterUC: ruc, logger: logger}
	g := e.Group("/api/register")
	{
		g.POST("", h.Register)
	}
}

// Register POST /api/register
func (rh *RegisterHandler) Register(c *gin.Context) {
	req := &model.RegisterRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithError(400, fmt.Errorf("Invalid register request"))
		return
	}
	if err := rh.RegisterUC.Register(req); err != nil {
		rh.logger.Error(err.Error())
		c.AbortWithError(503, fmt.Errorf("Failed register process"))
		return
	}
	c.JSON(200, &model.RegisterResponse{})
}
