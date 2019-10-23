package option

type Option struct {
	Port             int
	UseRedis         bool
	RedisHost        string // IP:port
	RedisMaxIdle     int
	RedisActive      int
	RedisIdleTimeout int // second
	RedisKeyExpire   int // second
	GeoResolution    int
}

var (
	Opt = &Option{}
)
