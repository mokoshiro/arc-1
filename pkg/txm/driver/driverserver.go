package driver

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
	"math"

	"github.com/Bo0km4n/arc/pkg/txm/executor"
	"github.com/Bo0km4n/arc/pkg/txm/executor/client"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

type DriverServer struct {
	geoHashAccuracy  uint
	geoHashMaxRadius float32
	executorDNS      *executorDNSClient
	locationHistory  *locationHistory
	httpClient       *http.Client
}


func NewDriverServer(db *sql.DB, kvs *redis.Pool) *DriverServer {
	tr := &http.Transport{
		MaxIdleConns:       512,
		MaxConnsPerHost:    256,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	httpclient := &http.Client{Transport: tr}
	return &DriverServer{
		httpClient:       httpclient,
		geoHashAccuracy:  uint(driverConf.GeoHashAccuracy),
		geoHashMaxRadius: getMaxRadius(driverConf.GeoHashAccuracy),
		executorDNS:      newExecutorDNSClient(db),
		locationHistory:  newLocationHistory(kvs),
	}
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
	api := e.Group("/api")
	//fmt.Println(ds.executorDNS)
	{
		api.POST("/peer", ds.StorePeer)
		api.PUT("/peer/location", ds.UpdatePeerLocation)
		api.GET("/peer", ds.ShowPeer)
		api.GET("/peer/lookup", ds.LookupPeer)
	}
}

func (ds *DriverServer) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (ds *DriverServer) StorePeer(c *gin.Context) {
	req := &RegisterRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		logger.Error(err.Error())
		return
	}
	if err := ds.storePeer(req); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		logger.Error(err.Error())
		return
	}
}

func (ds *DriverServer) UpdatePeerLocation(c *gin.Context) {
	req := &UpdatePeerLocationRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		logger.Error(err.Error())
		return
	}
	if err := ds.updatePeerLocation(req); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		logger.Error(err.Error())
		return
	}
}

func (ds *DriverServer) storePeer(req *RegisterRequest) error {
	// Step 1: parse target geohash
	hash := encodeGeoHash(req.Latitude, req.Longitude, ds.geoHashAccuracy)
	//fmt.Println(hash)
	// Step 2: get executor host address by geohash
	executorAddr := ds.resolveExecutorAddress(hash,req.Latitude,req.Longitude)
	fmt.Println(executorAddr)
	// Step 3: send request to store peer information
	executorClient := client.NewExecutorClient(ds.httpClient, "http://"+executorAddr)
	ctx := context.Background()
	storePeerRequest := &executor.PreparePutPeerRequest{
		PeerID:     req.PeerID,
		Addr:       req.Addr, //ピアのアドレス
		Credential: "dummy-credential",
		Longitude:  req.Longitude,
		Latitude:   req.Latitude,
	}
	if _, err := executorClient.StorePeer(ctx, storePeerRequest); err != nil {
		return err
	}
	// Step 4: store pair of <peer : executor> to locationHistory
	//peerを保存するredisの位置を覚えておく
	return ds.locationHistory.Put(req.PeerID, executorAddr)
}

func (ds *DriverServer) updatePeerLocation(req *UpdatePeerLocationRequest) error {
	// Step 1: send request to get prev peer information
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	prevExecutorAddr, err := ds.locationHistory.Get(req.PeerID)
	if err != nil {
		return err
	}
	prevc := client.NewExecutorClient(ds.httpClient, "http://"+prevExecutorAddr)
	prevPeerRow, err := prevc.SelectPeer(ctx, req.PeerID)
	if err != nil {
		return err
	}
	// Step 2: delete prev peer informatio
	if _, err := prevc.DeletePeer(ctx, &executor.DeletePeerRequest{PeerID: req.PeerID}); err != nil {
		return err
	}
	// Step 3: parse target geohash
	hash := encodeGeoHash(req.Latitude, req.Longitude, ds.geoHashAccuracy)
	// Step 4: get executor host by geohash
	executorAddr := ds.resolveExecutorAddress(hash,req.Latitude,req.Longitude)
	// Step 5: send update request to executor
	executorClient := client.NewExecutorClient(ds.httpClient, "http://"+executorAddr)
	execReq := &executor.UpdatePeerRequest{
		PeerID:     prevPeerRow.PeerID,
		Longitude:  req.Longitude,
		Latitude:   req.Latitude,
		Credential: prevPeerRow.Credential,
		Addr:       prevPeerRow.Addr,
	}
	if _, err := executorClient.UpdatePeerLocation(ctx, execReq); err != nil {
		return err
	}
	// Step 6: update pair of <peer : executor> to locationHistory
	return ds.locationHistory.Put(req.PeerID, executorAddr)
}

