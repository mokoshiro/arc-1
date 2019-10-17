package repository

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/metadata/api/proto"
	metadataclient "github.com/Bo0km4n/arc/pkg/metadata/client"
	"golang.org/x/xerrors"
)

type MetadataRepository interface {
	Register(ctx context.Context, peerID, addr string) error
}

type metadataRepository struct {
	metadataclient.Client
}

func (mr *metadataRepository) Register(ctx context.Context, peerID, addr string) error {
	_, err := mr.Client.Register(ctx, &proto.RegisterRequest{
		Id:   peerID,
		Addr: addr,
	})

	if err != nil {
		return xerrors.Errorf("metadataRepository.Register: %w", err)
	}
	return nil
}

func NewMetadataRepository(client metadataclient.Client) MetadataRepository {
	return &metadataRepository{
		Client: client,
	}
}
