package repository

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/metadata/api/proto"
	metadataclient "github.com/Bo0km4n/arc/pkg/metadata/client"
	"golang.org/x/xerrors"
)

type MetadataRepository interface {
	Register(ctx context.Context, peerID, addr string) error
	GetMember(ctx context.Context, req *proto.GetMemberRequest) (*proto.GetMemberResponse, error)
	Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error)
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

func (mr *metadataRepository) GetMember(ctx context.Context, req *proto.GetMemberRequest) (*proto.GetMemberResponse, error) {
	return mr.Client.GetMember(ctx, req)
}

func (mr *metadataRepository) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return mr.Client.Delete(ctx, req)
}

func NewMetadataRepository(client metadataclient.Client) MetadataRepository {
	return &metadataRepository{
		Client: client,
	}
}
