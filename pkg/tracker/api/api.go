package api

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/tracker/api/proto"
	"github.com/Bo0km4n/arc/pkg/tracker/usecase"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type trackerAPI struct {
	logger     *zap.Logger
	registerUC usecase.RegisterUsecase
}

func NewtrackerAPI(ruc usecase.RegisterUsecase) *trackerAPI {
	return &trackerAPI{
		registerUC: ruc,
	}
}

func (api *trackerAPI) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	if err := api.registerUC.Register(req); err != nil {
		return nil, err
	}
	return &proto.RegisterResponse{}, nil
}

func (api *trackerAPI) Embed(server *grpc.Server, logger *zap.Logger) {
	api.logger = logger
	proto.RegisterTrackerServer(server, api)
}
