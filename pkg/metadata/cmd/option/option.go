package option

type Option struct {
	RedisHost        string // IP:port
	RedisMaxIdle     int
	RedisActive      int
	RedisIdleTimeout int // second
}

var (
	Opt = &Option{}
)
