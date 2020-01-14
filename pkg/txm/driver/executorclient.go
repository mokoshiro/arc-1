package driver

import (
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/patrickmn/go-cache"
)

type executorClient struct {
	target              string
	timeout             time.Duration
	maxIdleConnsPerHost int
	conn                *http.Client
}

type executorClients struct {
	executorDNS *redis.Pool
	c           *cache.Cache
}

func newExecutorClients() *executorClients {
	newClients := &executorClients{
		executorDNS: kvs(
			driverConf.Redis.Host,
			driverConf.Redis.MaxIdle,
			driverConf.Redis.Active,
			driverConf.Redis.IdleTimeout,
		),
		c: cache.New(time.Duration(driverConf.CacheExpire)*time.Minute, time.Duration(driverConf.CacheExpire+5)*time.Minute),
	}
	return newClients
}
