package option

type Option struct {
	MetadataHost     string
	TrackerHost      string
	RedisHost        string // IP:port
	RedisMaxIdle     int
	RedisActive      int
	RedisIdleTimeout int // second
	RedisKeyExpire   int // second
}

var (
	Opt = &Option{}
)
