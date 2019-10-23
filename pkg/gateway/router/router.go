package router

import (
	"context"
	"fmt"

	"github.com/Bo0km4n/arc/pkg/gateway/api/handler"
	"github.com/Bo0km4n/arc/pkg/gateway/cmd/option"
	"github.com/Bo0km4n/arc/pkg/gateway/domain/repository"
	"github.com/Bo0km4n/arc/pkg/gateway/infra/db"
	"github.com/Bo0km4n/arc/pkg/gateway/usecase"
	metadataclient "github.com/Bo0km4n/arc/pkg/metadata/client"
	trackerclient "github.com/Bo0km4n/arc/pkg/tracker/client"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Router struct {
	mc     metadataclient.Client
	tc     trackerclient.Client
	engine *gin.Engine
}

// New returns a router object binding some functions.
func New(ctx context.Context, logger *zap.Logger, opt *option.Option) (*Router, error) {
	r := &Router{}
	r.engine = gin.New()
	r.engine.Use(gin.Recovery())
	r.engine.Use(gin.Logger())
	if opt.MetadataHost != "" {
		cli, err := metadataclient.NewClient(
			ctx,
			opt.MetadataHost,
			grpc.WithInsecure(),
		)
		if err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("connected metadata host: %s", opt.MetadataHost))
		r.mc = cli
	}

	if opt.TrackerHost != "" {
		cli, err := trackerclient.NewClient(
			ctx,
			opt.TrackerHost,
			grpc.WithInsecure(),
		)
		if err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("connected tracker host: %s", opt.TrackerHost))
		r.tc = cli
	}

	{
		r.engine.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"Message": "pong",
			})
		})
	}

	{
		// Member
		lockerRepo := repository.NewLockerRepository(db.RedisPool, 10, int64(100), int64(1), int64(100))
		metadataRepo := repository.NewMetadataRepository(r.mc)
		trackerRepo := repository.NewTrackerRepository(r.tc)
		muc := usecase.NewMemberUsecase(metadataRepo, trackerRepo, lockerRepo)
		handler.MemberResource(r.engine, muc, logger)
	}

	return r, nil
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}

func (r *Router) Close() {
	r.mc.Close()
	r.tc.Close()
}
