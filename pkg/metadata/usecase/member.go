package usecase

import (
	"context"

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
	return nil, nil
}

func NewMemberUsecase(mr repository.MemberRepository) MemberUsecase {
	return &memberUsecase{
		mr: mr,
	}
}
