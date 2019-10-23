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
	registerUC usecase.MemberUsecase
}

func NewtrackerAPI(ruc usecase.MemberUsecase) *trackerAPI {
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

func (api *trackerAPI) GetMemberByRadius(ctx context.Context, req *proto.GetMemberByRadiusRequest) (*proto.GetMemberByRadiusResponse, error) {
	res, err := api.registerUC.GetMemberByRadius(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (api *trackerAPI) Embed(server *grpc.Server, logger *zap.Logger) {
	api.logger = logger
	proto.RegisterTrackerServer(server, api)
}
