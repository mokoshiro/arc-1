package usecase

import (
	"context"

	"fmt"

	"github.com/Bo0km4n/arc/pkg/metadata/api/proto"
	"github.com/Bo0km4n/arc/pkg/metadata/domain/repository"
)

type MemberUsecase interface {
	Register(peerID, addr string) error
	GetMember(ctx context.Context, req *proto.GetMemberRequest) (*proto.GetMemberResponse, error)
}

type memberUsecase struct {
	mr repository.MemberRepository
}

func (mu *memberUsecase) Register(peerID, addr string) error {
	return mu.mr.Register(peerID, addr)
}

func (mu *memberUsecase) GetMember(ctx context.Context, req *proto.GetMemberRequest) (*proto.GetMemberResponse, error) {
	peerAddrs, err := mu.mr.GetMember(ctx, req.PeerIds)
	if err != nil {
		return nil, err
	}
	if len(peerAddrs) != len(req.PeerIds) {
		return nil, fmt.Errorf("Not matched get peer addrs slice length, expected=%d, got=%d", len(req.PeerIds), len(peerAddrs))
	}
	res := &proto.GetMemberResponse{
		Members: make([]*proto.MetadataMember, len(req.PeerIds)),
	}
	for i := range peerAddrs {
		res.Members[i] = &proto.MetadataMember{
			PeerId: req.PeerIds[i],
			Addr:   peerAddrs[i],
		}
	}
	return res, nil
}

func NewMemberUsecase(mr repository.MemberRepository) MemberUsecase {
	return &memberUsecase{
		mr: mr,
	}
}
