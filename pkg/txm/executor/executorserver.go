package executor

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type ExecutorServer struct {
	db *sql.DB
}

func NewExecutorServer(db *sql.DB) *ExecutorServer {
	return &ExecutorServer{
		db: db,
	}
}

func (es *ExecutorServer) run() {
	app := gin.Default()
	es.register(app)
	srv := &http.Server{
		Addr:    ":" + executorConf.Port,
		Handler: app,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Fatalln("Server closed with error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Println("Failed to gracefully shutdown:", err)
	}
	log.Println("Driver shutdown")
}

func (es *ExecutorServer) register(e *gin.Engine) {
	e.GET("/health", es.healthCheck)
	api := e.Group("/api")
	{
		api.POST("/peer/prepare", es.PreparePutPeer)
		api.POST("/peer/commit", es.CommitPutPeer)
		api.POST("/peer/rollback", es.RollbackPutPeer)
	}
}

func (es *ExecutorServer) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (es *ExecutorServer) PreparePutPeer(c *gin.Context) {
	req := &PreparePutPeerRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"message": "invalid json on PreparePutPeer"})
		return
	}
}

func (es *ExecutorServer) CommitPutPeer(c *gin.Context) {
	req := &CommitPutPeerRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"message": "invalid json on PreparePutPeer"})
		return
	}
}
func (es *ExecutorServer) RollbackPutPeer(c *gin.Context) {
	req := &RollbackPutPeerRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"message": "invalid json on PreparePutPeer"})
		return
	}
}
