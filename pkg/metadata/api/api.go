package api

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/metadata/api/proto"
	"github.com/Bo0km4n/arc/pkg/metadata/usecase"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type metadataAPI struct {
	logger     *zap.Logger
	registerUC usecase.RegisterUsecase
}

func NewMetadataAPI(ruc usecase.RegisterUsecase) *metadataAPI {
	return &metadataAPI{
		registerUC: ruc,
	}
}

func (api *metadataAPI) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return nil, nil
}

func (api *metadataAPI) Embed(server *grpc.Server, logger *zap.Logger) {
	api.logger = logger
	proto.RegisterMetadataServer(server, api)
}
