package executor

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type ExecutorServer struct {
	db          *sql.DB
	putTasks    *cache.Cache
	updateTasks *cache.Cache
}

func NewExecutorServer(db *sql.DB) *ExecutorServer {
	return &ExecutorServer{
		db:          db,
		putTasks:    cache.New(time.Duration(executorConf.CacheExpire)*time.Minute, time.Duration(executorConf.CacheExpire+5)*time.Minute),
		updateTasks: cache.New(time.Duration(executorConf.CacheExpire)*time.Minute, time.Duration(executorConf.CacheExpire+5)*time.Minute),
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
	log.Println("Executor shutdown")
}

func (es *ExecutorServer) register(e *gin.Engine) {
	e.GET("/health", es.healthCheck)
	api := e.Group("/api")
	{
		api.POST("/peer", es.PreparePutPeer)
		api.PUT("/peer", es.UpdatePeer)
		api.DELETE("/peer", es.DeletePeer)
		api.GET("/peer", es.ShowPeer)
		api.GET("/resource", es.ResourceUsage)
	}
}

func (es *ExecutorServer) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// Insert //

func (es *ExecutorServer) PreparePutPeer(c *gin.Context) {
	req := &PreparePutPeerRequest{}

	if err := c.BindJSON(req); err != nil {
		logger.Error(err)
		c.JSON(400, gin.H{"message": "invalid json of PreparePutPeer"})
		return
	}
	if err := es.preparePutPeer(req); err != nil {
		logger.Error(err)
		c.JSON(500, gin.H{"message": "Failed insert peer information"})
	}

	c.JSON(200, &PreparePutPeerResponse{})
}

func (es *ExecutorServer) preparePutPeer(req *PreparePutPeerRequest) error {
	ctx := context.Background()
	task := &PutTask{
		tx:    es.db,
		state: StateNone,
	}
	if err := task.prepare(ctx, req); err != nil {
		return err
	}
	return nil
}

// Update //

func (es *ExecutorServer) UpdatePeer(c *gin.Context) {
	req := &UpdatePeerRequest{}
	if err := c.BindJSON(req); err != nil {
		logger.Warn(err)
		c.JSON(400, gin.H{"message": "invalid json of PreparePutPeer"})
		return
	}
	if err := es.updatePeer(req); err != nil {
		logger.Error(err)
		c.JSON(500, gin.H{"message": "Fatal update peer"})
		return
	}
	c.JSON(200, &UpdatePeerLocationResponse{})
}

func (es *ExecutorServer) updatePeer(req *UpdatePeerRequest) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Duration(executorConf.Mysql.Timeout)*time.Second)
	tx, err := es.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	task := &UpdateTask{
		tx:    tx,
		state: StateNone,
	}
	return task.prepare(ctx, req)
}

// Delete
func (es *ExecutorServer) DeletePeer(c *gin.Context) {
	req := &DeletePeerRequest{}
	if err := c.BindJSON(req); err != nil {
		logger.Warn(err)
		c.JSON(400, gin.H{"message": "invalid json of PreparePutPeer"})
		return
	}
	if err := es.deletePeer(req); err != nil {
		logger.Error(err)
		c.JSON(500, gin.H{"message": "Fatal delete peer"})
		return
	}
	c.JSON(200, &DeletePeerResponse{})
}

func (es *ExecutorServer) deletePeer(req *DeletePeerRequest) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Duration(executorConf.Mysql.Timeout)*time.Second)
	tx, err := es.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	return deletePeerTx(tx, req)
}

// Get
func (es *ExecutorServer) ShowPeer(c *gin.Context) {
	id := c.Query("peer_id")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(executorConf.Mysql.Timeout)*time.Second)
	defer cancel()
	p, err := selectRowPeer(es.db, id)
	if err != nil {
		c.JSON(404, gin.H{"message": fmt.Sprintf("not found peer=%d", id)})
		return
	}
	c.JSON(200, p)
}

// Util

func (es *ExecutorServer) ResourceUsage(c *gin.Context) {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}
	vcpu, err := cpu.Percent(time.Second*2, false)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	average := func(xs []float64) float64 {
		total := 0.0
		for _, v := range xs {
			total += v
		}
		return total / float64(len(xs))
	}
	cpuAverage := average(vcpu)

	resp := &ResourceUsageResponse{
		MemUsedPercent: vmem.UsedPercent,
		CpuUsedPercent: cpuAverage,
	}
	c.JSON(200, resp)
}
