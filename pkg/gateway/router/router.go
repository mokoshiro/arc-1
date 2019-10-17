package router

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/gateway/api/handler"
	"github.com/Bo0km4n/arc/pkg/gateway/cmd/option"
	"github.com/Bo0km4n/arc/pkg/gateway/domain/repository"
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
		logger.Info("connected metadata host")
		r.mc = cli
	}

	if opt.TrackerHost != "" {
		cli, err := trackerclient.NewClient(
			ctx,
			opt.MetadataHost,
			grpc.WithInsecure(),
		)
		if err != nil {
			return nil, err
		}
		logger.Info("connected tracker host")
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
		// Register
		metadataRepo := repository.NewMetadataRepository(r.mc)
		trackerRepo := repository.NewTrackerRepository(r.tc)
		ruc := usecase.NewRegisterUsecase(metadataRepo, trackerRepo)
		handler.RegisterResource(r.engine, ruc)
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
