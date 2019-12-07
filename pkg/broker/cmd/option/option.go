package option

type Option struct {
	Port              int
	UseRedis          bool
	Ws                bool
	RedisHost         string // IP:port
	RedisMaxIdle      int
	RedisActive       int
	RedisIdleTimeout  int // second
	RedisKeyExpire    int // second
	MysqlHost         string
	MysqlPort         string
	MysqlUser         string
	MysqlPassword     string
	MysqlDatabase     string
	MysqlMaxOpenConns int
	MysqlMaxIdleConns int
	GeoResolution     int
}

var (
	Opt = &Option{}
)
