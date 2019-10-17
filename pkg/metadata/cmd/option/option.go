package option

type Option struct {
	Port             string
	UseRedis         bool
	RedisHost        string // IP:port
	RedisMaxIdle     int
	RedisActive      int
	RedisIdleTimeout int // second
	RedisKeyExpire   int // second
}

var (
	Opt = &Option{}
)
