package handler

import (
	"github.com/Bo0km4n/arc/pkg/gateway/usecase"
	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	RegisterUC usecase.RegisterUsecase
}

func RegisterResource(e *gin.Engine, ruc usecase.RegisterUsecase) {
	h := RegisterHandler{RegisterUC: ruc}
	e.Group("/api/register")
	{
		e.POST("", h.Register)
	}
}

// Register POST /api/register
func (rh *RegisterHandler) Register(c *gin.Context) {
	c.JSON(200, gin.H{"message": "registered"})
}
