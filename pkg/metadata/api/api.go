package api

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/metadata/api/proto"
	"github.com/Bo0km4n/arc/pkg/metadata/usecase"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type metadataAPI struct {
	logger   *zap.Logger
	memberUC usecase.MemberUsecase
}

func NewMetadataAPI(ruc usecase.MemberUsecase) *metadataAPI {
	return &metadataAPI{
		memberUC: ruc,
	}
}

func (api *metadataAPI) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	if err := api.memberUC.Register(req.Id, req.Addr); err != nil {
		return nil, err
	}
	return &proto.RegisterResponse{}, nil
}

func (api *metadataAPI) GetMember(ctx context.Context, req *proto.GetMemberRequest) (*proto.GetMemberResponse, error) {
	return api.memberUC.GetMember(ctx, req)
}

func (api *metadataAPI) Embed(server *grpc.Server, logger *zap.Logger) {
	api.logger = logger
	proto.RegisterMetadataServer(server, api)
}
