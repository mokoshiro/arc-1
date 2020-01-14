package driver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

type DriverServer struct {
	geoHashAccuracy  int
	geoHashMaxRadius float32
	executorDNS      *redis.Pool
	executorClients  *executorClients
}

func NewDriverServer(kvs *redis.Pool) *DriverServer {
	return &DriverServer{
		geoHashAccuracy:  driverConf.GeoHashAccuracy,
		geoHashMaxRadius: getMaxRadius(driverConf.GeoHashAccuracy),
		executorDNS:      kvs,
		executorClients:  nil,
	}
}

var accRadius = map[int]float32{ // km
	1: 4989.60000,
	2: 1012.50000,
	3: 155.92500,
	4: 31.64062,
	5: 4872.66,
	6: 0.98877,
	7: 0.15227,
}

func getMaxRadius(acc int) float32 {
	if acc > 7 {
		return accRadius[7]
	}
	return accRadius[acc]
}

func (ds *DriverServer) run() {
	app := gin.Default()
	ds.register(app)
	srv := &http.Server{
		Addr:    ":" + driverConf.Port,
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

func (ds *DriverServer) register(e *gin.Engine) {
	e.GET("/health", ds.healthCheck)
}

func (ds *DriverServer) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