func (ds *DriverServer) ShowPeer(c *gin.Context) {
	id := c.Query("id")
	executorAddr, err := ds.locationHistory.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"message": fmt.Sprintf("not found peer=%s", id)})
		return
	}
	executorClient := client.NewExecutorClient(ds.httpClient, "http://"+executorAddr)
	p, err := executorClient.SelectPeer(context.Background(), id)
	if err != nil {
		c.JSON(404, gin.H{"message": fmt.Sprintf("not found peer=%s", id)})
		return
	}
	c.JSON(200, p)
}

func (ds *DriverServer) LookupPeer(c *gin.Context) {
	req := &LookupRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		logger.Error(err.Error())
		return
	}
	res, err := ds.lookupPeer(req)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		logger.Error(err.Error())
		return
	}
	c.JSON(200, res)
}

func (ds *DriverServer) lookupPeer(req *LookupRequest) ([]string, error) {
	var indexs []string
	var peer []string
	indexs = lookupGeoHashIndexes(req.Latitude,req.Longitude,3,req.Radius)
	fmt.Println(indexs)
	for _,hash := range indexs{
		exists, err :=ds.isExistExcutor(hash)
		if err != nil{
			log.Fatal(err)
			return nil,err
		}else if !exists{
			//not exist
			fmt.Println(hash," not exist")
		}else{
			//executorにリクエストをおくる
			fmt.Println(hash," exist")
			addr := ds.getExecutorAddress(hash)
			executorClient := client.NewExecutorClient(ds.httpClient, "http://"+addr)
			execReq := &executor.LookupRequest{
				Radius:     req.Radius,
				Longitude:  req.Longitude,
				Latitude:   req.Latitude,
			}
			p, err := executorClient.GetPeer(context.Background(), execReq)
			if err != nil{
				return nil,err
			}
			peer = append(peer,p)
		}
	}

	fmt.Println(peer)
	return peer, nil
}

func (ds *DriverServer) getExecutorAddress(hash string)string {
	query := "SELECT host from executor where geohash = ?"
	var host string
	err := ds.executorDNS.db.QueryRow(query,hash).Scan(&host)
	if err != nil{
		log.Fatal(err)
	}
	//fmt.Println(host)
	return host
}

func (ds *DriverServer) resolveExecutorAddress(hash string,lng float64,lat float64) string {
	// TODO: Implement
	//fmt.Println(hash)
	exists, err := ds.isExistExcutor(hash)
	if err != nil {
		log.Fatal(err)
		return "err"
	}else if !exists{
		fmt.Println("not exist")
		//ds.newExecutorDNS(1)
		host := ds.newExecutorDNS(lng,lat)
		//host = host + ":" + port
		return ds.putNewHash(host,hash)
	}else{
		fmt.Println("exist")
		//return address
		var host string
		query := "SELECT host from executor where  geohash = ?"
		err:= ds.executorDNS.db.QueryRow(query, hash).Scan(&host)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println("host=",host)
		return host
	}
}

func (ds *DriverServer) isExistExcutor(hash string) (bool, error) {
	var exists bool
	
	query := "SELECT exists (SELECT host from executor where  geohash = ?)"
	//fmt.Println(query)
	err := ds.executorDNS.db.QueryRow(query, hash).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}else{
	}
	return exists, nil
}

func (ds *DriverServer) newExecutorDNS(lng float64,lat float64)string{
	query := "SELECT host ,longitude ,latitude from executor where  geohash = ?"
	//用意されたDB
	geohash := "@"
	var host []string
	var distance []float64
	rows, err := ds.executorDNS.db.Query(query,geohash)
	if err != nil{
		log.Fatal(err)
	}
	for rows.Next(){
		type Executor struct{
			Host string
			Lng float64
			Lat float64
		}
		var executor Executor
		rows.Scan(&executor.Host,&executor.Lng,&executor.Lat)
		host = append(host,executor.Host)
		distance = append(distance,ds.distance(lng,lat,executor.Lng,executor.Lat))
	}
	return host[ds.min(distance)]
}

func (ds *DriverServer)putNewHash(host string ,hash string) string{
	statement := fmt.Sprintf("INSERT INTO executor(host,geohash) VALUES (?, ?)")
	ins, err := ds.executorDNS.db.Prepare(statement)
	if err != nil{
		log.Fatal(err)
		return "err"
	}
	_, err = ins.Exec(host,hash)
	if err != nil {
		log.Fatal(err)
		return "err"
	}
	return host
}

func (ds *DriverServer)distance(lng1 float64,lat1 float64,lng2 float64,lat2 float64) float64{
	return math.Sqrt((lng1-lng2)*(lng1-lng2)+(lat1-lat2)*(lat1-lat2))
}

func (ds *DriverServer)min(distance []float64) int {
	num := 0
	min := distance[num]
	for i, _ := range distance {
		if distance[i] < min{
			min = distance[i]
			num = i
		}
	}
	return num
}