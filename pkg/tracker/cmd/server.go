package cmd

import (
	"context"
	"log"

	"net/http"

	"github.com/Bo0km4n/arc/pkg/tracker/api"
	"github.com/Bo0km4n/arc/pkg/tracker/cmd/option"
	"github.com/Bo0km4n/arc/pkg/tracker/infra/db"
	"github.com/Bo0km4n/arc/pkg/tracker/logger"
	"github.com/Bo0km4n/arc/pkg/tracker/socket"
	"github.com/spf13/cobra"
)

// serverCmd represents the web command
var serverCmd = &cobra.Command{
	Use: "server",
	PreRun: func(cmd *cobra.Command, args []string) {
		db.InitDB()
		logger.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Run Gateway API task
		Server(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().BoolVarP(&option.Opt.Ws, "ws", "", true, "use websockets")
	serverCmd.Flags().IntVarP(&option.Opt.Port, "port", "p", 50051, "listen port")
	serverCmd.Flags().BoolVarP(&option.Opt.UseRedis, "use_redis", "", true, "use redis")
	serverCmd.Flags().StringVarP(&option.Opt.RedisHost, "redis_host", "", "127.0.0.1:6379", "redis host address")
	serverCmd.Flags().IntVarP(&option.Opt.RedisMaxIdle, "redis_max_idle", "", 32, "redis max idle connection")
	serverCmd.Flags().IntVarP(&option.Opt.RedisActive, "redis_max_active", "", 64, "redis max active connection")
	serverCmd.Flags().IntVarP(&option.Opt.RedisIdleTimeout, "redis_idle_timeout", "", 240, "redis idle timeout connection")
	serverCmd.Flags().IntVarP(&option.Opt.RedisKeyExpire, "redis_key_expire", "", 86400, "redis key expire")
	serverCmd.Flags().IntVarP(&option.Opt.GeoResolution, "geo_resolution", "", 9, "Geo hash resolution")
	serverCmd.Flags().StringVarP(&option.Opt.MysqlHost, "mysql_host", "", "127.0.0.1", "mysql host address")
	serverCmd.Flags().StringVarP(&option.Opt.MysqlPort, "mysql_port", "", "3306", "mysql port")
	serverCmd.Flags().StringVarP(&option.Opt.MysqlPassword, "mysql_password", "", "root", "mysql password")
	serverCmd.Flags().StringVarP(&option.Opt.MysqlUser, "mysql_user", "", "root", "mysql user")
	serverCmd.Flags().StringVarP(&option.Opt.MysqlDatabase, "mysql_database", "", "arc", "mysql database")
}

// Server returns API object
func Server(ctx context.Context) {
	trackerAPI := api.NewHTTPTrackerAPI()

	if !option.Opt.Ws {
		trackerAPI.Run()
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			socket.Serve(w, r)
		})
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal(err)
		}
	}
}
